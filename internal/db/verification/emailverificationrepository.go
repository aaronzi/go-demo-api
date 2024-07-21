package db

import (
	b64 "encoding/base64"
	"go-demo-api/internal/util"
	"os"
)

type EmailVerificationRepository struct {
	VerificationRepository *VerificationRepository
}

func (e *EmailVerificationRepository) CreateVerification(email string) (string, error) {
	verificationCode, err := e.VerificationRepository.CreateVerification(email)
	if err != nil {
		return "", err
	}

	var server = os.Getenv("SERVER_ADDRESS")

	encodedMail := b64.RawURLEncoding.EncodeToString([]byte(email))
	encodedCode := b64.RawURLEncoding.EncodeToString([]byte(verificationCode))

	sendMailError := util.SendEmail(email, "Verification Code", "<a href='"+server+"/verify?email="+encodedMail+"&code="+encodedCode+"'>Click here to verify your email</a>")
	if sendMailError != nil {
		return "", sendMailError
	}

	return verificationCode, nil
}

func (e *EmailVerificationRepository) Verify(email string, verificationCode string) (string, error) {
	if errtype, err := e.VerificationRepository.Verify(email, verificationCode); err != nil {
		return errtype, err
	}
	return "", nil
}
