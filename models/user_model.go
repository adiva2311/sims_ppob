package models

type User struct {
	ID           uint   `json:"id"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Balance      int64  `json:"balance"`
	ProfileImage string `json:"profile_image"`
}
