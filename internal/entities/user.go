package entities

import (
	"time"
	"uuid"

	"gorm.io/gorm"
)

type User struct {
	UUID       uuid.UUID `gorm:"type:uuid;primary_key"`
	InternalID uint
	Name       string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.NewV7()
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return nil
}
