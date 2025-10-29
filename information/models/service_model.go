package models

type Services struct {
	ID            uint   `json:"id"`
	ServiceCode   string `json:"service_code" validate:"required"`
	ServiceName   string `json:"service_name" validate:"required"`
	ServiceIcon   string `json:"service_icon" validate:"required"`
	ServiceTariff int64  `json:"service_tariff" validate:"required"`
}

func (Services) TableName() string {
	return "services"
}
