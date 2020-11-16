package model

type ArticleLike struct {
	Userid      int       `db:"userid" form:"userid" json:"userid"`
	Articleid   int       `db:"articleid" json:"articleid"`
  Created     string    `db:"created"`
}