package database

import (
	"errors"
	"time"

	"github.com/egomes/sharing-things/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection(env *env.Env) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(env.DatabaseURL), &gorm.Config{
		TranslateError: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return nil, errors.Join(err, errors.New("failed to connect database"))
	}

	config, err := db.DB()

	if err != nil {
		return nil, errors.Join(err, errors.New("failed to get database config"))
	}

	config.SetMaxIdleConns(env.DatabaseIddleConnections)
	config.SetMaxOpenConns(env.DatabaseMaxConnections)
	config.SetConnMaxLifetime(time.Minute * time.Duration(env.DatabaseMaxConnectionsLifetime))
	config.SetConnMaxIdleTime(time.Minute * time.Duration(env.DatabaseMaxIdleTime))

	return db, nil
}
