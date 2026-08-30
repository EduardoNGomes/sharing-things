package main

import (
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/egomes/schedule/internal/application/services/user"
	"github.com/egomes/schedule/internal/env"
	"github.com/egomes/schedule/internal/infra/cryptography"
	"github.com/egomes/schedule/internal/infra/database"
	gormrepositories "github.com/egomes/schedule/internal/infra/gorm-repositories"
	"github.com/egomes/schedule/internal/infra/http/routes"
	"github.com/egomes/schedule/internal/infra/http/server"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func main() {
	verifiedEnvs, err := env.New()

	if err != nil {
		log.Fatal(err)
	}

	databaseConnection, err := database.DatabaseConnection(verifiedEnvs)

	if err != nil {
		log.Fatal(err)
	}

	dbConifg, err := databaseConnection.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer dbConifg.Close()

	if err != nil {
		log.Fatal(err)
	}

	services := initServices(databaseConnection)

	r := routes.NewRoutes(services.createUserService)

	mux := r.New()

	if err := server.StartServer(verifiedEnvs, mux); err != nil {
		log.Fatal(err)
	}
}

type Services struct {
	createUserService *user.CreateUserService
}

func initServices(db *gorm.DB) *Services {
	userRepository := gormrepositories.GormUserRepository{DB: db}

	createUserService := user.NewCreateUserService(cryptography.BCryptEncrypter{},
		userRepository,
	)

	return &Services{
		createUserService: createUserService,
	}
}
