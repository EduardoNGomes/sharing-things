package main

import (
	"log"

	"github.com/egomes/schedule/internal/env"
	"github.com/egomes/schedule/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	verifiedEnvs, err := env.New()

	if err != nil {
		log.Fatal(err)
	}

	if err := server.StartServer(verifiedEnvs); err != nil {
		log.Fatal(err)
	}
}
