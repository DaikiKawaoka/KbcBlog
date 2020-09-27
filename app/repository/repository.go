package repository

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// SetDB ...
func SetDB(d *sqlx.DB) {
	db = d
}