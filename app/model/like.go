package model

// ArticleLike ...
type ArticleLike struct {
	Userid      int       `db:"userid" form:"userid" json:"userid"`
	Articleid   int       `db:"articleid" json:"articleid"`
  Created     string    `db:"created"`
}

// QuestionLike ...
type QuestionLike struct {
	Userid      int       `db:"userid" form:"userid" json:"userid"`
	Questionid   int       `db:"questionid" json:"questionid"`
  Created     string    `db:"created"`
}

// ArticleCommentLike ...
type ArticleCommentLike struct {
	Userid             int       `db:"userid" form:"userid" json:"userid"`
	ArticleCommentid   int       `db:"article_commentid" json:"articleid"`
  Created            string    `db:"created"`
}

// QuestionCommentLike ...
type QuestionCommentLike struct {
	Userid              int       `db:"userid" form:"userid" json:"userid"`
	QuestionCommentid   int       `db:"question_commentid" json:"questionid"`
  Created             string    `db:"created"`
}

// Like ...
type Like struct {
	LikeCount   int      `json:"likeCount"`
	IsLike      bool     `json:"isLike"`
}