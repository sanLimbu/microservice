package cmd

import (
	"authentication/data"
	"database/sql"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}
