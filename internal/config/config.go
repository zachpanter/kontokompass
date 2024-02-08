package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBPort   int
	DBHost   string
	DBPass   string
	DBUser   string
	DBSchema string
}

func NewConfig() *Config {

	portString := os.Getenv("LC_DB_PORT")
	portInt, _ := strconv.Atoi(portString)
	config := &Config{
		DBPort:   portInt,
		DBHost:   os.Getenv("LC_DB_HOST"),
		DBPass:   os.Getenv("LC_DB_PASS"),
		DBUser:   os.Getenv("LC_DB_USER"),
		DBSchema: os.Getenv("LC_DB_SCHEMA"),
	}
	return config
}
