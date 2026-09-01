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
