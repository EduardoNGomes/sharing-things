package factories

import (
	"github.com/egomes/schedule/internal/application/services/user"
	"github.com/egomes/schedule/internal/infra/cryptography"
	gormrepositories "github.com/egomes/schedule/internal/infra/gorm-repositories"
	"gorm.io/gorm"
)

func CreateUserServiceFactory(db *gorm.DB) *user.CreateUserService {
	userRepository := gormrepositories.GormUserRepository{DB: db}

	return user.NewCreateUserService(cryptography.BCryptEncrypter{},
		userRepository,
	)

}
