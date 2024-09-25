// config_test.go
package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func createEnvFile(path string, content string) error {
	f, err := os.Create(path + "/app.env")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

func deleteEnvFile(path string) error {
	return os.Remove(path + "/app.env")
}

func TestLoadConfig(t *testing.T) {
	// Set up test environment
	tempDir := t.TempDir()

	envContent := `
		DB_DRIVER=postgres
		DB_SOURCE=postgres://user:password@localhost:5432/dbname?sslmode=disable
`

	err := createEnvFile(tempDir, envContent)
	require.NoError(t, err)

	t.Run("ValidConfig", func(t *testing.T) {
		config, err := LoadConfig(tempDir)
		require.NoError(t, err)
		require.Equal(t, "postgres", config.DBDriver)
		require.Equal(t, "postgres://user:password@localhost:5432/dbname?sslmode=disable", config.DBSource)
	})

	t.Run("MissingConfigFile", func(t *testing.T) {
		err := deleteEnvFile(tempDir)
		require.NoError(t, err)

		_, err = LoadConfig(tempDir)
		require.Error(t, err)
	})

	t.Run("InvalidConfigFormat", func(t *testing.T) {
		invalidEnvContent := `
			INVALID_CONFIG_LINE
			DB_SOURCE = postgres://user:password@localhost:5432/dbname?sslmode=disable
`
		err := createEnvFile(tempDir, invalidEnvContent)
		require.NoError(t, err)

		_, err = LoadConfig(tempDir)
		require.Error(t, err)
	})
}
