package middlewares

import (
	"fmt"
	"net/http"
	"sims_ppob/dto"
	"sims_ppob/utils"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Cek apakah header Authorization ada dan dimulai dengan "Bearer "
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, dto.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: "Header Authorization tidak valid",
				Data:    nil,
			})
		}

		// Ambil token dari header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return utils.GetSecretKey(), nil
		})
		if err != nil || token == nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, dto.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: "Token tidak valid atau kadaluwarsa",
				Data:    nil,
			})
		}

		// safe conversion
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, dto.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: "Gagal parsing klaim token",
				Data:    nil,
			})
		}

		email, ok := claims["email"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, dto.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: "Klaim email tidak ditemukan dalam token",
				Data:    nil,
			})
		}

		c.Set("userEmail", email)

		// Kirim request ke handler berikutnya
		return next(c)
	}
}
