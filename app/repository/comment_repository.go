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

	if err := db.Select(&comments, query, articleid); err != nil {
		return nil, err
	}

	return comments, nil
}

func ArticleCommentCreate(comment *model.ArticleComment) (sql.Result, error) {
	now := time.Now().In(jst)
  comment.Created = now.Format("2006/01/02 15:04:05")
  comment.Updated = now.Format("2006/01/02 15:04:05")
  query := `INSERT INTO article_comments (userid, articleid, text, created, updated)
  VALUES (:userid, :articleid, :text, :created, :updated);`
	tx := db.MustBegin()

  res, err := tx.NamedExec(query, comment)
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  tx.Commit()
  return res, nil
}

func ArticleCommentUpdate(comment *model.ArticleComment) (sql.Result, error) {
	now := time.Now().In(jst)
	comment.Updated = now.Format("2006/01/02 15:04:05")
	query := `UPDATE article_comments
	SET text = :text,
			updated = :updated
	WHERE id = :id;`
	tx := db.MustBegin()

	res, err := tx.NamedExec(query, comment)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, nil
}

func ArticleCommentDelete(id int) error {
	query := "DELETE FROM article_comments WHERE id = ?"
	tx := db.MustBegin()
	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}



//QuestionComment

func QuestionCommentListByCursor(questionId int) ([]*model.QuestionComment, error) {

	query := `SELECT c.id id,c.userid userid,c.questionid questionid,u.name name,c.text text,c.created created,c.updated updated
	FROM question_comments c,users u
	WHERE c.userid = u.id and c.questionid = ?
	ORDER BY c.id desc`

	var comments []*model.QuestionComment

	if err := db.Select(&comments, query, questionId); err != nil {
		return nil, err
	}

	return comments, nil
}

func QuestionCommentCreate(comment *model.QuestionComment) (sql.Result, error) {
	now := time.Now().In(jst)
  comment.Created = now.Format("2006/01/02 15:04:05")
  comment.Updated = now.Format("2006/01/02 15:04:05")
  query := `INSERT INTO question_comments (userid, questionid, text, created, updated)
  VALUES (:userid, :questionid, :text, :created, :updated);`

  tx := db.MustBegin()

  res, err := tx.NamedExec(query, comment)
  if err != nil {
    tx.Rollback()
    return nil, err
  }
  tx.Commit()
  return res, nil
}

func QuestionCommentUpdate(comment *model.QuestionComment) (sql.Result, error) {
	now := time.Now().In(jst)
	comment.Updated = now.Format("2006/01/02 15:04:05")
	query := `UPDATE question_comments
	SET text = :text,
			updated = :updated
	WHERE id = :id;`
	tx := db.MustBegin()

	res, err := tx.NamedExec(query, comment)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, nil
}

func QuestionCommentDelete(id int) error {
	query := "DELETE FROM question_comments WHERE id = ?;"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}