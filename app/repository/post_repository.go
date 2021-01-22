package repository

import (
"app/model"
)

// PostCount 投稿数を取得
func PostCount(user *model.User) error{
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

	if err := db.Select(&tags, query,userid); err != nil {
		return nil, err
	}

	return tags, nil
}