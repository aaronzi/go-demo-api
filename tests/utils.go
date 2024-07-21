package tests

import (
	"go-demo-api/internal/auth"
	utils "go-demo-api/internal/util"
	"log"
)

func GenerateToken() (string, error) {
	secret, fileReadError := utils.ReadFile("/workspace/privatekey.txt")

	if fileReadError != nil {
		log.Fatalf("Error reading file: %v", fileReadError)
		return "", fileReadError
	}

	jwt, err := auth.GenerateJWT("testUserID", secret)
	if err != nil {
		log.Fatalf("Error generating JWT: %v", err)
		return "", err
	}

	return jwt, nil
}
