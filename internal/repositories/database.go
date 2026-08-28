package repositories

import (
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection(databaseUrl string) error {
	_, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return errors.Join(err, errors.New("failed to connect database"))
	}

	return nil
}
