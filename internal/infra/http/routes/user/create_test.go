package user_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"uuid"

	"github.com/egomes/schedule/internal/application/services/user"
	databaseTest "github.com/egomes/schedule/internal/test/database"

	"github.com/egomes/schedule/internal/infra/database"
	"github.com/egomes/schedule/internal/infra/factories"
	"github.com/egomes/schedule/internal/infra/http/routes"
)

func TestSignUpRoute(t *testing.T) {

	t.Run("[E2E]-[/SIGUP] Create user", func(t *testing.T) {
		service := createServiceTest(t)

		handler := routes.NewRoutes(service).New()
		server := httptest.NewServer(handler)

		t.Cleanup(server.Close)

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
		service := createServiceTest(t)

		handler := routes.NewRoutes(service).New()
		server := httptest.NewServer(handler)

		t.Cleanup(server.Close)

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

func createServiceTest(t *testing.T) *user.CreateUserService {
	t.Helper()
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

	return factories.CreateUserServiceFactory(database)

}
