package controllers

import (
	"database/sql"
	"sims_ppob/services"

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
	validate    *sql.DB
}

func NewUserController(db *sql.DB) UserControllerImpl {
	services := services.NewUserService(db)
	return UserControllerImpl{
		UserService: services,
	}
}
