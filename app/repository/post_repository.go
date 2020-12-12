package repository

import (
"app/model"
// "math"
)

// 投稿数を取得
func PostCount(user *model.User) error{
	// query := `SELECT COUNT(*) AS PostCount
	// 					FROM
	// 					(
	// 						SELECT * FROM articles WHERE userid = ?
	// 						UNION
	// 						SELECT * FROM questions WHERE userid = ?
	// 						) AS posttable`;
	query := `SELECT COUNT(*) AS PostCount FROM articles WHERE userid = ?`;
	if err := db.Get(user, query, user.ID); err != nil {
		return err
	}
	return nil
}

// GetArticleTags ユーザ全記事のタグ配列取得
func GetArticleTags(userid int) ([]*string, error) {

	var tags []*string
	query := `SELECT a.tag tags
	FROM articles a inner join users u on a.userid = u.id
	WHERE a.userid = u.id and a.userid = ?;`

	// クエリ結果を格納する変数、クエリ文字列、パラメータを指定してクエリを実行します。
	if err := db.Select(&tags, query,userid); err != nil {
		return nil, err
	}

	return tags, nil
}