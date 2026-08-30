package user

import (
	"context"
	"errors"

	"github.com/egomes/schedule/internal/domain/cryptography"
	"github.com/egomes/schedule/internal/domain/entities"
	domainErrors "github.com/egomes/schedule/internal/domain/errors"
	"github.com/egomes/schedule/internal/domain/repositories"
)

type CreateUserServiceDTO struct {
	Email    string
	Name     string
	Password string
}

type CreateUserService struct {
	encrypter      cryptography.Encrypter
	userRepository repositories.UserRepository
}

func NewCreateUserService(encrypter cryptography.Encrypter, userRepository repositories.UserRepository) *CreateUserService {
	return &CreateUserService{
		encrypter:      encrypter,
		userRepository: userRepository,
	}
}

func (c *CreateUserService) Handler(ctx context.Context, dto CreateUserServiceDTO) error {
	user, err := c.userRepository.GetUserByEmail(ctx, dto.Email)

	if err != nil && !errors.Is(err, domainErrors.UserNotFoundError) {
		return err
	}

	if user != nil {
		return domainErrors.UserAlreadyExistsError
	}

	encryptedPassword, err := c.encrypter.Encrypt(dto.Password)

	if err != nil {
		return err
	}

	newUser := entities.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: encryptedPassword,
	}

	if err := c.userRepository.CreateUser(ctx, newUser); err != nil {
		if errors.Is(err, domainErrors.UserAlreadyExistsError) {
			return domainErrors.UserAlreadyExistsError
		}
		return err
	}

	return nil

}
