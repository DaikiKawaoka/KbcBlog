package repository

import (
	"github.com/jmoiron/sqlx"
	"time"
)

var db *sqlx.DB
var jst,_ = time.LoadLocation("Asia/Tokyo")

// SetDB ...
func SetDB(d *sqlx.DB) {
	db = d
}