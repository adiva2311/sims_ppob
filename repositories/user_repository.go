package repositories

import (
	"database/sql"
	"fmt"
	"sims_ppob/models"
)

type UserRepository interface {
	Register(user models.User) error
	CheckEmailValid(email string) (*models.User, error)
	CheckEmailExists(email string) (bool, error)
	GetProfile(user_id uint) (*models.User, error)
	UpdateProfile(user_id uint, user *models.User) error
	UpdateImage(user_id uint, user *models.User) error
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

// Register implements UserRepository.
func (u *UserRepositoryImpl) Register(user models.User) error {
	query := "INSERT INTO users (email, password, first_name, last_name) VALUES (?, ?, ?, ?)"
	result, err := u.DB.Exec(query, user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = uint(userID)
	return nil
}

// CheckEmailExists implements UserRepository.
func (u *UserRepositoryImpl) CheckEmailExists(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = ? AND deleted_at IS NULL"
	fmt.Println(u.DB.QueryRow(query, email).Scan(&count))
	err := u.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CheckEmailValid implements UserRepository.
func (u *UserRepositoryImpl) CheckEmailValid(email string) (*models.User, error) {
	var user models.User
	query := "SELECT email FROM users WHERE email = ? AND deleted_at IS NULL"
	result := u.DB.QueryRow(query, email).Scan(&user.Email)
	if result != nil {
		return nil, result
	}
	return &user, nil
}

// GetProfile implements UserRepository.
func (u *UserRepositoryImpl) GetProfile(user_id uint) (*models.User, error) {
	query := "SELECT id, email, first_name, last_name, profile_image FROM users WHERE id = ? AND deleted_at IS NULL"
	var user models.User
	result := u.DB.QueryRow(query, user_id).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.ProfileImage)
	if result != nil {
		return nil, result
	}
	return &user, nil
}

// UpdateImage implements UserRepository.
func (u *UserRepositoryImpl) UpdateImage(user_id uint, user *models.User) error {
	panic("unimplemented")
}

// UpdateProfile implements UserRepository.
func (u *UserRepositoryImpl) UpdateProfile(user_id uint, user *models.User) error {
	panic("unimplemented")
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
