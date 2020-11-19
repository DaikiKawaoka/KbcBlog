package repository

import (
"time"
"app/model"
)

// Article Like

func GetArticleLike(userId int, articleId int) (int, error) {
	query := `SELECT COUNT(*) FROM article_likes WHERE userid = ? AND articleid = ?;`
	var count int
	if err := db.Get(&count, query, userId, articleId); err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticleLikeCount(articleId int) (int,error) {
	query := `SELECT COUNT(*) FROM article_likes WHERE articleid = ?;`
	var count int
	if err := db.Get(&count, query, articleId); err != nil {
		return 0, err
	}
	return count, nil
}

func GetMyLikePost(articleId int) (int,error) {
	query := `SELECT COUNT(*) FROM article_likes WHERE articleid = ?;`
	var count int
	if err := db.Get(&count, query, articleId); err != nil {
		return 0, err
	}
	return count, nil
}

func CreateArticleLike(like *model.ArticleLike) error {

	now := time.Now().In(jst)
  like.Created = now.Format("2006/01/02 15:04:05")

  query := `INSERT INTO article_likes (userid, articleid, created)
	VALUES (:userid, :articleid, :created);`

  tx := db.MustBegin()
  _, err := tx.NamedExec(query, like)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return nil
}

func DeleteArticleLike(userid int, articleid int) error {
	query := "DELETE FROM article_likes WHERE userid = ? AND articleid = ?"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, userid, articleid); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}



// Question Like

func GetQuestionLike(userId int, questionId int) (int, error) {
	query := `SELECT COUNT(*) FROM question_likes WHERE userid = ? AND questionid = ?;`
	var count int
	if err := db.Get(&count, query, userId, questionId); err != nil {
		return 0, err
	}
	return count, nil
}

func GetQuestionLikeCount(questionId int) (int,error) {
	query := `SELECT COUNT(*) FROM question_likes WHERE questionid = ?;`
	var count int
	if err := db.Get(&count, query, questionId); err != nil {
		return 0, err
	}
	return count, nil
}

func CreateQuestionLike(like *model.QuestionLike) error {

	now := time.Now().In(jst)
  like.Created = now.Format("2006/01/02 15:04:05")

  query := `INSERT INTO question_likes (userid, questionid, created)
	VALUES (:userid, :questionid, :created);`

  tx := db.MustBegin()
  _, err := tx.NamedExec(query, like)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return nil
}

func DeleteQuestionLike(userid int, questionid int) error {
	query := "DELETE FROM question_likes WHERE userid = ? AND questionid = ?"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, userid, questionid); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}