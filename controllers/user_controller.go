package controllers

import (
	"database/sql"
	"net/http"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/services"
	"sims_ppob/utils"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetProfile(c echo.Context) error
	UpdateProfile(c echo.Context) error
	UpdateImage(c echo.Context) error
}

type UserControllerImpl struct {
	UserService services.UserService
	// validate    *sql.DB
}

func (u *UserControllerImpl) Register(c echo.Context) error {
	userPayload := new(models.User)
	if err := c.Bind(userPayload); err != nil {
		return err
	}

	if userPayload.ProfileImage == "" {
		userPayload.ProfileImage = "./img/profile_image/default_user.png"
	}

	// Validasi input from user
	if errs := utils.ValidateStruct(userPayload); errs != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
			Data:    errs,
		})
	}

	result, err := u.UserService.Register(models.User{
		Email:        userPayload.Email,
		Password:     userPayload.Password,
		FirstName:    userPayload.FirstName,
		LastName:     userPayload.LastName,
		ProfileImage: userPayload.ProfileImage,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal register: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil register",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (u *UserControllerImpl) Login(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (u *UserControllerImpl) GetProfile(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (u *UserControllerImpl) UpdateProfile(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (u *UserControllerImpl) UpdateImage(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func NewUserController(db *sql.DB) UserControllerImpl {
	services := services.NewUserService(db)
	return UserControllerImpl{
		UserService: services,
	}
}
