package repositories

import (
	"context"

	"github.com/egomes/sharing-things/internal/domain/entities"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	CreateUser(ctx context.Context, user entities.User) error
}
