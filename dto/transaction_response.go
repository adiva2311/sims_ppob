package dto

import "sims_ppob/models"

type BalanceResponse struct {
	Balance int64 `json:"balance"`
}

func ToBalanceResponse(user models.User) BalanceResponse {
	return BalanceResponse{
		Balance: user.Balance,
	}
}

type TransactionResponse struct {
	InvoiceNumber   string `json:"invoice_number"`
	ServiceCode     string `json:"service_code"`
	ServiceName     string `json:"service_name"`
	TransactionType string `json:"transaction_type"`
	TotalAmount     int64  `json:"total_amount"`
	CreatedAt       string `json:"created_at"`
}

func ToTransactionResponse(service models.Services, transaction models.Transaction) TransactionResponse {
	return TransactionResponse{
		InvoiceNumber:   transaction.InvoiceNumber,
		ServiceCode:     service.ServiceCode,
		ServiceName:     service.ServiceName,
		TransactionType: transaction.TransactionType,
		TotalAmount:     transaction.TotalAmount,
		CreatedAt:       transaction.CreatedAt,
	}
}

type TransactionHistoryResponse struct {
	InvoiceNumber   string `json:"invoice_number"`
	TransactionType string `json:"transaction_type"`
	Description     string `json:"description"`
	TotalAmount     int64  `json:"total_amount"`
	CreatedAt       string `json:"created_at"`
}

func ToTransactionHistoryResponse(transactions []models.Transaction) []TransactionHistoryResponse {
	var response []TransactionHistoryResponse
	for _, transaction := range transactions {
		response = append(response, TransactionHistoryResponse{
			InvoiceNumber:   transaction.InvoiceNumber,
			TransactionType: transaction.TransactionType,
			Description:     transaction.Description,
			TotalAmount:     transaction.TotalAmount,
			CreatedAt:       transaction.CreatedAt,
		})
	}
	return response
}
