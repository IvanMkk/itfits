package schema

import (
	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountDetailsResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type GetAccountParams struct {
	Id uuid.UUID
}
