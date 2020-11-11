package repository

import (
	"database/sql"
	// "math"
	"app/model"
)

func GetMyUser(id int) (*model.User , error) {
  // クエリ文字列を生成します。
	query := `SELECT id,mail,name,comment
	FROM users
	WHERE id = ?;`

	// クエリ結果を格納する変数を宣言します。
	// 複数件取得の場合はスライスでしたが、一件取得の場合は構造体になります。
	var user model.User

	// 結果を格納する構造体、クエリ文字列、パラメータを指定して SQL を実行します。
	// 一件取得の場合は db.Get()
	if err := db.Get(&user, query, id); err != nil {
		// エラーが発生した場合はエラーを返却します。
		return nil, err
	}

	// エラーがない場合は記事データを返却します。
	return &user, nil
}

// ArticleCreate ...
func UserCreate(user *model.CreateUser) (sql.Result, error) {

	//パスワードハッシュ化
	_, err := user.PasswordHash()
	if err != nil{
		// エラー内容を返却します。
    return nil, err
	}

  // クエリ文字列を生成します。
  query := `INSERT INTO users (mail,passhash,name)
	VALUES (:mail, :passhash, :name);`

  // トランザクションを開始します。
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行します。
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
  // 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
  res, err := tx.NamedExec(query, user)
  if err != nil {
    // エラーが発生した場合はロールバックします。
    tx.Rollback()

    // エラー内容を返却します。
    return nil, err
  }

  // SQL の実行に成功した場合はコミットします。
  tx.Commit()

  // SQL の実行結果を返却します。
  return res, nil
}

func UserGetByID(id int) (*model.User, error) {
	// クエリ文字列を生成します。
	query := `SELECT id,mail,name,comment
	FROM users
	WHERE id = ?`

	var user model.User

	// 結果を格納する構造体、クエリ文字列、パラメータを指定して SQL を実行します。
	// 一件取得の場合は db.Get()
	if err := db.Get(&user, query, id); err != nil {
		return nil, err
	}

	// エラーがない場合は記事データを返却します。
	return &user, nil
}

func UserUpdate(user *model.User) error {

	// クエリ文字列を生成します。
	query := `UPDATE users
	SET name = :name,
			comment = :comment
	WHERE id = :id;`

	// トランザクションを開始します。
	tx := db.MustBegin()

	_, err := tx.NamedExec(query, user)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return err
	}

	// エラーがない場合はコミットします。
	tx.Commit()

	// SQL の実行結果を返却します。
	return nil
}