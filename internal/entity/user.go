package entity

import (
	"github.com/wordict/w-api/internal/core/errors"
	"net/mail"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequestDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (sr SignupRequest) Validate() error {
	_, err := mail.ParseAddress(sr.Email)
	if err != nil {
		return errors.InvalidEmail
	}
	return nil
}

type SignupResponse struct{}
