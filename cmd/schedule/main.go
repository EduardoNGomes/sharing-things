package main

import (
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/egomes/schedule/internal/env"
	"github.com/egomes/schedule/internal/infra/database"
	"github.com/egomes/schedule/internal/infra/factories"
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

	dbConfig, err := databaseConnection.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer dbConfig.Close()

	services := initServices(databaseConnection, verifiedEnvs)

	r := routes.NewRoutes(services)

	mux := r.New()

	if err := server.StartServer(verifiedEnvs, mux); err != nil {
		log.Fatal(err)
	}
}

func initServices(db *gorm.DB, env *env.Env) *routes.Services {
	createUserService := factories.CreateUserServiceFactory(db)
	loginService := factories.LoginServiceFactory(db, env)

	return &routes.Services{
		CreateUserService: createUserService,
		LoginService:      loginService,
	}
}
