package api

import (
	b64 "encoding/base64"
	"encoding/json"
	"go-demo-api/internal/db"
	"net/http"
)

type APIUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserHandler struct {
	Repo db.UserRepositoryInterface
}

type VerificationHandler struct {
	VerificationRepository db.VerificationRepositoryInterface
}

// RegisterUser registers a new user in the system.
// @Summary Register a new user
// @Description Registers a new user with the provided username, email, and password.
// @Tags users
// @Accept json
// @Produce json
// @Param user body APIUser true "User to register"
// @Success 201 {string} string "Successfully registered the user"
// @Failure 400 {string} string "Invalid request parameters"
// @Failure 500 {string} string "Internal server error"
// @Router /users/register [post]
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user APIUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Repo.RegisterUser(user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// VerifyUser godoc
// @Summary Verify user
// @Description Verifies a user using base64 URL encoded email and verification code.
// @Tags verification
// @Accept  json
// @Produce  json
// @Param   email   query    string     true  "Base64 URL Encoded Email"
// @Param   code    query    string     true  "Base64 URL Encoded Verification Code"
// @Success 200  {string}  string  "User verified successfully"
// @Failure 400  {string}  string  "Invalid email or code"
// @Failure 500  {string}  string  "Verification failed"
// @Router /verify [get]
func (h *VerificationHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	email_byte, mail_decode_err := b64.RawURLEncoding.DecodeString(r.URL.Query().Get("email"))
	code_byte, code_decode_err := b64.RawURLEncoding.DecodeString(r.URL.Query().Get("code"))

	if mail_decode_err != nil || code_decode_err != nil {
		http.Error(w, "Invalid email or code", http.StatusBadRequest)
		return
	}

	email := string(email_byte)
	code := string(code_byte)

	error_type, err := h.VerificationRepository.Verify(email, code)
	if err != nil {
		if error_type == "system" {
			http.Error(w, "Verification failed", http.StatusInternalServerError)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User verified successfully"))
}
