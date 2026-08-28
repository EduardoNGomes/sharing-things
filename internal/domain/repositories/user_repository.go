package repositories

import (
	"github.com/egomes/schedule/internal/domain/entities"
)

type UserRepository interface {
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(user entities.User) error
}
