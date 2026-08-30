package cryptography

import "golang.org/x/crypto/bcrypt"

type BCryptEncrypter struct{}

func (e BCryptEncrypter) Encrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (e BCryptEncrypter) Comparer(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
