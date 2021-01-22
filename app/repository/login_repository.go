package repository

import (
	// "database/sql"
	// "math"
	"app/model"
)

// GetLoginPassHash ...
func GetLoginPassHash(mail string) (*model.LoginUser , error) {
	var loginUser model.LoginUser

	query := `SELECT id,mail,passhash
	FROM users
	WHERE mail = ?;`

	if err := db.Get(&loginUser, query, mail); err != nil {
		return nil ,err
	}

	return &loginUser, nil
}

// GetGuestUser ...
func GetGuestUser() (*model.LoginUser , error) {
	var loginUser model.LoginUser

	query := `SELECT id
	FROM users
	WHERE mail = "kbc920437694@stu.kawahara.ac.jp";`

	if err := db.Get(&loginUser, query); err != nil {
		return nil ,err
	}

	return &loginUser, nil
}