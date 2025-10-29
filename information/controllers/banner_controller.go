package controllers

import (
	"database/sql"
	"net/http"
	"sims_ppob/dto"
	"sims_ppob/information/services"

	"github.com/labstack/echo/v4"
)

type BannerController interface {
	FindAllBanners(c echo.Context) error
}

type BannerControllerImpl struct {
	BannerService services.BannerService
}

func (b *BannerControllerImpl) FindAllBanners(c echo.Context) error {
	result, err := b.BannerService.FindAllBanners()
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

func NewBannerController(db *sql.DB) BannerControllerImpl {
	service := services.NewBannerService(db)
	controller := BannerControllerImpl{
		BannerService: service,
	}
	return controller
}
