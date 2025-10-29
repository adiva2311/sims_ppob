package models

type Transaction struct {
	ID              uint    `json:"id"`
	InvoiceNumber   string  `json:"invoice_number" validate:"required"`
	UserID          uint    `json:"user_id"`
	ServiceID       uint    `json:"service_id"`
	TotalAmount     float64 `json:"total_amount" validate:"required"`
	TransactionType string  `json:"transaction_type" validate:"required"`
	CreatedAt       string  `json:"created_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}
