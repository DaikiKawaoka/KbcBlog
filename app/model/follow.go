package model

import (
	"database/sql"
)

// フォローする人は「follower」
// フォローされる人は「followed」
type Following struct {
	ID                 int       `db:"id" json:"id"`
	FollowerId         int       `db:"followerid" json:"followerid"`
	FollowedId         int       `db:"followedid" json:"followedid"`
  Created            string    `db:"created"`
}

type Follow struct {
	IsFollow              bool     `json:"isfollow"`
	FollowerCount         int       `db:"followercount" json:"followerCount"`
	FollowedCount         int       `db:"followedcount" json:"followedCount"`
}

type FUser struct {
	ID          int              `db:"id" json:"id"`
	Name        string           `db:"name" json:"name" validate:"required,min=4,max=20"`
	Comment     sql.NullString   `db:"comment" json:"comment" validate:"max=150"`
	IsFollowing bool             `db:"isfollowing" json:"isfollowing"`
	// PostCount int      `db:"PostCount" json:"postCount"` //投稿数
	// Github      sql.NullString    `db:"github" json:"github"`
	// Website     sql.NullString    `db:"website" json:"website"`
	// Languages   sql.NullString    `db:"languages" json:"languages"`
}