package schema

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	TokenType   string `json:"token_type"`
	AuthToken   string `json:"auth_token"`
	GeneratedAt string `json:"generated_at"`
	ExpiresAt   string `json:"expires_at"`
}
