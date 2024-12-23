package settings_test

import (
	"os"
	"testing"

	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultENV     = ".env"
	expectedDSN    = "postgres://test_user:test_pass@localhost:5432/test_db?sslmode=disable"
	expectedXToken = "1234567890"
)

func setupTestEnv() {
	os.Setenv("DB_SQL_PROTO", "postgres")
	os.Setenv("DB_SQL_HOST", "localhost")
	os.Setenv("DB_SQL_PORT", "5432")
	os.Setenv("DB_SQL_USERNAME", "test_user")
	os.Setenv("DB_SQL_PASSWORD", "test_pass")
	os.Setenv("DB_SQL_DATABASE", "test_db")
	os.Setenv("DB_SQL_SSL", "false")
	os.Setenv("XTOKEN", expectedXToken)
}

func cleanupTestEnv() {
	os.Unsetenv("DB_SQL_PROTO")
	os.Unsetenv("DB_SQL_HOST")
	os.Unsetenv("DB_SQL_PORT")
	os.Unsetenv("DB_SQL_USERNAME")
	os.Unsetenv("DB_SQL_PASSWORD")
	os.Unsetenv("DB_SQL_DATABASE")
	os.Unsetenv("DB_SQL_SSL")
	os.Unsetenv("XTOKEN")
}

func TestLoadEnv(t *testing.T) {
	viper.Reset() // Reset viper state

	// Simulate a valid file path
	err := os.WriteFile(defaultENV, []byte(`
DB_SQL_PROTO=postgres
DB_SQL_HOST=localhost
DB_SQL_PORT=5432
DB_SQL_USERNAME=test_user
DB_SQL_PASSWORD=test_pass
DB_SQL_DATABASE=test_db
DB_SQL_SSL=false
XTOKEN=1234567890
`), 0o644)
	defer os.Remove(defaultENV) // Clean up after the test

	assert.NoError(t, err, "Error writing .env file")
	err = settings.LoadEnv(defaultENV)
	assert.NoError(t, err, "LoadEnv should not return an error")

	// Test invalid file path
	err = settings.LoadEnv("invalid.env")
	assert.Error(t, err, "LoadEnv should return an error for invalid file paths")
}

func TestInitConfig(t *testing.T) {
	// show files in current directory
	files, err := os.ReadDir("../../../") // "." represents the current directory
	if err != nil {
		t.Error("Error reading directory:", err)
		return
	}

	t.Log("Current directory contents:")

	for _, file := range files {
		t.Logf(" - %s (directory: %t)\n", file.Name(), file.IsDir())
	}

	viper.Reset() // Reset viper state

	// Simulate a valid file path
	err = os.WriteFile(defaultENV, []byte(`
DB_SQL_PROTO=postgres
DB_SQL_HOST=localhost
DB_SQL_PORT=5432
DB_SQL_USERNAME=test_user
DB_SQL_PASSWORD=test_pass
DB_SQL_DATABASE=test_db
DB_SQL_SSL=false
XTOKEN=1234567890
`), 0o644)
	require.NoError(t, err)
	defer os.Remove(defaultENV) // Clean up after the test

	cfg := settings.InitConfig(defaultENV, settings.WithDSN())
	assert.NotNil(t, cfg, "Config should not be nil")
	setupTestEnv()

	defer cleanupTestEnv()

	assert.Equal(t, expectedDSN, cfg.DSN, "DSN does not match expected value")

	assert.Equal(t, "localhost", cfg.DB.Host, "Host should match environment variable")
	assert.Equal(t, "5432", cfg.DB.Port, "Port should match environment variable")
}

func TestInitConfig_PanicOnInvalidFile(t *testing.T) {
	viper.Reset() // Reset viper state to avoid interference

	invalidFilePath := "non_existent.env"

	// Ensure the InitConfig function panics on invalid file paths
	require.Panics(t, func() {
		settings.InitConfig(invalidFilePath)
	}, "InitConfig should panic when provided with an invalid file path")
}

func TestLoadConfig(t *testing.T) {
	setupTestEnv()

	defer cleanupTestEnv()

	viper.Reset() // Reset viper state

	// Simulate a valid file path
	err := os.WriteFile(defaultENV, []byte(`
DB_SQL_PROTO=postgres
DB_SQL_HOST=localhost
DB_SQL_PORT=5432
DB_SQL_USERNAME=test_user
DB_SQL_PASSWORD=test_pass
DB_SQL_DATABASE=test_db
DB_SQL_SSL=false
XTOKEN=1234567890
`), 0o644)
	require.NoError(t, err)

	defer os.Remove(defaultENV) // Clean up after the test
	cfg := settings.LoadConfig(defaultENV)
	assert.NotNil(t, cfg, "Config should not be nil")

	assert.Equal(t, expectedDSN, cfg.DSN, "DSN does not match expected value")

	// Ensure singleton is used
	cfgAgain := settings.LoadConfig(defaultENV)
	assert.Equal(t, cfg, cfgAgain, "Singleton instance should remain consistent")
}

func TestReload(t *testing.T) {
	setupTestEnv()

	defer cleanupTestEnv()

	cfg := &settings.Settings{}
	cfg.Reload(settings.WithDSN(), settings.WithXToken())

	assert.Equal(t, expectedDSN, cfg.DSN, "DSN does not match expected value")
	assert.Equal(t, expectedXToken, cfg.XToken, "X-Token does not match expected value")
}

func TestWithDSN(t *testing.T) {
	setupTestEnv()

	defer cleanupTestEnv()

	cfg := &settings.Settings{}
	withDSN := settings.WithDSN()
	withDSN(cfg)

	assert.Equal(t, expectedDSN, cfg.DSN, "DSN does not match expected value")

	expectedDB := settings.DB{
		Proto:    "postgres",
		Host:     "localhost",
		Port:     "5432",
		Username: "test_user",
		Password: "test_pass",
		Database: "test_db",
	}
	assert.Equal(t, expectedDB, cfg.DB, "Database configuration does not match expected values")
}

func TestWithXToken(t *testing.T) {
	setupTestEnv()

	defer cleanupTestEnv()

	cfg := &settings.Settings{}
	withXToken := settings.WithXToken()
	withXToken(cfg)

	assert.Equal(t, expectedXToken, cfg.XToken, "X-Token does not match expected value")
}
