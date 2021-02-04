package repository

import (
	"database/sql"
	"app/model"
)

// GetMyUser ...
func GetMyUser(id int) (*model.User , error) {
	query := `SELECT id,mail,name,comment,sex,imgpath
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

	if user.Sex == 1 {
		user.ImgPath = "https://s3.ap-northeast-1.amazonaws.com/www.kbcblog-s3.com/userIcon/man.png"
	}else{
		user.ImgPath = "https://s3.ap-northeast-1.amazonaws.com/www.kbcblog-s3.com/userIcon/woman.png"
	}

  query := `INSERT INTO users (mail,passhash,name,sex,imgpath)
	VALUES (:mail, :passhash, :name, :sex, :imgpath);`

  tx := db.MustBegin()
  res, err := tx.NamedExec(query, user)
  if err != nil {
    if re := tx.Rollback(); re != nil {
			return nil, re
		}
    return nil, err
  }
  if re := tx.Commit(); re != nil {
		return nil, re
	}
  return res, nil
}

// UserGetByID ...
func UserGetByID(id int) (*model.User, error) {
	query := `SELECT id,mail,name,comment,github,website,languages,sex,imgpath
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
			mail = :mail,
			comment = :comment,
			github = :github,
			website = :website,
			languages = :languages
	WHERE id = :id;`

	tx := db.MustBegin()
	_, err := tx.NamedExec(query, user)
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

// UserImgUpdate ...
func UserImgUpdate(user *model.UserImgUpdateClass) error {

	query := `UPDATE users
	SET imgpath = :imgpath
	WHERE id = :id;`

	tx := db.MustBegin()
	_, err := tx.NamedExec(query, user)
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

// UserImgDelete ...
func UserImgDelete(user *model.User) error {

	var query string
	if user.Sex == 1 {
		user.ImgPath = "https://s3.ap-northeast-1.amazonaws.com/www.kbcblog-s3.com/userIcon/man.png"
	}else{
		user.ImgPath = "https://s3.ap-northeast-1.amazonaws.com/www.kbcblog-s3.com/userIcon/woman.png"
	}
	query = `UPDATE users
	SET imgpath = :imgpath
	WHERE id = :id;`

	tx := db.MustBegin()
	_, err := tx.NamedExec(query, user)
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