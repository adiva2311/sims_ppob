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
	GetProfile(email string) (dto.UserProfileResponse, error)
	UpdateProfile(email string, updateProfileRequest models.User) (dto.UserProfileResponse, error)
	UpdateImage(email string, updateImageRequest models.User) (dto.UserProfileResponse, error)
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
	jwtToken, err := utils.GenerateJWT(int64(user.ID), user.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return *dto.ToLoginResponse(jwtToken), nil
}

// GetProfile implements UserService.
func (u *UserServiceImpl) GetProfile(email string) (dto.UserProfileResponse, error) {
	// Fetch user profile from repository
	userProfile, err := u.userRepository.GetProfile(email)
	if err != nil {
		return dto.UserProfileResponse{}, err
	}

	return *dto.ToUserProfileResponse(*userProfile), nil
}

// UpdateProfile implements UserService.
func (u *UserServiceImpl) UpdateProfile(email string, updateProfileRequest models.User) (dto.UserProfileResponse, error) {
	user := &models.User{
		FirstName: updateProfileRequest.FirstName,
		LastName:  updateProfileRequest.LastName,
	}

	_, err := u.userRepository.UpdateProfile(email, user)
	if err != nil {
		return dto.UserProfileResponse{}, err
	}

	userProfile, _ := u.userRepository.GetProfile(email)

	return *dto.ToUserProfileResponse(*userProfile), nil
}

// UpdateImage implements UserService.
func (u *UserServiceImpl) UpdateImage(email string, updateImageRequest models.User) (dto.UserProfileResponse, error) {
	user := &models.User{
		ProfileImage: updateImageRequest.ProfileImage,
	}

	_, err := u.userRepository.UpdateImage(email, user)
	if err != nil {
		return dto.UserProfileResponse{}, err
	}

	userProfile, _ := u.userRepository.GetProfile(email)

	return *dto.ToUserProfileResponse(*userProfile), nil
}

func NewUserService(db *sql.DB) UserService {
	return &UserServiceImpl{
		userRepository: repositories.NewUserRepository(db),
	}
}
