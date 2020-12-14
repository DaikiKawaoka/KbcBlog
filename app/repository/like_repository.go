package repository

import (
"time"
"app/model"
)



// GetArticleLike ...
func GetArticleLike(userID int, articleID int) (int, error) {
	query := `SELECT COUNT(*) FROM article_likes WHERE userid = ? AND articleid = ?;`
	var count int
	if err := db.Get(&count, query, userID, articleID); err != nil {
		return 0, err
	}
	return count, nil
}

// GetArticleLikeCount ...
func GetArticleLikeCount(articleID int) (int,error) {
	query := `SELECT COUNT(*) FROM article_likes WHERE articleid = ?;`
	var count int
	if err := db.Get(&count, query, articleID); err != nil {
		return 0, err
	}
	return count, nil
}

// GetMyLikePost ...
func GetMyLikePost(articleID int) (int,error) {
	query := `SELECT COUNT(*) FROM article_likes WHERE articleid = ?;`
	var count int
	if err := db.Get(&count, query, articleID); err != nil {
		return 0, err
	}
	return count, nil
}

// CreateArticleLike ...
func CreateArticleLike(like *model.ArticleLike) error {

	now := time.Now()
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

// DeleteArticleLike ...
func DeleteArticleLike(userid int, articleid int) error {
	query := "DELETE FROM article_likes WHERE userid = ? AND articleid = ?"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, userid, articleid); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}





// GetQuestionLike ...
func GetQuestionLike(userID int, questionID int) (int, error) {
	query := `SELECT COUNT(*) FROM question_likes WHERE userid = ? AND questionid = ?;`
	var count int
	if err := db.Get(&count, query, userID, questionID); err != nil {
		return 0, err
	}
	return count, nil
}

// GetQuestionLikeCount ...
func GetQuestionLikeCount(questionID int) (int,error) {
	query := `SELECT COUNT(*) FROM question_likes WHERE questionid = ?;`
	var count int
	if err := db.Get(&count, query, questionID); err != nil {
		return 0, err
	}
	return count, nil
}

// CreateQuestionLike ...
func CreateQuestionLike(like *model.QuestionLike) error {
	now := time.Now()
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

// DeleteQuestionLike ...
func DeleteQuestionLike(userID int, questionID int) error {
	query := "DELETE FROM question_likes WHERE userid = ? AND questionid = ?"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, userID, questionID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}




// GetArticleCommentLike ...
func GetArticleCommentLike(userID int, articleCommentID int) (int, error) {
	query := `SELECT COUNT(*) FROM article_comment_likes WHERE userid = ? AND article_commentid = ?;`
	var count int
	if err := db.Get(&count, query, userID, articleCommentID); err != nil {
		return 0, err
	}
	return count, nil
}

// GetArticleCommentLikeCount ...
func GetArticleCommentLikeCount(articleCommentID int) (int,error) {
	query := `SELECT COUNT(*) FROM article_comment_likes WHERE article_commentid = ?;`
	var count int
	if err := db.Get(&count, query, articleCommentID); err != nil {
		return 0, err
	}
	return count, nil
}

// CreateArticleCommentLike ...
func CreateArticleCommentLike(like *model.ArticleCommentLike) error {

	now := time.Now()
  like.Created = now.Format("2006/01/02 15:04:05")

  query := `INSERT INTO article_comment_likes (userid, article_commentid, created)
	VALUES (:userid, :article_commentid, :created);`

  tx := db.MustBegin()
  _, err := tx.NamedExec(query, like)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return nil
}

// DeleteArticleCommentLike ...
func DeleteArticleCommentLike(userID int, articleCommentID int) error {
	query := "DELETE FROM article_comment_likes WHERE userid = ? AND article_commentid = ?"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, userID, articleCommentID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}




// GetQuestionCommentLike ...
func GetQuestionCommentLike(userID int, questionCommentID int) (int, error) {
	query := `SELECT COUNT(*) FROM question_comment_likes WHERE userid = ? AND question_commentid = ?;`
	var count int
	if err := db.Get(&count, query, userID, questionCommentID); err != nil {
		return 0, err
	}
	return count, nil
}

// GetQuestionCommentLikeCount ...
func GetQuestionCommentLikeCount(questionCommentID int) (int,error) {
	query := `SELECT COUNT(*) FROM question_comment_likes WHERE question_commentid = ?;`
	var count int
	if err := db.Get(&count, query, questionCommentID); err != nil {
		return 0, err
	}
	return count, nil
}

// CreateQuestionCommentLike ...
func CreateQuestionCommentLike(like *model.QuestionCommentLike) error {
	now := time.Now()
  like.Created = now.Format("2006/01/02 15:04:05")

  query := `INSERT INTO question_comment_likes (userid, question_commentid, created)
	VALUES (:userid, :question_commentid, :created);`

  tx := db.MustBegin()
  _, err := tx.NamedExec(query, like)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return nil
}

// DeleteQuestionCommentLike ...
func DeleteQuestionCommentLike(userid int, questionCommentID int) error {
	query := "DELETE FROM question_comment_likes WHERE userid = ? AND question_commentid = ?"
	tx := db.MustBegin()

	if _, err := tx.Exec(query, userid, questionCommentID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}