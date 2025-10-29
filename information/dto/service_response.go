package dto

import "sims_ppob/information/models"

type ServiceResponse struct {
	ServiceCode   string `json:"service_code"`
	ServiceName   string `json:"service_name"`
	ServiceIcon   string `json:"service_icon"`
	ServiceTariff int64  `json:"service_tariff"`
}

func ToserviceResponse(services []models.Services) []ServiceResponse {
	var response []ServiceResponse
	for _, service := range services {
		response = append(response, ServiceResponse{
			ServiceCode:   service.ServiceCode,
			ServiceName:   service.ServiceName,
			ServiceIcon:   service.ServiceIcon,
			ServiceTariff: service.ServiceTariff,
		})
	}
	return response
}
