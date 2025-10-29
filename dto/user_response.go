package dto

import "sims_ppob/models"

type RegisterResponse struct {
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ProfileImage string `json:"profile_image"`
}

func ToUserResponse(user models.User) RegisterResponse {
	return RegisterResponse{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		ProfileImage: user.ProfileImage,
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

func ToLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		Token: token,
	}
}

type UserProfileResponse struct {
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	ProfileImage string `json:"profile_image"`
}

func ToUserProfileResponse(user models.User) *UserProfileResponse {
	return &UserProfileResponse{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		ProfileImage: user.ProfileImage,
	}
}
