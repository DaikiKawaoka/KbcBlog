package repository

import (
"database/sql"
"time"

"app/model"
)

func ArticleList() ([]*model.Article, error) {
	query := `SELECT id,title,body FROM articles;`

	var articles []*model.Article
	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}

	return articles, nil
}

// ArticleCreate ...
func ArticleCreate(article *model.Article) (sql.Result, error) {
  // 現在日時を取得します
  now := time.Now()

  // 構造体に現在日時を設定します。
  article.Created = now
  article.Updated = now

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