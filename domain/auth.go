package domain

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AuthUseCase interface {
	GenerateToken(username, password string) (string, error)
	ValidateToken(token string) error
}