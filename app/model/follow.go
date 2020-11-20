package model

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