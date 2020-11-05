package repository

import (
"database/sql"
"time"
"math"
"app/model"
)

func QuestionListByCursor(cursor int) ([]*model.Question, error) {

	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	query := `SELECT q.id id,q.userid userid,u.name name,q.title title,q.body body,q.created created,q.updated updated
	FROM questions q,users u
	WHERE q.id < ? and q.userid = u.id
	ORDER BY q.id desc
	LIMIT 10`

	questions := make([]*model.Question, 0, 10)

	if err := db.Select(&questions, query, cursor); err != nil {
		return nil, err
	}

	return questions, nil
}

func QuestionGetByID(id int) (*model.Question, error) {

	query := `SELECT q.id id,q.userid userid,u.name name,q.title title,q.body body,q.created created,q.updated updated
	FROM questions q,users u
	WHERE q.id = ? and q.userid = u.id;`

	var question model.Question

	if err := db.Get(&question, query, id); err != nil {
		return nil, err
	}

	return &question, nil
}

func QuestionCreate(question *model.Question) (sql.Result, error) {

  now := time.Now()

  question.Created = now
  question.Updated = now

  query := `INSERT INTO questions (userid,title, body, created, updated)
  VALUES (:userid, :title, :body, :created, :updated);`

  tx := db.MustBegin()

  res, err := tx.NamedExec(query, question)
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  tx.Commit()
  return res, nil
}

func QuestionUpdate(question *model.Question) (sql.Result, error) {
	// 現在日時を取得します
	now := time.Now()

	// 構造体に現在日時を設定します。
	question.Updated = now

	// クエリ文字列を生成します。
	query := `UPDATE questions
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
	res, err := tx.NamedExec(query, question)

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

func QuestionDelete(id int) error {
	// 記事データを削除するクエリ文字列を生成します。
	query := "DELETE FROM questions WHERE id = ?"

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