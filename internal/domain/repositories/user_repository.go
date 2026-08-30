package repositories

import (
	"context"

	"github.com/egomes/schedule/internal/domain/entities"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	CreateUser(ctx context.Context, user entities.User) error
}
