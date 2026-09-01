package inmemoryrepositories

import (
	"context"
	"slices"
	"uuid"

	"github.com/egomes/schedule/internal/domain/entities"
	"github.com/egomes/schedule/internal/domain/errors"
)

type InMemoryUserRepository struct {
	DB []entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		DB: []entities.User{},
	}
}

func (r *InMemoryUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {

	index := slices.IndexFunc(r.DB, func(user entities.User) bool {
		return user.Email == email
	})

	if index == -1 {
		return nil, errors.UserNotFoundError
	}

	userDB := r.DB[index]

	user := entities.User{
		Email:      userDB.Email,
		Name:       userDB.Name,
		Password:   userDB.Password,
		InternalID: userDB.InternalID,
		UUID:       userDB.UUID,
	}

	return &user, nil
}

func (r *InMemoryUserRepository) CreateUser(ctx context.Context, user entities.User) error {
	ID := uint(len(r.DB))
	UUID := uuid.NewV7()

	newUser := entities.User{
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		InternalID: &ID,
		UUID:       &UUID,
	}

	r.DB = append(r.DB, newUser)

	return nil
}
