package main

import (
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/egomes/schedule/internal/env"
	"github.com/egomes/schedule/internal/infra/http/routes"
	"github.com/egomes/schedule/internal/infra/http/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	verifiedEnvs, err := env.New()

	if err != nil {
		log.Fatal(err)
	}

	mux := routes.New()

	if err := server.StartServer(verifiedEnvs, mux); err != nil {
		log.Fatal(err)
	}
}
