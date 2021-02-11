package repository

import "app/model"


// KBCArticleRankingTop10 ...
func KBCArticleRankingTop10(rankingType int, period string) ([]*model.RankingUser, error) {

	var query string
	if rankingType == 0 {
		if period == "all"{
			query = `SELECT u.id id, u.name name, u.mail mail , u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(al.id) count
			FROM users u LEFT JOIN articles a on a.userid = u.id
			LEFT join article_likes al on al.articleid = a.id
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name,u.comment ORDER BY count desc LIMIT 10;`
		}else{
			// 1ヶ月以内
			query = `SELECT DISTINCT * FROM((SELECT u.id id, u.name name, u.mail mail , u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(al.id) count
			FROM users u LEFT JOIN articles a on a.userid = u.id
			LEFT join article_likes al on al.articleid = a.id
			WHERE u.id <> 920437694 and al.created >= DATE_ADD(NOW(), interval -1 month)
			GROUP BY u.id,u.name)
			UNION
			(SELECT u.id id, u.name name, u.mail mail ,u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(al.id) count
			FROM users u LEFT JOIN articles a on a.userid = u.id
			LEFT join article_likes al on al.articleid = a.id
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name
			HAVING count = 0)) AS new ORDER BY count desc LIMIT 10;`
		}
	}else{
		if period == "all"{
			query = `SELECT u.id id, u.name name, u.mail mail , u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(a.id) count
			FROM users u LEFT JOIN articles a on a.userid = u.id
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name,u.comment ORDER BY count desc LIMIT 10;`
		}else{
			// 1ヶ月以内
			query = `SELECT DISTINCT * FROM((SELECT u.id id, u.name name, u.mail mail ,u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(a.id) count
			FROM users u LEFT JOIN articles a on a.userid = u.id
			WHERE u.id <> 920437694 and a.created >= DATE_ADD(NOW(), interval -1 month)
			GROUP BY u.id,u.name,u.comment)
			UNION
			(SELECT u.id id, u.name name, u.mail mail ,u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(a.id) count
			FROM users u LEFT JOIN articles a on a.userid = u.id
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name,u.comment
			HAVING count = 0)) AS new ORDER BY count desc LIMIT 10;`
		}
	}
	rankingUsers := make([]*model.RankingUser, 0, 10)
	if err := db.Select(&rankingUsers, query); err != nil {
		return nil, err
	}
	return rankingUsers, nil
}

// KBCQuestionRankingTop10 ...
func KBCQuestionRankingTop10(rankingType int, period string) ([]*model.RankingUser, error) {

	var query string
	if rankingType == 0 {
		// 質問数
		if period == "all"{
			query = `SELECT u.id id, u.name name, u.mail mail ,u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(q.id) count
			FROM users u LEFT JOIN questions q on q.userid = u.id
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name,u.comment
			ORDER BY count desc
			LIMIT 10;`
		}else{
			query = `SELECT DISTINCT * FROM((SELECT u.id id, u.name name, u.mail mail , u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(q.id) count
			FROM users u LEFT JOIN questions q on q.userid = u.id
			WHERE u.id <> 920437694 and q.created >= DATE_ADD(NOW(), interval -1 month)
			GROUP BY u.id,u.name)
			UNION
			(SELECT u.id id, u.name name, u.mail mail ,u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(q.id) count
			FROM users u LEFT JOIN questions q on q.userid = u.id
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name
			HAVING count = 0)) AS new ORDER BY count desc LIMIT 10;`
		}
	}else{
		// 回答数
		if period == "all"{
			query = `SELECT u.id id, u.name name, u.mail mail , u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(qc.id) count
			FROM users u LEFT join question_comments qc on u.id = qc.userid
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name,u.comment ORDER BY count desc LIMIT 10;`
		}else{
			query = `SELECT DISTINCT * FROM((SELECT u.id id, u.name name, u.mail mail , u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(qc.id) count
			FROM users u LEFT join question_comments qc on u.id = qc.userid
			WHERE u.id <> 920437694 and qc.created >= DATE_ADD(NOW(), interval -1 month)
			GROUP BY u.id,u.name)
			UNION
			(SELECT u.id id, u.name name, u.mail mail ,u.comment comment, u.imgpath imgpath, u.sex sex, COUNT(qc.id) count
			FROM users u LEFT join question_comments qc on u.id = qc.userid
			WHERE u.id <> 920437694
			GROUP BY u.id,u.name
			HAVING count = 0)) AS new ORDER BY count desc LIMIT 10;`
		}
	}
	rankingUsers := make([]*model.RankingUser, 0, 10)
	if err := db.Select(&rankingUsers, query); err != nil {
		return nil, err
	}
	return rankingUsers, nil
}