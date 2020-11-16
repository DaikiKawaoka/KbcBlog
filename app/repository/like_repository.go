package repository

import (
"time"
"app/model"
)

func GetArticleLike(userId int, articleId int) (*model.ArticleLike, error) {
	// クエリ文字列を生成します。
	query := `SELECT * FROM article_likes WHERE userid = ? AND articleid = ?;`

	var like model.ArticleLike

	if err := db.Get(&like, query, userId, articleId); err != nil {
		return nil, err
	}

	// エラーがない場合は記事データを返却します。
	return &like, nil
}

func CreateArticleLike(like *model.ArticleLike) error {
	// 現在日時を取得します

	now := time.Now().In(jst)

  // 構造体に現在日時を設定します。
  like.Created = now.Format("2006/01/02 15:04:05")

  // クエリ文字列を生成します。
  query := `INSERT INTO article_likes (userid, articleid, created)
  VALUES (:userid, :articleid, :created);`

  // トランザクションを開始します。
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行します。
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
  // 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
  _, err := tx.NamedExec(query, like)
  if err != nil {
    // エラーが発生した場合はロールバックします。
    tx.Rollback()

    // エラー内容を返却します。
    return err
  }

  // SQL の実行に成功した場合はコミットします。
  tx.Commit()

  // SQL の実行結果を返却します。
  return nil
}

// ArticleDelete ...
func DeleteArticleLike(userid int, articleid int) error {
	// 記事データを削除するクエリ文字列を生成します。
	query := "DELETE FROM article_likes WHERE userid = ? AND articleid = ?"

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列とパラメータを指定して SQL を実行します。
	if _, err := tx.Exec(query, userid, articleid); err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()
		// エラー内容を返却します。
		return err
	}
	return tx.Commit()
}