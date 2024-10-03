package handlers

import (
	"database/sql"
)

var DB *sql.DB

func SetDB(database *sql.DB) {
	DB = database
}
