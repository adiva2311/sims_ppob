package dto

type BalanceRequest struct {
	Balance int64 `json:"balance"`
}

type TopUpRequest struct {
	TransactionType string `json:"transaction_type"`
	TopUpAmount     int64  `json:"top_up_amount"`
}

type PaymentRequest struct {
	ServiceCode string `json:"service_code"`
}
