package repository

import (
"database/sql"
"time"
"app/model"
)

// ArticleCommentListByCursor ...
func ArticleCommentListByCursor(articleid int) ([]*model.ArticleComment, error) {

	query := `SELECT c.id id,c.userid userid,c.articleid articleid,u.name name, u.mail mail, u.imgpath imgpath, u.sex sex,c.text text,c.created created,c.updated updated
	FROM article_comments c,users u
	WHERE c.userid = u.id and c.articleid = ?
	ORDER BY c.id desc`

	var comments []*model.ArticleComment

	if err := db.Select(&comments, query, articleid); err != nil {
		return nil, err
	}

	return comments, nil
}

// ArticleCommentCreate ...
func ArticleCommentCreate(comment *model.ArticleComment) (sql.Result, error) {
	now := time.Now()
  comment.Created = now.Format("2006/01/02 15:04:05")
  comment.Updated = now.Format("2006/01/02 15:04:05")
  query := `INSERT INTO article_comments (userid, articleid, text, created, updated)
  VALUES (:userid, :articleid, :text, :created, :updated);`
	tx := db.MustBegin()

  res, err := tx.NamedExec(query, comment)
  if err != nil {
    if re := tx.Rollback(); re != nil {
			return nil,re
		}
    return nil, err
	}
	if re := tx.Commit(); re != nil {
		return nil,re
	}
  return res, nil
}

// ArticleCommentUpdate ...
func ArticleCommentUpdate(comment *model.ArticleComment) (sql.Result, error) {
	now := time.Now()
	comment.Updated = now.Format("2006/01/02 15:04:05")
	query := `UPDATE article_comments
	SET text = :text,
			updated = :updated
	WHERE id = :id;`
	tx := db.MustBegin()

	res, err := tx.NamedExec(query, comment)

	if err != nil {
		if re := tx.Rollback(); re != nil {
			return nil,re
		}
		return nil, err
	}
	if re := tx.Commit(); re != nil {
		return nil,re
	}

	return res, nil
}

// ArticleCommentDelete ...
func ArticleCommentDelete(id int) error {
	query := "DELETE FROM article_comments WHERE id = ?"
	tx := db.MustBegin()
	if _, err := tx.Exec(query, id); err != nil {
		if re := tx.Rollback(); re != nil {
			return re
		}
		return err
	}
	if re := tx.Commit(); re != nil {
		return re
	}
	return nil
}



//QuestionCommentListByCursor ...
func QuestionCommentListByCursor(questionID int) ([]*model.QuestionComment, error) {

	query := `SELECT c.id id,c.userid userid, u.mail mail,c.questionid questionid,u.name name, u.imgpath imgpath, u.sex sex,c.text text,c.created created,c.updated updated
	FROM question_comments c,users u
	WHERE c.userid = u.id and c.questionid = ?
	ORDER BY c.id desc`

	var comments []*model.QuestionComment

	if err := db.Select(&comments, query, questionID); err != nil {
		return nil, err
	}

	return comments, nil
}

// QuestionCommentCreate ...
func QuestionCommentCreate(comment *model.QuestionComment) (sql.Result, error) {
	now := time.Now()
  comment.Created = now.Format("2006/01/02 15:04:05")
  comment.Updated = now.Format("2006/01/02 15:04:05")
  query := `INSERT INTO question_comments (userid, questionid, text, created, updated)
  VALUES (:userid, :questionid, :text, :created, :updated);`

  tx := db.MustBegin()

  res, err := tx.NamedExec(query, comment)
  if err != nil {
    if re := tx.Rollback(); re != nil {
			return nil,re
		}
    return nil, err
	}
	if re := tx.Commit(); re != nil {
		return nil, re
	}

  return res, nil
}

// QuestionCommentUpdate ...
func QuestionCommentUpdate(comment *model.QuestionComment) (sql.Result, error) {
	now := time.Now()
	comment.Updated = now.Format("2006/01/02 15:04:05")
	query := `UPDATE question_comments
	SET text = :text,
			updated = :updated
	WHERE id = :id;`
	tx := db.MustBegin()

	res, err := tx.NamedExec(query, comment)

	if err != nil {
		if re := tx.Rollback(); re != nil {
			return nil,re
		}
		return nil, err
	}
	if re := tx.Commit(); re != nil {
		return nil,re
	}
	return res, nil
}

// QuestionCommentDelete ...
func QuestionCommentDelete(id int) error {
	query := "DELETE FROM question_comments WHERE id = ?;"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, id); err != nil {
		if re := tx.Rollback(); re != nil {
			return re
		}
		return err
	}
	if re := tx.Commit(); re != nil {
		return re
	}
	return nil
}