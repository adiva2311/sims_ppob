package services

import (
	"database/sql"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/repositories"
)

type TransactionService interface {
	Balance(email string) (dto.BalanceResponse, error)
	TopUp(email string, topUp models.User) (int64, error)
	Payment(payment dto.PaymentRequest, email string) (int64, error)
	PaymentHistory(email string) ([]dto.TransactionHistoryResponse, error)
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
	return dto.BalanceResponse{Balance: int64(currentBalance)}, nil
}

// Payment implements TransactionService.
func (t *TransactionServiceImpl) Payment(payment dto.PaymentRequest, email string) (int64, error) {
	panic("unimplemented")
}

// PaymentHistory implements TransactionService.
func (t *TransactionServiceImpl) PaymentHistory(email string) ([]dto.TransactionHistoryResponse, error) {
	panic("unimplemented")
}

// TopUp implements TransactionService.
func (t *TransactionServiceImpl) TopUp(email string, topUp models.User) (int64, error) {
	user := &models.User{
		Balance: topUp.Balance,
	}

	newBalance, err := t.TransactionRepository.TopUp(email, int(user.Balance))
	if err != nil {
		return 0, err
	}

	return int64(newBalance), nil
}

func NewTransactionService(db *sql.DB) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: repositories.NewTransactionRepository(db),
	}
}
