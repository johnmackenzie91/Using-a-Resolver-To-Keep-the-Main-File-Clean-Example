package db

import (
	"fmt"
	"os"
)

type Config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

func GetConfig() Config {
	return Config{
		dbUser: getMandidatoryEnvVar("DB_USER"),
		dbPass: getMandidatoryEnvVar("DB_PASS"),
		dbHost: getMandidatoryEnvVar("DB_HOST"),
		dbPort: getMandidatoryEnvVar("DB_PORT"),
		dbName: getMandidatoryEnvVar("DB_NAME"),
	}
}

func (c Config) GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true", c.dbUser, c.dbPass, c.dbHost, c.dbPort, c.dbName)
}

func getMandidatoryEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("mandidatory env var not set %s", key))
	}
	return val
}
