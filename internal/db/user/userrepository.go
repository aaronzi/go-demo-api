package db

import (
	"database/sql"
	"errors"
	"log"

	auth "go-demo-api/internal/auth"
	db "go-demo-api/internal/db"
	utils "go-demo-api/internal/util"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var UserRepo_NotFoundError = errors.New("identifier or password wrong")

type UserRepository struct {
	DB                     *sql.DB
	VerificationRepository db.VerificationRepositoryInterface
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
		log.Fatalf("Error hashing password: %v", err)
		return err
	}
	_, err = repo.DB.Exec("INSERT INTO Users (id, username, email, password_hash) VALUES (?, ?, ?, ?)", userid, username, email, hashedpassword)

	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
		return err
	}

	_, err = repo.VerificationRepository.CreateVerification(email)

	if err != nil {
		log.Fatalf("Error inserting verification: %v", err)
		return err
	}

	return nil
}

func (repo *UserRepository) LoginUser(identifier string, password string) (string, error) {
	// Changed the error message to 'identifier' to generalize username/email
	var user User

	// Adjust the SQL query to check both the username and email fields
	err := repo.DB.QueryRow("SELECT * FROM Users WHERE username = ? OR email = ?", identifier, identifier).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return "", UserRepo_NotFoundError
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", UserRepo_NotFoundError
	}
	secret, fileReadError := utils.ReadFile("/workspace/privatekey.txt")

	if fileReadError != nil {
		log.Fatalf("Error reading file: %v", fileReadError)
		return "", fileReadError
	}

	jwt, err := auth.GenerateJWT(user.ID, secret)
	if err != nil {
		log.Fatalf("Error generating JWT: %v", err)
		return "", err
	}

	return jwt, nil
}