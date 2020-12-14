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