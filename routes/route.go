package routes

import (
	"net/http"
	"sims_ppob/dto"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {
	// db, err := config.InitDB()
	// if err != nil {
	// 	log.Fatal("Failed Connect to Database")
	// }

	e.GET("/health", func(c echo.Context) error {
		apiResponse := dto.ApiResponse{
			Status:  http.StatusOK,
			Message: "API is healthy",
		}

		return c.JSON(http.StatusOK, apiResponse)
	})

	// USER ROUTES
	// UserController := controllers.NewUserController(db)
	// e.POST("/registration", UserController.Register)
	// e.POST("/login", UserController.Login)
	// e.GET("/profile", UserController.Logout, middlewares.JWTMiddleware)
	// e.PUT("/profile/update", UserController.UpdateUser, middlewares.JWTMiddleware)
	// e.PUT("/profile/image", UserController.UpdateUser, middlewares.JWTMiddleware)

	// INFORMATION ROUTES
	// e.GET("/banner", UserController.RefreshToken)
	// e.GET("/services", UserController.RefreshToken, middlewares.JWTMiddleware)

	// TRANSACTION ROUTES
	// e.GET("/balance", UserController.RefreshToken, middlewares.JWTMiddleware)
	// e.POST("/topup", UserController.RefreshToken, middlewares.JWTMiddleware)
	// e.POST("/transaction", UserController.RefreshToken, middlewares.JWTMiddleware)
	// e.GET("/transaction/history", UserController.RefreshToken, middlewares.JWTMiddleware)
}
