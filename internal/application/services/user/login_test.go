package user

import (
	"context"
	"math"
	"testing"
	"uuid"

	userEntities "github.com/egomes/sharing-things/internal/domain/entities"
	"github.com/egomes/sharing-things/internal/test/cryptography"
	inmemoryrepositories "github.com/egomes/sharing-things/internal/test/in-memory-repositories"
)

func TestLoginService(t *testing.T) {
	t.Run("[SERVICE] Login user", func(t *testing.T) {

		encrypter := cryptography.MockCryptographer{}
		inMemoryUserRepository := inmemoryrepositories.NewInMemoryUserRepository()
		token := cryptography.MockToken{}

		createFakeUser(t, "test@test.com", "test", encrypter, inMemoryUserRepository)

		loginUserService := NewLoginService(encrypter, token, inMemoryUserRepository)

		loginUserServiceDTO := LoginServiceDTO{
			Email:    "test@test.com",
			Password: "test",
		}

		ctx := context.Background()

		response, err := loginUserService.Handler(ctx, loginUserServiceDTO)

		if response == nil || err != nil {
			t.Errorf("Should not return error, got %v", response)
		}

		user, err := loginUserService.userRepository.GetUserByEmail(ctx, "test@test.com")

		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if user.GetUUID() == nil {
			t.Errorf("Should return user with uuid, got %v", user.GetUUID())
		}

		verifyToken, err := token.ValidateToken(response.Token)

		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if verifyToken == false {
			t.Errorf("Should return user with uuid, got %s \n user uuid: %s", response.Token, user.GetUUID())
		}

		if user.Email != "test@test.com" {
			t.Errorf("Should return user with email test@test.com, got %v", user.Email)
		}
	})
}

func createFakeUser(t *testing.T, email, password string, encrypter cryptography.MockCryptographer, inMemoryUserRepository *inmemoryrepositories.InMemoryUserRepository) {
	t.Helper()
	encrypterPassword, _ := encrypter.Encrypt(password)

	internalID := uint(math.MaxUint32)
	uuid := uuid.NewV7()

	inMemoryUserRepository.CreateUser(context.Background(), userEntities.User{
		Email:      email,
		Name:       "Test",
		Password:   encrypterPassword,
		InternalID: &internalID,
		UUID:       &uuid,
	})
}
