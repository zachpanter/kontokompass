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

	portString := os.Getenv("DB_PORT_LOCAL")
	portInt, _ := strconv.Atoi(portString)
	config := &Config{
		DBPort:   portInt,
		DBHost:   os.Getenv("DB_HOST_LOCAL"),
		DBPass:   os.Getenv("DB_PASS_LOCAL"),
		DBUser:   os.Getenv("DB_USER_LOCAL"),
		DBSchema: os.Getenv("DB_NAME_LOCAL"),
	}
	return config
}
