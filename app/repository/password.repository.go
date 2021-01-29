package repository

import (
	"app/model"
)

// GetPassword ...
func GetPassword(ID int) (string,error) {
	query := `SELECT passhash
	FROM users
	WHERE id = ?;`
	var pass string
	if err := db.Get(&pass, query, ID); err != nil {
		return "", err
	}

	return pass, nil
}

// PasswordUpdate ...
func PasswordUpdate(pass *model.Password) error {

	query := `UPDATE users
	SET passhash = :newpassword
	WHERE id = :id;`

	tx := db.MustBegin()
	_, err := tx.NamedExec(query, pass)
	if err != nil {
		if re := tx.Rollback(); re != nil {
			return re
		}
		return err
	}
	if re := tx.Commit(); re != nil {
		return re
	}
	return nil
}