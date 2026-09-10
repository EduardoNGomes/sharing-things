package user

import (
	"context"
	"testing"

	domainErrors "github.com/egomes/sharing-things/internal/domain/errors"
	"github.com/egomes/sharing-things/internal/test/cryptography"
	inmemoryrepositories "github.com/egomes/sharing-things/internal/test/in-memory-repositories"
)

func TestCreateUserService(t *testing.T) {

	t.Run("[SERVICE] Should create user a new user", func(t *testing.T) {
		createUserService := NewCreateUserService(cryptography.MockCryptographer{}, inmemoryrepositories.NewInMemoryUserRepository())

		createUserServiceDTO := CreateUserServiceDTO{
			Email:    "test@test.com",
			Name:     "Test",
			Password: "test",
		}

		ctx := context.Background()

		err := createUserService.Handler(ctx, createUserServiceDTO)

		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		user, err := createUserService.userRepository.GetUserByEmail(ctx, "test@test.com")

		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if user.Email != "test@test.com" {
			t.Errorf("Should return user with email test@test.com, got %v", user.Email)
		}
	})

	t.Run("[SERVICE] Should return error if user already exists", func(t *testing.T) {
		createUserService := NewCreateUserService(cryptography.MockCryptographer{}, inmemoryrepositories.NewInMemoryUserRepository())

		createUserServiceDTO := CreateUserServiceDTO{
			Email:    "test@test.com",
			Name:     "Test",
			Password: "test",
		}

		ctx := context.Background()

		err := createUserService.Handler(ctx, createUserServiceDTO)

		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		err = createUserService.Handler(ctx, createUserServiceDTO)

		if err != domainErrors.UserAlreadyExistsError {
			t.Errorf("Should return error UserAlreadyExistsError, got %v", err)
		}
	})

}
