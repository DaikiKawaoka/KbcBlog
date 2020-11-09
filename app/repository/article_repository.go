package repository

import (
"database/sql"
"time"
"math"
"app/model"
)


// func ArticleList() ([]*model.Article, error) {
// 	query := `SELECT id,title,body FROM articles;`

// 	var articles []*model.Article
// 	if err := db.Select(&articles, query); err != nil {
// 		return nil, err
// 	}

// 	return articles, nil
// }

// ArticleListByCursor ...
func ArticleListByCursor(cursor int) ([]*model.Article, error) {
	// 引数で渡されたカーソルの値が 0 以下の場合は、代わりに int 型の最大値で置き換えます。
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	// ID の降順に記事データを 10 件取得するクエリ文字列を生成します。
	query := `SELECT a.id id,a.userid userid,u.name name,a.title title,a.body body,a.created created,a.updated updated
	FROM articles a,users u
	WHERE a.id < ? and a.userid = u.id
	ORDER BY a.id desc
	LIMIT 10`

	// クエリ結果を格納するスライスを初期化します。
	// 10 件取得すると決まっているため、サイズとキャパシティを指定しています。
	articles := make([]*model.Article, 0, 10)

	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
	if err := db.Select(&articles, query, cursor); err != nil {
		return nil, err
	}

	return articles, nil
}

// ArticleGetByID ...
func ArticleGetByID(id int) (*model.Article, error) {
	// クエリ文字列を生成します。
	query := `SELECT a.id id,a.userid userid,u.name name,a.title title,a.body body,a.created created,a.updated updated
	FROM articles a,users u
	WHERE a.id = ? and a.userid = u.id;`

	// クエリ結果を格納する変数を宣言します。
	// 複数件取得の場合はスライスでしたが、一件取得の場合は構造体になります。
	var article model.Article

	// 結果を格納する構造体、クエリ文字列、パラメータを指定して SQL を実行します。
	// 一件取得の場合は db.Get()
	if err := db.Get(&article, query, id); err != nil {
		// エラーが発生した場合はエラーを返却します。
		return nil, err
	}

	// エラーがない場合は記事データを返却します。
	return &article, nil
}

// ArticleCreate ...
func ArticleCreate(article *model.Article) (sql.Result, error) {
  // 現在日時を取得します
	now := time.Now().In(jst)

  // 構造体に現在日時を設定します。
  article.Created = now.Format("2006/01/02 15:04:05")
  article.Updated = now.Format("2006/01/02 15:04:05")

  // クエリ文字列を生成します。
  query := `INSERT INTO articles (userid,title, body, created, updated)
  VALUES (:userid, :title, :body, :created, :updated);`

  // トランザクションを開始します。
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行します。
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
  // 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
  res, err := tx.NamedExec(query, article)
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

// ArticleUpdate ...
func ArticleUpdate(article *model.Article) (sql.Result, error) {
	// 現在日時を取得します
	now := time.Now().In(jst)

	// 構造体に現在日時を設定します。
	article.Updated = now.Format("2006/01/02 15:04:05")

	// クエリ文字列を生成します。
	query := `UPDATE articles
	SET title = :title,
			body = :body,
			updated = :updated
	WHERE id = :id;`

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と引数で渡ってきた構造体を指定して、SQL を実行します。
	// クエリ文字列内の :title, :body, :id には、
	// 第 2 引数の Article 構造体の Title, Body, ID が bind されます。
	// 構造体に db タグで指定した値が紐付けされます。
	res, err := tx.NamedExec(query, article)

	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラーを返却します。
		return nil, err
	}

	// エラーがない場合はコミットします。
	tx.Commit()

	// SQL の実行結果を返却します。
	return res, nil
}

// ArticleDelete ...
func ArticleDelete(id int) error {
	// 記事データを削除するクエリ文字列を生成します。
	query := "DELETE FROM articles WHERE id = ?"

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列とパラメータを指定して SQL を実行します。
	if _, err := tx.Exec(query, id); err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()

		// エラー内容を返却します。
		return err
	}

	// エラーがない場合はコミットします。
	return tx.Commit()
}