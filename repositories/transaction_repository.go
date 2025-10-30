package repositories

import (
	"database/sql"
	"fmt"
	"sims_ppob/models"
)

type TransactionRepository interface {
	Balance(email string) (int, error)
	TopUp(email string, amount int) (int, error)
	Payment(email string, serviceCode string, amount int) error
	PaymentHistory(email string) ([]models.Transaction, error)
}

type TransactionRepositoryImpl struct {
	db *sql.DB
}

// Balance implements TransactionRepository.
func (t *TransactionRepositoryImpl) Balance(email string) (int, error) {
	var balance int
	query := "SELECT balance FROM users WHERE email = ? AND deleted_at IS NULL"
	err := t.db.QueryRow(query, email).Scan(&balance)
	if err != nil {
		return 0, err
	}
	fmt.Println("Current balance:", balance)
	return balance, nil
}

// TopUp implements TransactionRepository.
func (t *TransactionRepositoryImpl) TopUp(email string, amount int) (int, error) {
	balance, err := t.Balance(email)
	if err != nil {
		return 0, err
	}

	totalBalance := balance + amount
	fmt.Println("Total balance after top-up:", totalBalance)

	query := "UPDATE users SET balance = balance + ? WHERE email = ? AND deleted_at IS NULL"
	result, err := t.db.Exec(query, totalBalance, email)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	fmt.Println("Rows affected:", rowsAffected)

	return totalBalance, nil
}

// Payment implements TransactionRepository.
func (t *TransactionRepositoryImpl) Payment(email string, serviceCode string, amount int) error {
	panic("unimplemented")
}

// PaymentHistory implements TransactionRepository.
func (t *TransactionRepositoryImpl) PaymentHistory(email string) ([]models.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}
