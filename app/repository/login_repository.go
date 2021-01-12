package repository

import (
	// "database/sql"
	// "math"
	"app/model"
)

// GetLoginPassHash ...
func GetLoginPassHash(mail string) (*model.LoginUser , error) {
	var loginUser model.LoginUser

	// クエリ文字列を生成します。
	query := `SELECT id,mail,passhash
	FROM users
	WHERE mail = ?;`

	if err := db.Get(&loginUser, query, mail); err != nil {
		// エラーが発生した場合はエラーを返却します。
		return nil ,err
	}

	// エラーがない場合は記事データを返却します。
	return &loginUser, nil
}

// GetGuestUser ...
func GetGuestUser() (*model.LoginUser , error) {
	var loginUser model.LoginUser

	// クエリ文字列を生成します。
	query := `SELECT id
	FROM users
	WHERE mail = "kbc920437694@stu.kawahara.ac.jp";`

	if err := db.Get(&loginUser, query); err != nil {
		// エラーが発生した場合はエラーを返却します。
		return nil ,err
	}

	// エラーがない場合は記事データを返却します。
	return &loginUser, nil
}