package cryptography

type Token interface {
	CreateToken(userID string) (string, error)
	ValidateToken(token string) (bool, error)
	GetUserUUID(tokenString string) (string, error)
}
