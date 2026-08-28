package main

import (
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/egomes/schedule/internal/env"
	"github.com/egomes/schedule/internal/repositories"
	"github.com/egomes/schedule/internal/routes"
	"github.com/egomes/schedule/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	verifiedEnvs, err := env.New()

	if err != nil {
		log.Fatal(err)
	}

	if err := repositories.DatabaseConnection(verifiedEnvs.DatabaseUrl); err != nil {
		panic(err)
	}

	mux := routes.New()

	if err := server.StartServer(verifiedEnvs, mux); err != nil {
		log.Fatal(err)
	}
}
