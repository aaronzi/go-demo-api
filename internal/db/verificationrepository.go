package db

import (
	"errors"
	"go-demo-api/internal/util"
	"log"
	"os"

	"database/sql"
	b64 "encoding/base64"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

type VerificationRepository struct {
	DB *sql.DB
}

func GenerateVerificationCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}

func (v *VerificationRepository) CreateVerification(email string) error {
	verificationCode := GenerateVerificationCode(6)

	_, err := v.DB.Exec("INSERT INTO Verifications (email, verification_code, verified) VALUES (?, ?, ?)", email, verificationCode, false)
	if err != nil {
		return err
	}

	env_error := godotenv.Load("/workspace/.env")
	if env_error != nil {
		return env_error
	}

	var server = os.Getenv("SERVER_ADDRESS")

	encodedMail := b64.RawURLEncoding.EncodeToString([]byte(email))
	encodedCode := b64.RawURLEncoding.EncodeToString([]byte(verificationCode))

	sendMailError := util.SendEmail(email, "Verification Code", "<a href='"+server+"/verify?email="+encodedMail+"&code="+encodedCode+"'>Click here to verify your email</a>")
	if sendMailError != nil {
		return sendMailError
	}

	return nil
}

func (v *VerificationRepository) Verify(email string, verificationCode string) (string, error) {
	log.Printf("Verifying email %s with code %s", email, verificationCode)

	result, select_err := v.DB.Query("SELECT * FROM Verifications WHERE email = ? AND verification_code = ? AND verified = ?", email, verificationCode, false)

	if select_err != nil {
		log.Fatalf(select_err.Error())
		return "system", select_err
	}

	defer result.Close()

	if !result.Next() {
		return "user", errors.New("invalid verification code")
	}

	_, err := v.DB.Exec("UPDATE Verifications SET verified = ? WHERE email = ? AND verification_code = ?", true, email, verificationCode)
	if err != nil {
		log.Fatalf(err.Error())
		return "system", err
	}

	return "", nil
}
