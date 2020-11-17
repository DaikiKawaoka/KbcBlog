package repository

import (
"app/model"
// "math"
)

// 投稿数を取得
func PostCount(user *model.User) error{
	query := `SELECT COUNT(*) AS PostCount
						FROM
						(
							SELECT * FROM articles WHERE userid = ?
							UNION
							SELECT * FROM questions WHERE userid = ?
							) AS posttable`;
	if err := db.Get(user, query, user.ID, user.ID); err != nil {
		return err
	}
	return nil
}

// func GetMyCommentPosts(cursor int,userid int) error{
// 	// 引数で渡されたカーソルの値が 0 以下の場合は、代わりに int 型の最大値で置き換えます。
// 	if cursor <= 0 {
// 		cursor = math.MaxInt32
// 	}

// 	// ID の降順に記事データを 10 件取得するクエリ文字列を生成します。
// 	query := `SELECT a.id id,a.userid userid,u.name name,a.title title,a.body body,a.created created,a.updated updated
// 	FROM articles a,users u
// 	WHERE a.id < ? and a.userid = u.id and a.userid = ?
// 	ORDER BY a.id desc
// 	LIMIT 10`

// 	// クエリ結果を格納するスライスを初期化します。
// 	// 10 件取得すると決まっているため、サイズとキャパシティを指定しています。
// 	articles := make([]*model.Article, 0, 10)

// 	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
// 	if err := db.Select(&articles, query, cursor,userid); err != nil {
// 		return nil, err
// 	}

// 	return articles, nil
// }