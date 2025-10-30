package models

type Transaction struct {
	ID              uint   `json:"id"`
	InvoiceNumber   string `json:"invoice_number" validate:"required"`
	UserID          uint   `json:"user_id"`
	ServiceID       uint   `json:"service_id"`
	TotalAmount     int64  `json:"total_amount" validate:"required, gte=1"`
	TransactionType string `json:"transaction_type" validate:"required"`
	Description     string `json:"description"`
	CreatedAt       string `json:"created_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}
