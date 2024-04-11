package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	PgUser    string
	PgPass    string
	PgHost    string
	PgPort    uint16
	PgDb      string
	PgSSLMode string
}

type ServerConfig struct {
	HTTPPort       string
	ServerEndpoint string
}

func GetDBConfig() (DBConfig, error) {
	pgPort, err := strconv.ParseInt(getVarFromEnv("PGPORT"), 0, 16)
	if err != nil {
		return DBConfig{}, err
	}

	return DBConfig{
		PgUser:    getVarFromEnv("PGUSER"),
		PgPass:    getVarFromEnv("PGPASSWORD"),
		PgHost:    getVarFromEnv("PGHOST"),
		PgPort:    uint16(pgPort),
		PgDb:      getVarFromEnv("PGDATABASE"),
		PgSSLMode: getVarFromEnv("PGSSLMODE"),
	}, nil
}

func GetServerConfig() ServerConfig {
	return ServerConfig{
		HTTPPort:       ":" + getVarFromEnv("HTTP_PORT"),
		ServerEndpoint: getVarFromEnv("SERVER_ENDPOINT"),
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func getVarFromEnv(varName string) string {
	return os.Getenv(varName)
}
