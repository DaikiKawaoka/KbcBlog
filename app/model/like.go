package model

type ArticleLike struct {
	Userid      int       `db:"userid" form:"userid" json:"userid"`
	Articleid   int       `db:"articleid" json:"articleid"`
  Created     string    `db:"created"`
}

type Like struct {
	LikeCount   int      `json:"likeCount"`
	IsLike      bool     `json:"isLike"`
}