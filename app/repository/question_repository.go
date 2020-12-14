package repository

import (
"database/sql"
"time"
"math"
"app/model"
)

//QuestionListByCursor ...
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

// QuestionGetByID ...
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

// QuestionCreate ...
func QuestionCreate(question *model.Question) (sql.Result, error) {
	now := time.Now()
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

// QuestionUpdate ...
func QuestionUpdate(question *model.Question) (sql.Result, error) {
	now := time.Now()
	question.Updated = now.Format("2006/01/02 15:04:05")

	query := `UPDATE questions
	SET title = :title,
			body = :body,
			updated = :updated
	WHERE id = :id;`

	tx := db.MustBegin()
	res, err := tx.NamedExec(query, question)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, nil
}

// QuestionDelete ...
func QuestionDelete(ID int) error {
	query := "DELETE FROM questions WHERE id = ?"
	tx := db.MustBegin()
	if _, err := tx.Exec(query, ID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// GetUserQuestion ...
func GetUserQuestion(cursor int,userID int) ([]*model.Question, error) {
	// 引数で渡されたカーソルの値が 0 以下の場合は、代わりに int 型の最大値で置き換えます。
	if cursor <= 0 {
		cursor = math.MaxInt32
	}
	// ID の降順に記事データを 10 件取得するクエリ文字列を生成します。
	query := `SELECT q.id id,q.userid userid,u.name name,q.title title,q.body body,q.created created,q.updated updated, COUNT(l.id) likecount
	FROM questions q inner join users u on q.userid = u.id
	left join question_likes l on q.id = l.questionid
	WHERE q.id < ? and q.userid = u.id and q.userid = ?
	GROUP BY q.id,q.userid,u.name,q.title,q.body,q.created,q.updated
	ORDER BY q.id desc
	LIMIT 10`

	questions := make([]*model.Question, 0, 10)
	if err := db.Select(&questions, query, cursor,userID); err != nil {
		return nil, err
	}
	return questions, nil
}

// GetAnswerQuestionList 回答した質問を取得
func GetAnswerQuestionList(cursor int, userID int) ([]*model.AnswerQuestion, error){
	if cursor <= 0 {
		cursor = math.MaxInt32
	}
	query := `SELECT q.id id,q.userid userid,u.name name,q.title title, c.text text,c.created created
	FROM questions q,users u,question_comments c
	WHERE q.id < ? AND q.userid = u.id AND q.id = c.questionid AND c.userid = ?
	ORDER BY q.id desc
	LIMIT 10`

	answerQuestions := make([]*model.AnswerQuestion, 0, 10)
	if err := db.Select(&answerQuestions, query, cursor ,userID); err != nil {
		return nil, err
	}
	return answerQuestions, nil
}