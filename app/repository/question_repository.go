package repository

import (
"database/sql"
"time"
"math"
"app/model"
)

// QuestionListByCursor ...
func QuestionListByCursor(cursor int, scope *model.Scope, userid int) ([]*model.Question, error) {

	questions := make([]*model.Question, 0, 10)
	var query string
	var query1 string
	var query2 string
	var query3 string

	if scope.FriendsOnly {
		// フォローしているユーザの投稿のみ
		query1 = "SELECT q.id id,q.userid userid,u.name name, u.imgpath imgpath, u.sex sex,q.title title,q.tag tag,q.created created,q.updated updated ,COUNT(ql.id) likecount FROM questions q inner join users u on q.userid = u.id inner join follows f on q.userid = f.followedid left join question_likes ql on q.id = ql.questionid "
	}else{
		// 全てのユーザの投稿
		query1 = "SELECT q.id id,q.userid userid,u.name name, u.imgpath imgpath, u.sex sex,q.title title,q.tag tag,q.category category,q.created created,q.updated updated ,COUNT(ql.id) likecount FROM questions q inner join users u on q.userid = u.id left join question_likes ql on q.id = ql.questionid "
	}

	if scope.Order == "new"{
		// IDの降順に記事データを 10 件取得
		query3 = " GROUP BY q.id,q.userid,u.name,q.title,q.tag,q.category,q.created,q.updated ORDER BY q.id desc LIMIT 10;"
		if cursor <= 0 {
			cursor = math.MaxInt32
		}
	}else{
		// いいね数の降順に記事データを 10 件取得
		query3 = " GROUP BY q.id,q.userid,u.name,q.title,q.tag,q.category,q.created,q.updated ORDER BY likecount desc , q.id desc LIMIT ?,10;"
	}

	if scope.Tag == "全て"{
		if scope.Text == ""{
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = "WHERE q.id < ? AND f.followerid = ? "
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor,userid); err != nil {
						return nil, err
					}
				}else{
					query2 = "WHERE f.followerid = ? "
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = "WHERE q.id < ? "
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor); err != nil {
						return nil, err
					}
				}else{
					query = query1 + query3
					if err := db.Select(&questions, query, cursor); err != nil {
						return nil, err
					}
				}
			}
		}else{
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = `WHERE q.id < ? AND q.title LIKE CONCAT("%",?,"%") AND f.followerid = ?  `
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor, scope.Text,userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE q.title LIKE CONCAT("%",?,"%") AND f.followerid = ?  `
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, scope.Text, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE q.id < ? AND q.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor, scope.Text); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE q.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, scope.Text, cursor); err != nil {
						return nil, err
					}
				}
			}
		}

	}else{
		if scope.Text == "" {
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = `WHERE q.id < ? AND q.tag LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor, scope.Tag, userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE q.tag LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, scope.Tag, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE q.id < ? AND q.tag LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor, scope.Tag); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE q.tag LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, scope.Tag, cursor); err != nil {
						return nil, err
					}
				}
			}
		}else{
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = `WHERE q.id < ? AND q.tag LIKE CONCAT("%",?,"%") AND q.title LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor, scope.Tag, scope.Text, userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE q.tag LIKE CONCAT("%",?,"%") AND q.title LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, scope.Tag, scope.Text, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE q.id < ? AND q.tag LIKE CONCAT("%",?,"%") AND q.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, cursor, scope.Tag, scope.Text); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE q.tag LIKE CONCAT("%",?,"%") AND q.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&questions, query, scope.Tag, scope.Text, cursor); err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return questions, nil
}

// QuestionGetByID ...
func QuestionGetByID(id int) (*model.Question, error) {
	query := `SELECT q.id id,q.userid userid,u.name name,u.imgpath imgpath, u.sex sex,q.title title,q.body body,q.tag tag,q.category category,q.created created,q.updated updated
	FROM questions q,users u
	WHERE q.id = ? and q.userid = u.id;`

	var question model.Question
	if err := db.Get(&question, query, id); err != nil {
		return nil, err
	}
	return &question, nil
}

// QuestionGetByID2 記事IDでユーザーIDだけ返す
func QuestionGetByID2(id int) (int, error) {
	query := `select userid from questions where id = ?;`
	var userid int
	if err := db.Get(&userid, query, id); err != nil {
		return 0, err
	}
	return userid, nil
}

// QuestionCreate ...
func QuestionCreate(question *model.Question) (sql.Result, error) {
	now := time.Now()
  question.Created = now.Format("2006/01/02 15:04:05")
  question.Updated = now.Format("2006/01/02 15:04:05")
  query := `INSERT INTO questions (userid, title, body, tag, category, created, updated)
  VALUES (:userid, :title, :body, :tag, :category, :created, :updated);`

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
			tag = :tag,
			category = :category,
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
	if cursor <= 0 {
		cursor = math.MaxInt32
	}
	// ID の降順に記事データを 10 件取得
	query := `SELECT q.id id,q.userid userid,u.name name,q.title title,q.body body,q.tag tag, q.category category ,q.created created,q.updated updated, COUNT(l.id) likecount
	FROM questions q inner join users u on q.userid = u.id
	left join question_likes l on q.id = l.questionid
	WHERE q.id < ? and q.userid = u.id and q.userid = ?
	GROUP BY q.id,q.userid,u.name,q.title,q.body,q.tag,q.category,q.created,q.updated
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