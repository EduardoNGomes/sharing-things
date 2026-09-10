package user

import (
	"context"

	"github.com/egomes/sharing-things/internal/domain/cryptography"
	domainErrors "github.com/egomes/sharing-things/internal/domain/errors"
	"github.com/egomes/sharing-things/internal/domain/repositories"
)

type LoginServiceDTO struct {
	Email    string
	Name     string
	Password string
}

type LoginService struct {
	encrypter      cryptography.Encrypter
	token          cryptography.Token
	userRepository repositories.UserRepository
}

type LoginServiceResponse struct {
	Token string
}

func NewLoginService(encrypter cryptography.Encrypter, token cryptography.Token, userRepository repositories.UserRepository) *LoginService {
	return &LoginService{
		encrypter:      encrypter,
		userRepository: userRepository,
		token:          token,
	}
}

func (c *LoginService) Handler(ctx context.Context, dto LoginServiceDTO) (*LoginServiceResponse, error) {
	user, err := c.userRepository.GetUserByEmail(ctx, dto.Email)

	if err != nil {
		switch err {
		case domainErrors.UserNotFoundError:
			return nil, domainErrors.UserNotFoundError
		default:
			return nil, err
		}
	}

	validatePasswword := c.encrypter.Comparer(dto.Password, user.Password)

	if validatePasswword == false {
		return nil, domainErrors.InvalidPasswordError
	}

	if user.GetUUID() == nil {
		return nil, domainErrors.InvalidUUIDError
	}

	token, err := c.token.CreateToken(user.GetUUID().String())

	if err != nil {
		return nil, err
	}

	return &LoginServiceResponse{
		Token: token,
	}, nil

}
