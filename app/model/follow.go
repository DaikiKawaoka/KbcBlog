package model

// フォローする人は「follower」
// フォローされる人は「followed」
type Follow struct {
	ID                 int       `db:"id" json:"id"`
	FollowerID         int       `db:"followerid" json:"followerid"`
	FollowedID         int       `db:"followedid" json:"followedid"`
  Created            string    `db:"created"`
}