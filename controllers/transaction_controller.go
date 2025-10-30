package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/services"
	"sims_ppob/utils"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TransactionController interface {
	Balance(c echo.Context) error
	TopUp(c echo.Context) error
	Payment(c echo.Context) error
	PaymentHistory(c echo.Context) error
}

type TransactionControllerImpl struct {
	TransactionService services.TransactionService
}

func (t *TransactionControllerImpl) Balance(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	result, err := t.TransactionService.Balance(userEmail)
	log.Println(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal mendapatkan Balance: " + err.Error(),
			Data:    nil,
		})
	}

	// currentBalance := result.Balance

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
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Unauthorized: Can't get user email",
		})
	}

	// Get user ID from JWT token
	userID, ok := c.Get("id").(int64)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Unauthorized: Can't get user ID",
		})
	}

	log.Println(userID)

	// Bind request body to TopUpRequest struct
	userTopUpPayload := new(dto.TopUpRequest)
	if err := c.Bind(userTopUpPayload); err != nil {
		return err
	}

	// Validasi input from user
	if err := utils.ValidateStruct(userTopUpPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
			Data:    err,
		})
	}

	result, err := t.TransactionService.TopUp(userEmail, int(userID), models.Transaction{
		TotalAmount: userTopUpPayload.TopUpAmount,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal Top Up: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Sukses Top Up",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (t *TransactionControllerImpl) Payment(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Unauthorized: Can't get user email",
		})
	}

	// Get user ID from JWT token
	userID, ok := c.Get("id").(int64)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Unauthorized: Can't get user ID",
		})
	}

	// Bind request body to TopUpRequest struct
	userPayload := new(dto.PaymentRequest)
	if err := c.Bind(userPayload); err != nil {
		return err
	}

	// Validasi input from user
	if err := utils.ValidateStruct(userPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
			Data:    err,
		})
	}

	userPayload.ServiceCode = strings.ToTitle(userPayload.ServiceCode)

	result, err := t.TransactionService.Payment(userEmail, int(userID), userPayload.ServiceCode, &models.Transaction{
		TransactionType: "PAYMENT",
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal Payment: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Sukses Payment",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (t *TransactionControllerImpl) PaymentHistory(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Unauthorized: Can't get user email",
		})
	}

	// Get Limit & Offset Param
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 3 // default
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}

	result, err := t.TransactionService.PaymentHistory(userEmail, limit, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal mendapatkan Payment History: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Sukses mendapatkan Payment History",
		Data: dto.Pagination{
			Limit:  limit,
			Offset: offset,
			Record: result,
		},
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func NewTransactionController(db *sql.DB) TransactionControllerImpl {
	service := services.NewTransactionService(db)
	controller := TransactionControllerImpl{
		TransactionService: service,
	}
	return controller
}
