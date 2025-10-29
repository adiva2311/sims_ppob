package services

import (
	"database/sql"
	"errors"
	"fmt"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/repositories"
	"sims_ppob/utils"
	"strings"
)

type UserService interface {
	Register(registerRequest models.User) (dto.RegisterResponse, error)
	Login(loginRequest dto.LoginRequest) (dto.LoginResponse, error)
	GetProfile()
	UpdateProfile()
	UpdateImage()
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
	// validate       *validator.Validate
}

// Register implements UserService.
func (u *UserServiceImpl) Register(registerRequest models.User) (dto.RegisterResponse, error) {
	// Check if email already exists
	exists, err := u.userRepository.CheckEmailExists(registerRequest.Email)
	if err != nil {
		return dto.RegisterResponse{}, errors.New("failed to check email")
	}
	if exists {
		return dto.RegisterResponse{}, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		return dto.RegisterResponse{}, errors.New("failed to hash password")
	}

	// Normalize email
	registerRequest.Email = strings.TrimSpace(strings.ToLower(registerRequest.Email))

	user := &models.User{
		Email:        registerRequest.Email,
		Password:     hashedPassword,
		FirstName:    registerRequest.FirstName,
		LastName:     registerRequest.LastName,
		ProfileImage: registerRequest.ProfileImage,
	}

	if err := u.userRepository.Register(*user); err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("failed to register user: %v", err)
	}

	return dto.ToUserResponse(*user), nil
}

// Login implements UserService.
func (u *UserServiceImpl) Login(loginRequest dto.LoginRequest) (dto.LoginResponse, error) {
	// Check if email is valid
	user, err := u.userRepository.CheckEmailValid(loginRequest.Email)
	if err != nil {
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	// Verify password
	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	// Generate JWT token
	// accessToken, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	// if err != nil {
	// 	return dto.LoginResponse{}, err
	// }

	accessToken := "dummy_token" // Placeholder for JWT token generation

	return *dto.ToLoginResponse(accessToken), nil
}

// GetProfile implements UserService.
func (u *UserServiceImpl) GetProfile() {
	panic("unimplemented")
}

// UpdateImage implements UserService.
func (u *UserServiceImpl) UpdateImage() {
	panic("unimplemented")
}

// UpdateProfile implements UserService.
func (u *UserServiceImpl) UpdateProfile() {
	panic("unimplemented")
}

func NewUserService(db *sql.DB) UserService {
	return &UserServiceImpl{
		userRepository: repositories.NewUserRepository(db),
		// validate:       validate,
	}
}
