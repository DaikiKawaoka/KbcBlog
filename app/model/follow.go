package model

import (
	"database/sql"
)

// フォローする人は「follower」
// フォローされる人は「followed」

// Following ...
type Following struct {
	ID                 int       `db:"id" json:"id"`
	FollowerID         int       `db:"followerid" json:"followerid"`
	FollowedID         int       `db:"followedid" json:"followedid"`
  Created            string    `db:"created"`
}

// Follow ...
type Follow struct {
	IsFollow              bool      `json:"isfollow"`
	FollowerCount         int       `db:"followercount" json:"followerCount"`
	FollowedCount         int       `db:"followedcount" json:"followedCount"`
}

// FUser ...
type FUser struct {
	ID          int              `db:"id" json:"id"`
	Name        string           `db:"name" json:"name"`
	Comment     sql.NullString   `db:"comment" json:"comment"`
	ImgPath     string    `db:"imgpath" json:"imgpath"`
	Sex         int       `db:"sex" json:"sex"`
	IsFollowing bool             `db:"isfollowing" json:"isfollowing"`
}