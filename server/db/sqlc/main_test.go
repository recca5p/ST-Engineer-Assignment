package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDbContext *sql.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Define PostgreSQL container request
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine", // Use a lightweight PostgreSQL image
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpassword",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(5 * time.Minute),
	}

	// Start the container
	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal("Failed to start container:", err)
	}
	defer func() {
		_ = postgresContainer.Terminate(ctx)
	}()

	// Get the host and port of the running container
	host, _ := postgresContainer.Host(ctx)
	port, _ := postgresContainer.MappedPort(ctx, "5432")

	// Create the connection string
	dsn := fmt.Sprintf("postgres://testuser:testpassword@%s:%s/testdb?sslmode=disable", host, port.Port())

	// Open a connection to the database
	testDbContext, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the test database:", err)
	}

	// Ensure the connection is ready
	if err = testDbContext.Ping(); err != nil {
		log.Fatal("Failed to ping the test database:", err)
	}

	// Run migrations
	if err := RunMigrations(testDbContext); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize queries
	testQueries = New(testDbContext)

	// Run tests
	code := m.Run()

	// Close the database connection
	if err := testDbContext.Close(); err != nil {
		log.Fatal("Failed to close the test database:", err)
	}

	os.Exit(code)
}

func RunMigrations(db *sql.DB) error {
	migrationsDir := "../migrations/"

	// Get all .up.sql files
	migrationFiles, err := filepath.Glob(path.Join(migrationsDir, "*.up.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration files: %w", err)
	}

	// Sort the migration files to ensure they're executed in order
	sort.Strings(migrationFiles)

	for _, file := range migrationFiles {
		sqlContent, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		// Split the SQL content into individual statements
		statements := strings.Split(string(sqlContent), ";")

		// Execute each statement
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}
			_, err := db.Exec(stmt)
			if err != nil {
				return fmt.Errorf("failed to execute migration statement from %s: %w", file, err)
			}
		}

		fmt.Printf("Executed migration: %s\n", path.Base(file))
	}

	return nil
}
