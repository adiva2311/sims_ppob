package services

import (
	"database/sql"
	"log"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/repositories"
)

type TransactionService interface {
	Balance(email string) (dto.BalanceResponse, error)
	TopUp(email string, user_id int, topUpAmount models.Transaction) (int64, error)
	Payment(email string, user_id int, serviceCode string, payment *models.Transaction) (*dto.PaymentResponse, error)
	PaymentHistory(email string, limit, offset int) ([]dto.PaymentHistoryResponse, error)
}

type TransactionServiceImpl struct {
	TransactionRepository repositories.TransactionRepository
}

// Balance implements TransactionService.
func (t *TransactionServiceImpl) Balance(email string) (dto.BalanceResponse, error) {
	currentBalance, err := t.TransactionRepository.Balance(email)
	if err != nil {
		return dto.BalanceResponse{}, err
	}
	return dto.ToBalanceResponse(*currentBalance), nil
}

// TopUp implements TransactionService.
func (t *TransactionServiceImpl) TopUp(email string, user_id int, topUpAmount models.Transaction) (int64, error) {
	log.Println("🔹 Repository.TopUp dipanggil")
	topUp, err := t.TransactionRepository.TopUp(email, user_id, &topUpAmount)
	if err != nil {
		return 0, err
	}
	return topUp.TotalAmount, nil
}

// Payment implements TransactionService.
func (t *TransactionServiceImpl) Payment(email string, user_id int, serviceCode string, payment *models.Transaction) (*dto.PaymentResponse, error) {
	log.Println("🔹 Repository.Payment dipanggil")
	paymentResult, err := t.TransactionRepository.Payment(email, user_id, serviceCode, payment)
	if err != nil {
		return nil, err
	}
	return paymentResult, nil
}

// PaymentHistory implements TransactionService.
func (t *TransactionServiceImpl) PaymentHistory(email string, limit, offset int) ([]dto.PaymentHistoryResponse, error) {
	log.Println("🔹 Repository.Payment History dipanggil")
	paymentHistory, err := t.TransactionRepository.PaymentHistory(email, limit, offset)
	if err != nil {
		return nil, err
	}
	return dto.ToTransactionHistoryResponse(paymentHistory), nil
}

func NewTransactionService(db *sql.DB) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: repositories.NewTransactionRepository(db),
	}
}
