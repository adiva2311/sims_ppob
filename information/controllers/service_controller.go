package controllers

import (
	"database/sql"
	"net/http"
	"sims_ppob/dto"
	"sims_ppob/information/services"

	"github.com/labstack/echo/v4"
)

type ServiceController interface {
	FindAllServices(c echo.Context) error
}

type ServiceControllerImpl struct {
	ServiceService services.ServiceService
}

func (s *ServiceControllerImpl) FindAllServices(c echo.Context) error {
	result, err := s.ServiceService.FindAllServices()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "Gagal mengambil data service: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Sukses mengambil data service",
		Data:    result,
	}
	return c.JSON(http.StatusOK, apiResponse)
}

func NewServiceController(db *sql.DB) ServiceControllerImpl {
	service := services.NewServiceService(db)
	controller := ServiceControllerImpl{
		ServiceService: service,
	}
	return controller
}
