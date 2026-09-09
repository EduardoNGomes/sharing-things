package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"uuid"

	"github.com/egomes/schedule/internal/domain/entities"
	"github.com/egomes/schedule/internal/env"
	"github.com/egomes/schedule/internal/infra/cryptography"
	"github.com/egomes/schedule/internal/infra/database"
	gormrepositories "github.com/egomes/schedule/internal/infra/gorm-repositories"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func DatabaseTestConfig(t *testing.T, schema string) *env.Env {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate test helper")
	}

	projectRoot := filepath.Clean(
		filepath.Join(filepath.Dir(filename), "../../.."),
	)

	envFile := filepath.Join(projectRoot, ".env")

	if err := godotenv.Load(envFile); err != nil {
		t.Fatalf("load %s: %v", envFile, err)
	}
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		t.Fatal("DATABASE_URL is required for E2E tests")
	}

	databaseURLTest := fmt.Sprintf(
		"%s&search_path=%s",
		databaseURL,
		schema,
	)
	t.Setenv("DATABASE_URL", databaseURLTest)

	env := &env.Env{
		DatabaseURL:                    databaseURLTest,
		DatabaseMaxConnections:         2,
		DatabaseIddleConnections:       2,
		DatabaseMaxConnectionsLifetime: 5,
		DatabaseMaxIdleTime:            5,
	}

	return env
}

func ApplyMigrations(t *testing.T, databaseURL string, schema string) {
	t.Helper()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve project root: caller information is unavailable")
	}

	projectRoot := filepath.Clean(filepath.Join(filepath.Dir(filename), "../../.."))
	cmd := exec.Command(
		"make",
		"migrate-up-tests",
		"DATABASE_URL="+databaseURL,
		"REVISIONS_SCHEMA="+schema,
	)
	cmd.Dir = projectRoot

	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("apply database migrations: %v\n%s", err, output)
	}
}

func CloseDatabase(
	t *testing.T,
	schema string,
	database *sql.DB,
) {
	t.Helper()

	t.Cleanup(func() {

		query := fmt.Sprintf(
			`DROP SCHEMA IF EXISTS "%s" CASCADE`,
			strings.ReplaceAll(schema, `"`, `""`),
		)

		if _, err := database.Exec(query); err != nil {
			t.Errorf("drop test schema: %v", err)
		}

		if err := database.Close(); err != nil {
			t.Errorf("close test database: %v", err)
		}
	})
}

func CreateFakeUserE2E(t *testing.T, user struct {
	Name     string
	Email    string
	Password string
}, database *gorm.DB) {
	t.Helper()
	ctx := context.Background()

	db := gormrepositories.GormUserRepository{DB: database}

	encryptedPassword, err := cryptography.BCryptEncrypter{}.Encrypt(user.Password)

	if err != nil {
		t.Fatal(err)
	}

	userDB := entities.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: encryptedPassword,
	}

	if err := db.CreateUser(ctx, userDB); err != nil {
		t.Fatal(err)
	}
}

type databaseTestUtils struct {
	Database *gorm.DB
	Envs     *env.Env
}

func CreataDatabaseTest(t *testing.T) databaseTestUtils {
	t.Helper()
	schema := "test_" + strings.ReplaceAll(
		uuid.New().String(),
		"-",
		"",
	)
	envTest := DatabaseTestConfig(t, schema)

	ApplyMigrations(t, envTest.DatabaseURL, schema)

	database, err := database.DatabaseConnection(envTest)

	if err != nil {
		t.Fatal(err)
	}

	dbConn, err := database.DB()
	if err != nil {
		t.Fatal(err)
	}

	CloseDatabase(t, schema, dbConn)

	return databaseTestUtils{
		Database: database,
		Envs:     envTest,
	}
}
