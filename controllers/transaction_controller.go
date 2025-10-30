package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/services"
	"sims_ppob/utils"

	"github.com/labstack/echo/v4"
)

type TransactionController interface {
	Balance(c echo.Context) error
	TopUp(c echo.Context) error
	Payment(c echo.Context) error
	PaymentHistory(c echo.Context) error
}

type TransactionControllerImpl struct {
	transactionService services.TransactionService
}

func (t *TransactionControllerImpl) Balance(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	result, err := t.transactionService.Balance(userEmail)
	log.Println(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal mendapatkan Balance: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Sukses mendapatkan Balance",
		Data:    result,
	}
	return c.JSON(http.StatusOK, apiResponse)
}

func (t *TransactionControllerImpl) TopUp(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	userPayload := new(dto.TopUpRequest)
	if err := c.Bind(userPayload); err != nil {
		return err
	}

	userPayload.TransactionType = "TOPUP"

	// Validasi input from user
	if err := utils.ValidateStruct(userPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
			Data:    err,
		})
	}

	result, err := t.transactionService.TopUp(userEmail, models.User{
		Balance: userPayload.TopUpAmount,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal melakukan TopUp: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "TopUp Sukses",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (t *TransactionControllerImpl) Payment(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (t *TransactionControllerImpl) PaymentHistory(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func NewTransactionController(db *sql.DB) TransactionControllerImpl {
	service := services.NewTransactionService(db)
	controller := TransactionControllerImpl{
		transactionService: service,
	}
	return controller
}
