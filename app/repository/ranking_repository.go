package repository

import (
// "database/sql"
// "time"
// "math"
"app/model"
)

// KBCRankingTop10 ...
func KBCRankingTop10(c string) ([]*model.RankingUser, error) {

	var query string
	if c == "like" {
		query = `SELECT u.id id, u.name name, COUNT(*) count
		FROM articles a inner join users u on a.userid = u.id
		right join article_likes al on al.articleid = a.id
		GROUP BY u.id,u.name
		ORDER BY count desc
		LIMIT 10`
	}else{
		query = `SELECT u.id id, u.name name, COUNT(*) count
		FROM users u right join articles a on a.userid = u.id
		GROUP BY u.id,u.name
		ORDER BY count desc
		LIMIT 10`
	}

	rankingUsers := make([]*model.RankingUser, 0, 10)

	if err := db.Select(&rankingUsers, query); err != nil {
		return nil, err
	}

	return rankingUsers, nil
}