package gormrepositories

import (
	"context"
	"errors"
	"time"
	"uuid"

	"github.com/egomes/schedule/internal/domain/entities"
	domainErrors "github.com/egomes/schedule/internal/domain/errors"
	"gorm.io/gorm"
)

type UserDB struct {
	UUID       uuid.UUID `gorm:"type:uuid;primary_key"`
	InternalID uint      `gorm:"uniqueIndex;not null;autoIncrement"`
	Name       string    `gorm:"type:varchar(50);not null"`
	Email      string    `gorm:"type:varchar(50);not null;uniqueIndex"`
	Password   string    `gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time `gorm:"default:now()"`
	UpdatedAt  time.Time `gorm:"default:now()"`
}

func (u *UserDB) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.NewV7()
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *UserDB) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *UserDB) TableName() string {
	return "users"
}

type GormUserRepository struct {
	DB *gorm.DB
}

func (r GormUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var userDB UserDB

	userDB, err := gorm.G[UserDB](r.DB).Where("email = ?", email).First(ctx)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.UserNotFoundError
		}

		return nil, err
	}

	user := entities.User{
		Email:    userDB.Email,
		Name:     userDB.Name,
		Password: userDB.Password,
		ID:       &userDB.InternalID,
		UUID:     &userDB.UUID,
	}

	return &user, nil
}

func (r GormUserRepository) CreateUser(ctx context.Context, user entities.User) error {
	newUser := UserDB{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	err := gorm.G[UserDB](r.DB).Create(ctx, &newUser)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domainErrors.UserAlreadyExistsError
		}

		return err
	}

	return nil
}
