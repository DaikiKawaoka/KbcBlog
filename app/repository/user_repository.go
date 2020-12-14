package repository

import (
	"database/sql"
	"app/model"
)

// GetMyUser ...
func GetMyUser(id int) (*model.User , error) {
	query := `SELECT id,mail,name,comment
	FROM users
	WHERE id = ?;`

	var user model.User
	if err := db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

// UserCreate ...
func UserCreate(user *model.CreateUser) (sql.Result, error) {

	//パスワードハッシュ化
	_, err := user.PasswordHash()
	if err != nil{
    return nil, err
	}

  query := `INSERT INTO users (mail,passhash,name)
	VALUES (:mail, :passhash, :name);`

  tx := db.MustBegin()
  res, err := tx.NamedExec(query, user)
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  tx.Commit()
  return res, nil
}

// UserGetByID ...
func UserGetByID(id int) (*model.User, error) {
	query := `SELECT id,mail,name,comment,github,website,languages
	FROM users
	WHERE id = ?`

	var user model.User
	if err := db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

// UserUpdate ...
func UserUpdate(user *model.User) error {

	query := `UPDATE users
	SET name = :name,
			comment = :comment,
			github = :github,
			website = :website,
			languages = :languages
	WHERE id = :id;`

	tx := db.MustBegin()
	_, err := tx.NamedExec(query, user)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}