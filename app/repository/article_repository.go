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
	// query4 := "AND u.id = f.followerid AND f.followerid = ? "

	if scope.FriendsOnly {
		// フォローしているユーザの投稿のみ
		query1 = "SELECT a.id id,a.userid userid,u.name name,a.title title,a.tag tag,a.created created,a.updated updated ,COUNT(al.id) likecount FROM articles a inner join users u on a.userid = u.id inner join follows f on a.userid = f.followedid left join article_likes al on a.id = al.articleid  "
	}else{
		// 全てのユーザの投稿
		query1 = "SELECT a.id id,a.userid userid,u.name name,a.title title,a.tag tag,a.created created,a.updated updated ,COUNT(al.id) likecount FROM articles a inner join users u on a.userid = u.id left join article_likes al on a.id = al.articleid "
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
					// query2 = "WHERE a.id < ? AND (f.followerid = ? OR f.followedid = ?) "
					query2 = "WHERE a.id < ? AND f.followerid = ? "
					query = query1 + query2 + query3
						// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, cursor,userid); err != nil {
						return nil, err
					}
				}else{
					// query2 = "WHERE a.id < ? AND (f.followerid = ? OR f.followedid = ?) "
					query2 = "WHERE f.followerid = ? "
					query = query1 + query2 + query3
						// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = "WHERE a.id < ? "
					query = query1 + query2 + query3
						// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, cursor); err != nil {
						return nil, err
					}
				}else{
					query = query1 + query3
						// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
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
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, cursor, scope.Text,userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.title LIKE CONCAT("%",?,"%") AND f.followerid = ?  `
					query = query1 + query2 + query3
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, scope.Text, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, cursor, scope.Text); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.title LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
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
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, cursor, scope.Tag, userid); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.tag LIKE CONCAT("%",?,"%") AND f.followerid = ? `
					query = query1 + query2 + query3
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, scope.Tag, userid, cursor); err != nil {
						return nil, err
					}
				}
			}else{
				if scope.Order == "new" {
					query2 = `WHERE a.id < ? AND a.tag LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, cursor, scope.Tag); err != nil {
						return nil, err
					}
				}else{
					query2 = `WHERE a.tag LIKE CONCAT("%",?,"%")`
					query = query1 + query2 + query3
					// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
					if err := db.Select(&articles, query, scope.Tag, cursor); err != nil {
						return nil, err
					}
				}
			}
			// query2 = `WHERE a.id < ? AND a.tag LIKE CONCAT("%",?,"%")`
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
	// クエリ文字列を生成します。
	query := `SELECT a.id id,a.userid userid,u.name name,a.title title,a.body body,a.tag tag,a.created created,a.updated updated
	FROM articles a,users u
	WHERE a.id = ? and a.userid = u.id;`

	// クエリ結果を格納する変数を宣言します。
	// 複数件取得の場合はスライスでしたが、一件取得の場合は構造体になります。
	var article model.Article

	// 結果を格納する構造体、クエリ文字列、パラメータを指定して SQL を実行します。
	// 一件取得の場合は db.Get()
	if err := db.Get(&article, query, id); err != nil {
		// エラーが発生した場合はエラーを返却します。
		return nil, err
	}

	// エラーがない場合は記事データを返却します。
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
  // 現在日時を取得します
	now := time.Now()

  // 構造体に現在日時を設定します。
  article.Created = now.Format("2006/01/02 15:04:05")
  article.Updated = now.Format("2006/01/02 15:04:05")

  // クエリ文字列を生成します。
  query := `INSERT INTO articles (userid,title, body, tag, created, updated)
  VALUES (:userid, :title, :body, :tag, :created, :updated);`

  // トランザクションを開始します。
  tx := db.MustBegin()

  // クエリ文字列と構造体を引数に渡して SQL を実行します。
  // クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換されます。
  // 構造体タグで指定してあるフィールドが対象となります。（`db:"title"` など）
  res, err := tx.NamedExec(query, article)
  if err != nil {
    // エラーが発生した場合はロールバックします。
    tx.Rollback()

    // エラー内容を返却します。
    return nil, err
  }

  // SQL の実行に成功した場合はコミットします。
  tx.Commit()

  // SQL の実行結果を返却します。
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
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, nil
}

// ArticleDelete ...
func ArticleDelete(id int) error {
	query := "DELETE FROM articles WHERE id = ?"
	tx := db.MustBegin()
	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
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
	LIMIT 10`
	articles := make([]*model.Article, 0, 10)
	if err := db.Select(&articles, query, cursor,userid); err != nil {
		return nil, err
	}

	return articles, nil
}