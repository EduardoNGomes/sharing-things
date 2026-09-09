package cryptography

import (
	"fmt"
	"strings"
)

type MockCryptographer struct{}

func (e MockCryptographer) Encrypt(password string) (string, error) {
	return fmt.Sprintf("%s-encrypted", password), nil
}

func (e MockCryptographer) Comparer(password, hash string) bool {
	hashWithPrefix := strings.Split(hash, "-encrypted")

	return hashWithPrefix[0] == password
}

type MockToken struct{}

func (e MockToken) CreateToken(userUUID string) (string, error) {
	return fmt.Sprintf("token-%s", userUUID), nil
}

func (e MockToken) ValidateToken(token string) (bool, error) {
	tokenWithPrefix := strings.Split(token, "token-")

	if len(tokenWithPrefix) != 2 {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}

func (e MockToken) GetUserUUID(tokenString string) (string, error) {
	token, err := e.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	if token == false {
		return "", fmt.Errorf("invalid token")
	}

	tokenWithPrefix := strings.Split(tokenString, "token-")

	return tokenWithPrefix[1], nil
}
