package routes

import (
	"log"
	"net/http"
	"sims_ppob/config"
	"sims_ppob/controllers"
	"sims_ppob/dto"
	"sims_ppob/middlewares"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed Connect to Database")
	}

	e.GET("/health", func(c echo.Context) error {
		apiResponse := dto.ApiResponse{
			Status:  http.StatusOK,
			Message: "API is healthy",
		}

		return c.JSON(http.StatusOK, apiResponse)
	})

	// USER ROUTES
	UserController := controllers.NewUserController(db)
	e.POST("/registration", UserController.Register)
	e.POST("/login", UserController.Login)
	e.GET("/profile", UserController.GetProfile, middlewares.JWTMiddleware)
	e.PUT("/profile/update", UserController.UpdateProfile, middlewares.JWTMiddleware)
	e.PUT("/profile/image", UserController.UpdateImage, middlewares.JWTMiddleware)

	// INFORMATION ROUTES
	InfoController := controllers.NewInfoController(db)
	e.GET("/banner", InfoController.FindAllBanners)
	e.GET("/services", InfoController.FindAllServices, middlewares.JWTMiddleware)

	// TRANSACTION ROUTES
	TransactionController := controllers.NewTransactionController(db)
	e.GET("/balance", TransactionController.Balance, middlewares.JWTMiddleware)
	// e.POST("/topup", UserController.RefreshToken, middlewares.JWTMiddleware)
	// e.POST("/transaction", UserController.RefreshToken, middlewares.JWTMiddleware)
	// e.GET("/transaction/history", UserController.RefreshToken, middlewares.JWTMiddleware)
}
