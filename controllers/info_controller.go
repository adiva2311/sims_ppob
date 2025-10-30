package controllers

import (
	"database/sql"
	"net/http"
	"sims_ppob/dto"
	"sims_ppob/services"

	"github.com/labstack/echo/v4"
)

type InfoController interface {
	FindAllBanners(c echo.Context) error
	FindAllServices(c echo.Context) error
}

type InfoControllerImpl struct {
	infoService services.InfoService
}

func (i *InfoControllerImpl) FindAllBanners(c echo.Context) error {
	result, err := i.infoService.FindAllBanners()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Gagal mengambil data banner: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Sukses mengambil data banner",
		Data:    result,
	}
	return c.JSON(http.StatusOK, apiResponse)
}

func (i *InfoControllerImpl) FindAllServices(c echo.Context) error {
	result, err := i.infoService.FindAllServices()
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

func NewInfoController(db *sql.DB) InfoControllerImpl {
	service := services.NewInfoService(db)
	controller := InfoControllerImpl{
		infoService: service,
	}
	return controller
}
