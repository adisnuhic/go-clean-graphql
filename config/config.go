package config

import (
	"os"
	"strconv"

	"github.com/adisnuhic/go-graphql/pkg/utils"
	"github.com/subosito/gotenv"
)

// DBConnection connection to a database
type DBConnection struct {
	DBDialect         string
	DBConnection      string
	DbMaxIdleConns    int
	DbMaxOpenConns    int
	DbConnMaxLifetime int
	DbLogging         bool
}

// JWT configuration
type JWTConf struct {
	Secret string
}

// Env conf
type Env struct {
	Env  string
	Port string
	Url  string
}

// AppConfig application configuration
type AppConfig struct {
	DBConnections map[string]DBConnection
	JWTConf
	Env
}

// Load app configuration
func Load() *AppConfig {
	gotenv.Load()

	// defaults
	maxIdleConnections := 10
	maxOpenConnections := 100
	maxConnectionsLifetime := 20

	port := "8080"
	URL := "http://localhost"

	if utils.GetEnv("PORT") == "" {
		os.Setenv("PORT", port)
	}
	if utils.GetEnv("URL") == "" {
		os.Setenv("URL", URL)
	}

	// customs
	if utils.GetEnv("DB_MAX_IDLE_CONNECTIONS") != "" {
		maxIdleConnections, _ = strconv.Atoi(utils.GetEnv("DB_MAX_IDLE_CONNECTIONS"))
	}

	if utils.GetEnv("DB_MAX_OPEN_CONNECTIONS") != "" {
		maxOpenConnections, _ = strconv.Atoi(utils.GetEnv("DB_MAX_OPEN_CONNECTIONS"))
	}

	if utils.GetEnv("DB_CONNECTION_MAX_LIFETIME_MINUTES") != "" {
		maxConnectionsLifetime, _ = strconv.Atoi(utils.GetEnv("DB_CONNECTION_MAX_LIFETIME_MINUTES"))
	}

	return &AppConfig{
		DBConnections: map[string]DBConnection{
			"development": {
				DBDialect:         "mysql",
				DBConnection:      utils.GetEnv("DB_DEV_CONNECTION"),
				DbMaxIdleConns:    maxIdleConnections,
				DbMaxOpenConns:    maxOpenConnections,
				DbConnMaxLifetime: maxConnectionsLifetime,
				DbLogging:         true,
			},
			"production": {
				DBDialect:         "mysql",
				DBConnection:      utils.GetEnv("DB_PROD_CONNECTION"),
				DbMaxIdleConns:    maxIdleConnections,
				DbMaxOpenConns:    maxOpenConnections,
				DbConnMaxLifetime: maxConnectionsLifetime,
				DbLogging:         true,
			},
		},
		JWTConf: JWTConf{
			Secret: utils.MustGetEnvMin("JWT_SECRET", 32),
		},
		Env: Env{
			Env:  utils.MustGetEnv("ENV"),
			Port: port,
			Url:  URL,
		},
	}
}
