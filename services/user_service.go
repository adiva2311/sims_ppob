package services

import (
	"database/sql"
	"errors"
	"sims_ppob/dto"
	"sims_ppob/models"
	"sims_ppob/repositories"
	"sims_ppob/utils"

	"github.com/go-playground/validator/v10"
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
	validate       *validator.Validate
}

// Register implements UserService.
func (u *UserServiceImpl) Register(registerRequest models.User) (dto.RegisterResponse, error) {
	// Check if email already exists
	exists, _ := u.userRepository.CheckEmailExists(registerRequest.Email)
	if !exists {
		return dto.RegisterResponse{}, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		return dto.RegisterResponse{}, errors.New("failed to hash password")
	}

	user := &models.User{
		Email:     registerRequest.Email,
		Password:  hashedPassword,
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
	}

	if err := u.userRepository.Register(user); err != nil {
		return dto.RegisterResponse{}, errors.New("failed to register user")
	}

	return *dto.ToUserResponse(user), nil
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

func NewUserService(db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: repositories.NewUserRepository(db),
		validate:       validate,
	}
}
