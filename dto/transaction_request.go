package dto

type BalanceRequest struct {
	Balance int64 `json:"balance"`
}

type TopUpRequest struct {
	TopUpAmount int64 `json:"top_up_amount" validate:"required,gt=0"`
}

type PaymentRequest struct {
	ServiceCode string `json:"service_code" validate:"required"`
}
