package cryptography

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (e JWT) CreateToken(userUUID string) (string, error) {

	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(e.secret)

	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "sharing-things",
			"sub": userUUID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	s, err := t.SignedString(key)

	if err != nil {
		return "", err
	}

	return s, nil
}

func (e JWT) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(e.secret), nil
	})

	if err != nil {
		return false, err
	}

	if token.Valid == false {
		return false, fmt.Errorf("token is invalid")
	}

	return true, nil
}

func (e JWT) GetUserUUID(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(e.secret), nil
	})

	if err != nil {
		return "", err
	}

	if token.Valid == false {
		return "", fmt.Errorf("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", fmt.Errorf("token is invalid")
	}

	if _, ok := claims["sub"]; !ok {
		return "", fmt.Errorf("token is invalid")
	}

	return claims["sub"].(string), nil
}
