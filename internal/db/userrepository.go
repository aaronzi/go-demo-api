package db

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}
type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (repo *UserRepository) RegisterUser(username string, email string, password string) error {
	userid := uuid.New().String()
	hashedpassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	_, err = repo.DB.Exec("INSERT INTO Users (id, username, email, password_hash) VALUES (?, ?, ?, ?)", userid, username, email, hashedpassword)

	if err != nil {
		return err
	}

	database, err := NewDB()

	// Instantiate the repository
	verificationRepo := &VerificationRepository{DB: database}
	verificationError := verificationRepo.CreateVerification(email)

	if verificationError != nil {
		return verificationError
	}

	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) CheckUser(username string, password string) error {
	var NotFoundError = errors.New("username or password wrong")
	var user User

	err := repo.DB.QueryRow("SELECT * FROM Users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return NotFoundError
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return NotFoundError
	}
	return nil
}
