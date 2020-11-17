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

// 回答した質問を取得
func GetAnswerQuestionList(cursor int) ([]*model.Question, error){
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

func QuestionCreate(question *model.Question) (sql.Result, error) {
	now := time.Now().In(jst)
  question.Created = now.Format("2006/01/02 15:04:05")
  question.Updated = now.Format("2006/01/02 15:04:05")
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
	now := time.Now().In(jst)

	// 構造体に現在日時を設定します。
	question.Updated = now.Format("2006/01/02 15:04:05")

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

func GetUserQuestion(cursor int,userid int) ([]*model.Question, error) {
	// 引数で渡されたカーソルの値が 0 以下の場合は、代わりに int 型の最大値で置き換えます。
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	// ID の降順に記事データを 10 件取得するクエリ文字列を生成します。
	query := `SELECT q.id id,q.userid userid,u.name name,q.title title,q.body body,q.created created,q.updated updated
	FROM questions q,users u
	WHERE q.id < ? and q.userid = u.id and q.userid = ?
	ORDER BY q.id desc
	LIMIT 10`

	// クエリ結果を格納するスライスを初期化します。
	// 10 件取得すると決まっているため、サイズとキャパシティを指定しています。
	questions := make([]*model.Question, 0, 10)

	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
	if err := db.Select(&questions, query, cursor,userid); err != nil {
		return nil, err
	}

	return questions, nil
}