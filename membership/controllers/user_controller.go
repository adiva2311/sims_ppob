package controllers

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sims_ppob/dto"
	req "sims_ppob/membership/dto"
	"sims_ppob/membership/models"
	"sims_ppob/membership/services"
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
	if err := utils.ValidateStruct(userPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
			Data:    err,
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

// Login implements UserController.
func (u *UserControllerImpl) Login(c echo.Context) error {
	userPayload := new(req.LoginRequest)
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

	result, err := u.UserService.Login(req.LoginRequest{
		Email:    userPayload.Email,
		Password: userPayload.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal login: " + err.Error(),
			Data:    err,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Login Sukses",
		Data:    result,
	}
	return c.JSON(http.StatusOK, apiResponse)
}

func (u *UserControllerImpl) GetProfile(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	// Call service to get user profile
	result, err := u.UserService.GetProfile(userEmail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal mendapatkan profil: " + err.Error(),
			Data:    nil,
		})
	}
	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mendapatkan profil",
		Data:    result,
	}
	return c.JSON(http.StatusOK, apiResponse)
}

func (u *UserControllerImpl) UpdateProfile(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	userPayload := new(req.UpdateProfileRequest)
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

	// Call service to update user profile
	result, err := u.UserService.UpdateProfile(userEmail, models.User{
		FirstName: userPayload.FirstName,
		LastName:  userPayload.LastName,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal memperbarui profil: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Update Pofile berhasil",
		Data:    result,
	}
	return c.JSON(http.StatusOK, apiResponse)
}

func storeImage(c echo.Context, email string) (string, error) {
	// Limit file size to 2MB
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, 2<<20)
	if err := c.Request().ParseMultipartForm(2 << 20); err != nil {
		return "", err
	}
	defer c.Request().MultipartForm.RemoveAll()

	// Get file from FORM
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	if file.Size > 2*1024*1024 {
		return "", fmt.Errorf("file size exceeds 2MB limit")
	}

	// Validate file type
	fileType := file.Header.Get("Content-Type")
	if fileType != "image/jpeg" && fileType != "image/png" && fileType != "image/jpg" {
		return "", fmt.Errorf("only JPEG, JPG, and PNG images are allowed")
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create destination file
	storagePath := "./img/profile_image/"
	file.Filename = email + "_profile" + filepath.Ext(file.Filename)
	path := filepath.Join(storagePath, file.Filename)
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the file content to destination
	if _, err = io.Copy(dst, src); err != nil {
		return "", c.JSON(http.StatusBadRequest, err)
	}

	return storagePath + file.Filename, nil
}

func (u *UserControllerImpl) UpdateImage(c echo.Context) error {
	// Get user Email from JWT token
	userEmail, ok := c.Get("userEmail").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	// Store image
	path, err := storeImage(c, userEmail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal menyimpan gambar: " + err.Error(),
			Data:    nil,
		})
	}

	// Call service to update user profile
	result, err := u.UserService.UpdateImage(userEmail, models.User{
		ProfileImage: path,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Gagal memperbarui gambar: " + err.Error(),
			Data:    nil,
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Update Pofile berhasil",
		Data:    result,
	}
	return c.JSON(http.StatusOK, apiResponse)
}

func NewUserController(db *sql.DB) UserControllerImpl {
	services := services.NewUserService(db)
	return UserControllerImpl{
		UserService: services,
	}
}
