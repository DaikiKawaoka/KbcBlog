package repository

import (
	"app/model"
)

func GetPassword(id int) (string,error) {
  // クエリ文字列を生成します。
	query := `SELECT passhash
	FROM users
	WHERE id = ?;`
	var pass string
	if err := db.Get(&pass, query, id); err != nil {
		return "", err
	}

	// エラーがない場合は記事データを返却します。
	return pass, nil
}

func PasswordUpdate(pass *model.Password) error {

	// クエリ文字列を生成します。
	query := `UPDATE users
	SET passhash = :newpassword
	WHERE id = :id;`

	// トランザクションを開始します。
	tx := db.MustBegin()
	_, err := tx.NamedExec(query, pass)
	if err != nil {
		tx.Rollback()
		return err
	}
	// エラーがない場合はコミットします。
	tx.Commit()
	// SQL の実行結果を返却します。
	return nil
}