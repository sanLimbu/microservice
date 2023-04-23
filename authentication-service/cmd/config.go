package cmd

import (
	"authentication/data"
	"database/sql"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
	users  data.UserInterface
}

func NewConfig(u data.UserInterface) *Config {
	return &Config{
		users: u,
	}
}
