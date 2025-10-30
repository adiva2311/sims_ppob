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

type TopUpResponse struct {
	Balance int64 `json:"balance"`
}

func ToTopUpResponse(user models.User) TopUpResponse {
	return TopUpResponse{
		Balance: user.Balance,
	}
}

type PaymentResponse struct {
	InvoiceNumber   string `json:"invoice_number"`
	ServiceCode     string `json:"service_code"`
	ServiceName     string `json:"service_name"`
	TransactionType string `json:"transaction_type"`
	TotalAmount     int64  `json:"total_amount"`
	CreatedAt       string `json:"created_at"`
}

func ToPaymentResponse(service models.Services, transaction models.Transaction) PaymentResponse {
	return PaymentResponse{
		InvoiceNumber:   transaction.InvoiceNumber,
		ServiceCode:     service.ServiceCode,
		ServiceName:     service.ServiceName,
		TransactionType: transaction.TransactionType,
		TotalAmount:     transaction.TotalAmount,
		CreatedAt:       transaction.CreatedAt,
	}
}

type PaymentHistoryResponse struct {
	InvoiceNumber   string `json:"invoice_number"`
	TransactionType string `json:"transaction_type"`
	Description     string `json:"description"`
	TotalAmount     int64  `json:"total_amount"`
	CreatedAt       string `json:"created_at"`
}

func ToTransactionHistoryResponse(transactions []PaymentHistoryResponse) []PaymentHistoryResponse {
	var response []PaymentHistoryResponse
	for _, transaction := range transactions {
		response = append(response, PaymentHistoryResponse{
			InvoiceNumber:   transaction.InvoiceNumber,
			TransactionType: transaction.TransactionType,
			Description:     transaction.Description,
			TotalAmount:     transaction.TotalAmount,
			CreatedAt:       transaction.CreatedAt,
		})
	}
	return response
}
