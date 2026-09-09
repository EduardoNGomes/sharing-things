package user_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egomes/schedule/internal/application/services/user"
	"github.com/egomes/schedule/internal/env"
	"github.com/egomes/schedule/internal/test/database"
	"gorm.io/gorm"

	"github.com/egomes/schedule/internal/infra/factories"
	"github.com/egomes/schedule/internal/infra/http/routes"
)

func TestSignInRoute(t *testing.T) {
	t.Run("[E2E]-[/SIGNIN] Should login user", func(t *testing.T) {
		databaseTest := database.CreataDatabaseTest(t)

		database.CreateFakeUserE2E(t, struct {
			Name     string
			Email    string
			Password string
		}{Name: "Alice", Email: "alice@example.com", Password: "secret123"}, databaseTest.Database)

		service := routes.Services{
			LoginService: loginServiceTest(t, databaseTest.Database, databaseTest.Envs),
		}

		handler := routes.NewRoutes(&service).New()

		server := httptest.NewServer(handler)

		t.Cleanup(server.Close)

		body := bytes.NewBufferString(`{
  "email": "alice@example.com",
  "password": "secret123"
}`)

		response, err := server.Client().Post(
			server.URL+"/signin",
			"application/json",
			body,
		)
		if err != nil {
			t.Fatal(err)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			responseBody, _ := io.ReadAll(response.Body)

			t.Fatalf(
				"expected status %d, got %d: %s",
				http.StatusCreated,
				response.StatusCode,
				responseBody,
			)
		}

		if response.Cookies()[0].Name != "token" {
			t.Fatalf("expected cookie token, got %s", response.Cookies()[0].Name)
		}
	})

	t.Run("[E2E]-[/SIGNIN]{INVALID PASSWORD} Shouln't login with invalid credentials", func(t *testing.T) {
		databaseTest := database.CreataDatabaseTest(t)

		database.CreateFakeUserE2E(t, struct {
			Name     string
			Email    string
			Password string
		}{Name: "Alice", Email: "alice@example.com", Password: "secret123"}, databaseTest.Database)

		service := routes.Services{
			LoginService: loginServiceTest(t, databaseTest.Database, databaseTest.Envs),
		}

		handler := routes.NewRoutes(&service).New()

		server := httptest.NewServer(handler)

		t.Cleanup(server.Close)

		body := bytes.NewBufferString(`{
  "email": "alice@example.com",
  "password": "secret1234"
}`)

		response, err := server.Client().Post(
			server.URL+"/signin",
			"application/json",
			body,
		)
		if err != nil {
			t.Fatal(err)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusConflict {
			responseBody, _ := io.ReadAll(response.Body)

			t.Fatalf(
				"expected status %d, got %d: %s",
				http.StatusCreated,
				response.StatusCode,
				responseBody,
			)
		}
	})

	t.Run("[E2E]-[/SIGNIN]{NO USER} Shouln't login with invalid credentials[NO USER]", func(t *testing.T) {
		databaseTest := database.CreataDatabaseTest(t)

		service := routes.Services{
			LoginService: loginServiceTest(t, databaseTest.Database, databaseTest.Envs),
		}

		handler := routes.NewRoutes(&service).New()

		server := httptest.NewServer(handler)

		t.Cleanup(server.Close)

		body := bytes.NewBufferString(`{
  "email": "alice@example.com",
  "password": "secret1234"
}`)

		response, err := server.Client().Post(
			server.URL+"/signin",
			"application/json",
			body,
		)
		if err != nil {
			t.Fatal(err)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusConflict {
			responseBody, _ := io.ReadAll(response.Body)

			t.Fatalf(
				"expected status %d, got %d: %s",
				http.StatusCreated,
				response.StatusCode,
				responseBody,
			)
		}
	})

}

func loginServiceTest(t *testing.T, dbConn *gorm.DB, env *env.Env) *user.LoginService {
	t.Helper()
	return factories.LoginServiceFactory(dbConn, env)
}
