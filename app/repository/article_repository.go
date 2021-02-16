package repository

import (
"database/sql"
"time"
"math"
"app/model"
)

// ArticleListByCursor ...
func ArticleListByCursor(cursor int, scope *model.Scope, userid int) ([]*model.Article, error) {

	articles := make([]*model.Article, 0, 10)
	var query string
	var query1 string
	var query2 string
	var query3 string

	if scope.FriendsOnly {
		// フォローしているユーザの投稿のみ
		query1 = "SELECT a.id id,a.userid userid,u.name name, u.imgpath imgpath, a.title title,a.tag tag,a.created created,a.updated updated ,COUNT(al.id) likecount FROM articles a inner join users u on a.userid = u.id inner join follows f on a.userid = f.followedid left join article_likes al on a.id = al.articleid  "
	}else{
		// 全てのユーザの投稿
		query1 = "SELECT a.id id,a.userid userid,u.name name, u.imgpath imgpath,a.title title,a.tag tag,a.created created,a.updated updated ,COUNT(al.id) likecount FROM articles a inner join users u on a.userid = u.id left join article_likes al on a.id = al.articleid "
	}

	if scope.Order == "new"{
		// IDの降順に記事データを 10 件取得
		query3 = " GROUP BY a.id,a.userid,u.name,a.title,a.created,a.updated ORDER BY a.id desc LIMIT 10;"
		if cursor <= 0 {
			cursor = math.MaxInt32
		}
	}else{
		// いいね数の降順に記事データを 10 件取得
		query3 = " GROUP BY a.id,a.userid,u.name,a.title,a.created,a.updated ORDER BY likecount desc , a.id desc LIMIT ?,10;"
	}

	if scope.Tag == "全て"{
		if scope.Text == ""{
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = "WHERE a.id < ? AND f.followerid = ? "
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor,userid); err != nil {
						return nil, err
					}
				}else{
					query2 = "WHERE f.followerid = ? "
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = "WHERE a.id < ? "
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor); err != nil {
						return nil, err
					}
				}else{
					query = query1 + query3
					if err := db.Select(&articles, query, cursor); err != nil {
						return nil, err
					}
				}
			}
		}else{
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.title LIKE CONCAT("%",?,"%") AND f.followerid = ?  `
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor, scope.Text,userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.title LIKE CONCAT("%",?,"%") AND f.followerid = ?  `
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, scope.Text, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor, scope.Text); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, scope.Text, cursor); err != nil {
						return nil, err
					}
				}
			}
		}

	}else{
		if scope.Text == "" {
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.tag LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor, scope.Tag, userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.tag LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, scope.Tag, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.tag LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor, scope.Tag); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.tag LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, scope.Tag, cursor); err != nil {
						return nil, err
					}
				}
			}
		}else{
			if scope.FriendsOnly{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.tag LIKE CONCAT("%",?,"%") AND a.title LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor, scope.Tag, scope.Text, userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.tag LIKE CONCAT("%",?,"%") AND a.title LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, scope.Tag, scope.Text, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.tag LIKE CONCAT("%",?,"%") AND a.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, cursor, scope.Tag, scope.Text); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.tag LIKE CONCAT("%",?,"%") AND a.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					if err := db.Select(&articles, query, scope.Tag, scope.Text, cursor); err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return articles, nil
}

// ArticleGetByID ...
func ArticleGetByID(id int) (*model.Article, error) {
	query := `SELECT a.id id,a.userid userid,u.name name, u.mail mail, u.imgpath imgpath, u.sex sex, a.title title,a.body body,a.tag tag,a.created created,a.updated updated
	FROM articles a,users u
	WHERE a.id = ? and a.userid = u.id;`

	var article model.Article
	if err := db.Get(&article, query, id); err != nil {
		return nil, err
	}

	return &article, nil
}

// ArticleGetByID2 記事IDでユーザーIDだけ返す
func ArticleGetByID2(id int) (int, error) {
	query := `select userid from articles where id = ?;`
	var userid int
	if err := db.Get(&userid, query, id); err != nil {
		return 0, err
	}
	return userid, nil
}

// ArticleCreate ...
func ArticleCreate(article *model.Article) (sql.Result, error) {
	now := time.Now()
  article.Created = now.Format("2006/01/02 15:04:05")
  article.Updated = now.Format("2006/01/02 15:04:05")

  query := `INSERT INTO articles (userid,title, body, tag, created, updated)
  VALUES (:userid, :title, :body, :tag, :created, :updated);`

  tx := db.MustBegin()

  res, err := tx.NamedExec(query, article)
  if err != nil {
    if re := tx.Rollback(); re != nil {
			return nil , re
		}
    return nil, err
	}
	if re := tx.Commit(); re != nil {
		return nil , re
	}
  return res, nil
}

// ArticleUpdate ...
func ArticleUpdate(article *model.Article) (sql.Result, error) {
	now := time.Now()
	article.Updated = now.Format("2006/01/02 15:04:05")

	query := `UPDATE articles
	SET title = :title,
			body = :body,
			updated = :updated,
			tag  = :tag
	WHERE id = :id;`
	tx := db.MustBegin()
	res, err := tx.NamedExec(query, article)
	if err != nil {
		if re := tx.Rollback(); re != nil {
			return nil , re
		}
		return nil, err
	}
	if re := tx.Commit(); re != nil {
		return nil , re
	}
	return res, nil
}

// ArticleDelete ...
func ArticleDelete(id int) error {
	query := "DELETE FROM articles WHERE id = ?"
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

// GetUserArticle ...
func GetUserArticle(cursor int,userid int) ([]*model.Article, error) {
	if cursor <= 0 {
		cursor = math.MaxInt32
	}
	query := `SELECT a.id id,a.userid userid,u.name name,a.title title,a.tag tag,a.created created,a.updated updated, COUNT(l.id) likecount
	FROM articles a inner join users u on a.userid = u.id
	left join article_likes l on a.id = l.articleid
	WHERE a.id < ? and a.userid = u.id and a.userid = ?
	GROUP BY a.id,a.userid,u.name,a.title,a.created,a.updated
	ORDER BY a.id desc
	LIMIT 30`
	articles := make([]*model.Article, 0, 30)
	if err := db.Select(&articles, query, cursor,userid); err != nil {
		return nil, err
	}

	return articles, nil
}

// GetFavoriteArticle ...
func GetFavoriteArticle(cursor int,userid int) ([]*model.FavoriteArticle, error) {
	if cursor <= 0 {
		cursor = math.MaxInt32
	}
	query := `SELECT a.id t_id,a.userid userid,u.name name, u.imgpath imgpath, a.title title,a.tag tag,a.updated updated,
	(select COUNT(l.id) from articles a inner join users u on a.userid = u.id right join article_likes l on a.id = l.articleid where a.id = t_id Group by a.id) as likecount,l.id likeid
	FROM articles a inner join users u on a.userid = u.id
	right join article_likes l on a.id = l.articleid
	WHERE l.id < ? and l.userid = ?
	GROUP BY t_id,userid,name,imgpath,title,updated,likeid
	ORDER BY likeid desc
	LIMIT 10`
	articles := make([]*model.FavoriteArticle, 0, 10)
	if err := db.Select(&articles, query, cursor,userid); err != nil {
		return nil, err
	}

	return articles, nil
}