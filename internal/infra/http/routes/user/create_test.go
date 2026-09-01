package user_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"uuid"

	databaseTest "github.com/egomes/schedule/internal/test/database"

	createUserService "github.com/egomes/schedule/internal/application/services/user"
	"github.com/egomes/schedule/internal/infra/cryptography"
	"github.com/egomes/schedule/internal/infra/database"
	gormrepositories "github.com/egomes/schedule/internal/infra/gorm-repositories"
	"github.com/egomes/schedule/internal/infra/http/routes"
)

func TestSignUpRoute(t *testing.T) {

	t.Run("[E2E]-[/SIGUP] Create user", func(t *testing.T) {
		schema := "test_" + strings.ReplaceAll(
			uuid.New().String(),
			"-",
			"",
		)
		envTest := databaseTest.DatabaseTestConfig(t, schema)

		databaseTest.ApplyMigrations(t, envTest.DatabaseURL, schema)

		database, err := database.DatabaseConnection(envTest)

		if err != nil {
			t.Fatal(err)
		}

		dbConn, err := database.DB()
		if err != nil {
			t.Fatal(err)
		}
		defer databaseTest.CloseDatabase(t, schema, dbConn)

		userRepository := gormrepositories.GormUserRepository{
			DB: database,
		}

		service := createUserService.NewCreateUserService(
			cryptography.BCryptEncrypter{},
			userRepository,
		)

		handler := routes.NewRoutes(service).New()
		server := httptest.NewServer(handler)

		defer t.Cleanup(server.Close)

		body := bytes.NewBufferString(`{
  "name": "Alice",
  "email": "alice@example.com",
  "password": "secret123"
}`)

		response, err := server.Client().Post(
			server.URL+"/signup",
			"application/json",
			body,
		)
		if err != nil {
			t.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusCreated {
			responseBody, _ := io.ReadAll(response.Body)

			t.Fatalf(
				"expected status %d, got %d: %s",
				http.StatusCreated,
				response.StatusCode,
				responseBody,
			)
		}
	})

	t.Run("[E2E]-[/SIGUP] Cannot create user with email already registered", func(t *testing.T) {
		schema := "test_" + strings.ReplaceAll(
			uuid.New().String(),
			"-",
			"",
		)
		envTest := databaseTest.DatabaseTestConfig(t, schema)

		databaseTest.ApplyMigrations(t, envTest.DatabaseURL, schema)

		database, err := database.DatabaseConnection(envTest)

		if err != nil {
			t.Fatal(err)
		}

		dbConn, err := database.DB()
		if err != nil {
			t.Fatal(err)
		}
		defer databaseTest.CloseDatabase(t, schema, dbConn)

		userRepository := gormrepositories.GormUserRepository{
			DB: database,
		}

		service := createUserService.NewCreateUserService(
			cryptography.BCryptEncrypter{},
			userRepository,
		)

		handler := routes.NewRoutes(service).New()
		server := httptest.NewServer(handler)

		defer t.Cleanup(server.Close)

		payload := []byte(`{
      "name": "Alice",
      "email": "alice@example.com",
      "password": "secret123"
  }`)
		if _, err := server.Client().Post(
			server.URL+"/signup",
			"application/json",
			bytes.NewBuffer(payload),
		); err != nil {
			t.Fatal(err)
		}

		response, err := server.Client().Post(
			server.URL+"/signup",
			"application/json",
			bytes.NewBuffer(payload),
		)
		if err != nil {
			t.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusConflict {
			responseBody, _ := io.ReadAll(response.Body)

			t.Fatalf(
				"expected status %d, got %d: %s",
				http.StatusConflict,
				response.StatusCode,
				responseBody,
			)
		}
	})

}
