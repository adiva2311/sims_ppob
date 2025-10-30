package repositories

import (
	"database/sql"
	"fmt"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/utils"
)

type TransactionRepository interface {
	Balance(email string) (*models.User, error)
	TopUp(email string, user_id int, topUp *models.Transaction) (*models.Transaction, error)
	Payment(email string, user_id int, serviceCode string, payment *models.Transaction) (*dto.PaymentResponse, error)
	PaymentHistory(email string, limit, offset int) ([]dto.PaymentHistoryResponse, error)
}

type TransactionRepositoryImpl struct {
	db *sql.DB
}

// Balance implements TransactionRepository.
func (t *TransactionRepositoryImpl) Balance(email string) (*models.User, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	var user models.User
	query := `SELECT balance FROM users WHERE email = ? AND deleted_at IS NULL`
	err = tx.QueryRow(query, email).Scan(&user.Balance)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// TopUp implements TransactionRepository.
func (t *TransactionRepositoryImpl) TopUp(email string, user_id int, topUp *models.Transaction) (*models.Transaction, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	// Get current balance
	var balance int
	query := `SELECT balance FROM users WHERE email = ? AND deleted_at IS NULL`
	err = tx.QueryRow(query, email).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update balance to user by adding topUpAmount
	newBalance := balance + int(topUp.TotalAmount)
	updateBalance := `UPDATE users SET balance = ? WHERE email = ? AND deleted_at IS NULL`
	_, err = tx.Exec(updateBalance, newBalance, email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	topUp.InvoiceNumber = "INV-" + utils.GenerateInvoiceID(8)
	fmt.Println(topUp.InvoiceNumber)
	topUp.Description = "Top Up Balance"
	topUp.TransactionType = "TOP_UP"

	// Insert transaction record
	insertTransaction := `INSERT INTO transactions (invoice_number, user_id, total_amount, transaction_type, description) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(insertTransaction, topUp.InvoiceNumber, user_id, topUp.TotalAmount, topUp.TransactionType, topUp.Description)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return topUp, nil
}

// Payment implements TransactionRepository.
func (t *TransactionRepositoryImpl) Payment(email string, user_id int, serviceCode string, payment *models.Transaction) (*dto.PaymentResponse, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	// Get current balance
	var balance int
	queryBalance := `SELECT balance FROM users WHERE email = ? AND deleted_at IS NULL`
	err = tx.QueryRow(queryBalance, email).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get service price
	var service models.Services
	queryService := `SELECT id, service_name, service_tariff FROM services WHERE service_code = ?`
	err = tx.QueryRow(queryService, serviceCode).Scan(&service.ID, &service.ServiceName, &service.ServiceTariff)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("service ataus layanan tidak ditemukan")
	}

	// Check if Balance is Enough
	if balance < int(service.ServiceTariff) {
		tx.Rollback()
		return nil, fmt.Errorf("saldo tidak cukup")
	}

	// Update balance to user by reducing servicePrice
	newBalance := balance - int(service.ServiceTariff)
	updateBalance := `UPDATE users SET balance = ? WHERE email = ? AND deleted_at IS NULL`
	_, err = tx.Exec(updateBalance, newBalance, email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	payment.Description = "Payment for " + service.ServiceName
	payment.TotalAmount = service.ServiceTariff
	payment.InvoiceNumber = "INV-" + utils.GenerateInvoiceID(8)

	// Insert Payment Transaction Record
	insertTransaction := `INSERT INTO transactions (invoice_number, user_id, service_id, total_amount, transaction_type, description)
							VALUES (?, ?, ?, ?, ?, ?)`
	_, err = tx.Exec(insertTransaction, payment.InvoiceNumber, user_id, service.ID,
		payment.TotalAmount, payment.TransactionType, payment.Description)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Response Payment
	var result dto.PaymentResponse
	query := `SELECT t.invoice_number, s.service_code, s.service_name, t.transaction_type, t.total_amount, t.created_at
				FROM transactions t JOIN services s ON t.service_id = s.id
				WHERE t.invoice_number = ?`
	err = tx.QueryRow(query, payment.InvoiceNumber).Scan(&result.InvoiceNumber,
		&result.ServiceCode, &result.ServiceName, &result.TransactionType,
		&result.TotalAmount, &result.CreatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &result, nil
}

// PaymentHistory implements TransactionRepository.
func (t *TransactionRepositoryImpl) PaymentHistory(email string, limit, offset int) ([]dto.PaymentHistoryResponse, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	var data []dto.PaymentHistoryResponse
	query := `SELECT t.invoice_number, t.transaction_type, t.description, t.total_amount, t.created_at
				FROM transactions t
				JOIN users u ON t.user_id = u.id
				WHERE u.email = ?
				ORDER BY t.created_at ASC
				LIMIT ? OFFSET ?`
	rows, err := tx.Query(query, email, limit, offset)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var transaction dto.PaymentHistoryResponse
		err := rows.Scan(&transaction.InvoiceNumber, &transaction.TransactionType,
			&transaction.Description, &transaction.TotalAmount, &transaction.CreatedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		data = append(data, transaction)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return data, nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}
