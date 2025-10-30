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
	GetProfile(email string) (*models.User, error)
	UpdateProfile(email string, user *models.User) (*models.User, error)
	UpdateImage(email string, user *models.User) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

// Register implements UserRepository.
func (u *UserRepositoryImpl) Register(user models.User) error {
	query := "INSERT INTO users (email, password, first_name, last_name, profile_image) VALUES (?, ?, ?, ?, ?)"
	result, err := u.DB.Exec(query, user.Email, user.Password, user.FirstName, user.LastName, user.ProfileImage)
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
	query := "SELECT email, password FROM users WHERE email = ? AND deleted_at IS NULL"
	err := u.DB.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetProfile implements UserRepository.
func (u *UserRepositoryImpl) GetProfile(email string) (*models.User, error) {
	var user models.User
	query := "SELECT email, first_name, last_name, profile_image FROM users WHERE email = ? AND deleted_at IS NULL"
	err := u.DB.QueryRow(query, email).Scan(&user.Email, &user.FirstName, &user.LastName, &user.ProfileImage)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateImage implements UserRepository.
func (u *UserRepositoryImpl) UpdateImage(email string, user *models.User) (*models.User, error) {
	query := "UPDATE users SET profile_image = ? WHERE email = ? AND deleted_at IS NULL"
	result, err := u.DB.Exec(query, user.ProfileImage, email)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return user, nil
	}
	return user, nil
}

// UpdateProfile implements UserRepository.
func (u *UserRepositoryImpl) UpdateProfile(email string, user *models.User) (*models.User, error) {
	query := "UPDATE users SET first_name = ?, last_name = ? WHERE email = ? AND deleted_at IS NULL"
	result, err := u.DB.Exec(query, user.FirstName, user.LastName, email)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return user, nil
	}
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
