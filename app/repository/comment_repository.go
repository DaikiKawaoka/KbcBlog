package repository

import (
"database/sql"
"time"
"app/model"
)

// ArticleComment

func ArticleCommentListByCursor(articleid int) ([]*model.ArticleComment, error) {

	query := `SELECT c.id id,c.userid userid,c.articleid articleid,u.name name,c.text text,c.created created,c.updated updated
	FROM article_comments c,users u
	WHERE c.userid = u.id and c.articleid = ?
	ORDER BY c.id desc`

	var comments []*model.ArticleComment

	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
	if err := db.Select(&comments, query, articleid); err != nil {
		return nil, err
	}

	return comments, nil
}

func ArticleCommentCreate(comment *model.ArticleComment) (sql.Result, error) {
	// 現在日時を取得します

	now := time.Now().In(jst)

  // 構造体に現在日時を設定します。
  comment.Created = now.Format("2006/01/02 15:04:05")
  comment.Updated = now.Format("2006/01/02 15:04:05")

  // クエリ文字列を生成します。
  query := `INSERT INTO article_comments (userid, articleid, text, created, updated)
  VALUES (:userid, :articleid, :text, :created, :updated);`

  // トランザクションを開始します。
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行します。
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
  // 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
  res, err := tx.NamedExec(query, comment)
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

func ArticleCommentUpdate(comment *model.ArticleComment) (sql.Result, error) {
	// 現在日時を取得します
	now := time.Now().In(jst)

	// 構造体に現在日時を設定します。
	comment.Updated = now.Format("2006/01/02 15:04:05")

	// クエリ文字列を生成します。
	query := `UPDATE article_comments
	SET text = :text,
			updated = :updated
	WHERE id = :id;`

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と引数で渡ってきた構造体を指定して、SQL を実行します。
	// クエリ文字列内の :title, :body, :id には、
	// 第 2 引数の Article 構造体の Title, Body, ID が bind されます。
	// 構造体に db タグで指定した値が紐付けされます。
	res, err := tx.NamedExec(query, comment)

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

func ArticleCommentDelete(id int) error {
	// 記事データを削除するクエリ文字列を生成します。
	query := "DELETE FROM article_comments WHERE id = ?"

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



//QuestionComment

func QuestionCommentListByCursor(questionId int) ([]*model.QuestionComment, error) {

	query := `SELECT c.id id,c.userid userid,c.questionid questionid,u.name name,c.text text,c.created created,c.updated updated
	FROM question_comments c,users u
	WHERE c.userid = u.id and c.questionid = ?
	ORDER BY c.id desc`

	var comments []*model.QuestionComment

	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
	if err := db.Select(&comments, query, questionId); err != nil {
		return nil, err
	}

	return comments, nil
}

func QuestionCommentCreate(comment *model.QuestionComment) (sql.Result, error) {
	// 現在日時を取得します

	now := time.Now().In(jst)

  // 構造体に現在日時を設定します。
  comment.Created = now.Format("2006/01/02 15:04:05")
  comment.Updated = now.Format("2006/01/02 15:04:05")

  // クエリ文字列を生成します。
  query := `INSERT INTO question_comments (userid, questionid, text, created, updated)
  VALUES (:userid, :questionid, :text, :created, :updated);`

  // トランザクションを開始します。
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行します。
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
  // 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
  res, err := tx.NamedExec(query, comment)
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

func QuestionCommentUpdate(comment *model.QuestionComment) (sql.Result, error) {
	// 現在日時を取得します
	now := time.Now().In(jst)

	// 構造体に現在日時を設定します。
	comment.Updated = now.Format("2006/01/02 15:04:05")

	// クエリ文字列を生成します。
	query := `UPDATE question_comments
	SET text = :text,
			updated = :updated
	WHERE id = :id;`

	// トランザクションを開始します。
	tx := db.MustBegin()

	// クエリ文字列と引数で渡ってきた構造体を指定して、SQL を実行します。
	// クエリ文字列内の :title, :body, :id には、
	// 第 2 引数の Article 構造体の Title, Body, ID が bind されます。
	// 構造体に db タグで指定した値が紐付けされます。
	res, err := tx.NamedExec(query, comment)

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

func QuestionCommentDelete(id int) error {
	// 記事データを削除するクエリ文字列を生成します。
	query := "DELETE FROM question_comments WHERE id = ?;"

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