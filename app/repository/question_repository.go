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

	query := `SELECT *
	FROM questions
	WHERE id = ?;`

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