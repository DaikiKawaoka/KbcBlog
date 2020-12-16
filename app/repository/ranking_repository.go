package repository

import (
// "database/sql"
// "time"
// "math"
"app/model"
)

// KBCArticleRankingTop10 ...
func KBCArticleRankingTop10(c string) ([]*model.RankingUser, error) {

	var query string
	if c == "like" {
		query = `SELECT u.id id, u.name name , u.comment comment, COUNT(al.id) count
		FROM users u LEFT JOIN articles a on a.userid = u.id
		LEFT join article_likes al on al.articleid = a.id
		GROUP BY u.id,u.name,u.comment ORDER BY count desc LIMIT 10;`
	}else{
		query = `SELECT u.id id, u.name name, u.comment comment, COUNT(a.id) count
		FROM users u LEFT JOIN articles a on a.userid = u.id
		GROUP BY u.id,u.name,u.comment ORDER BY count desc LIMIT 10;`
	}

	rankingUsers := make([]*model.RankingUser, 0, 10)

	if err := db.Select(&rankingUsers, query); err != nil {
		return nil, err
	}

	return rankingUsers, nil
}

// KBCQuestionRankingTop10 ...
func KBCQuestionRankingTop10(c string) ([]*model.RankingUser, error) {

	var query string
	if c == "post" {
		// 質問数
		query = `SELECT u.id id, u.name name,u.comment comment, COUNT(q.id) count
		FROM users u LEFT JOIN questions q on q.userid = u.id
		GROUP BY u.id,u.name,u.comment
		ORDER BY count desc
		LIMIT 10;`
	}else{
		// 回答数
		query = `SELECT u.id id, u.name name, u.comment comment, COUNT(qc.id) count
		FROM users u LEFT join question_comments qc on u.id = qc.userid
		GROUP BY u.id,u.name,u.comment ORDER BY count desc LIMIT 10;`
	}

	rankingUsers := make([]*model.RankingUser, 0, 10)

	if err := db.Select(&rankingUsers, query); err != nil {
		return nil, err
	}

	return rankingUsers, nil
}