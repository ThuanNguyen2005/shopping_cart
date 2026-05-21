package config

import (
	"fmt"
	"os"

	"23.usersapi/internal/utils"
)

type databaseConfig struct {
	usernameDB string
	passwordDB string
	hostDB     string
	portDB     string
	nameDB     string
	sslMode    string
}

type Config struct {
	Port     string
	dbConfig databaseConfig
}

func NewConfig() *Config {
	return &Config{
		Port: os.Getenv("PORTSERVER"),
		dbConfig: databaseConfig{
			usernameDB: utils.Getenv("USERNAME_DB", "postgres"),
			passwordDB: utils.Getenv("PASSWORD_DB", "postgres"),
			hostDB:     utils.Getenv("HOST", "localhost"),
			nameDB:     utils.Getenv("NAME_DB", "my_app"),
			portDB:     utils.Getenv("PORT", "5432"),
			sslMode:    utils.Getenv("SSLMODE", "disable"),
		},
	}
}

func (c *Config) DNS() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", c.dbConfig.usernameDB, c.dbConfig.passwordDB, c.dbConfig.hostDB, c.dbConfig.portDB, c.dbConfig.nameDB, c.dbConfig.sslMode)
}
