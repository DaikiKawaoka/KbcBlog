package model

import (
	"gopkg.in/go-playground/validator.v9"
)

// ArticleComment ...
type ArticleComment struct {
	ID          int       `db:"id" json:"id"`
	Userid      int       `db:"userid" form:"userid" json:"userid"`
	Articleid   int       `db:"articleid" json:"articleid"`
	UserName    string    `db:"name" json:"name"`
	Text        string    `db:"text" form:"text" json:"text" validate:"required"`
	Like   Like
	Created     string    `db:"created"`
	Updated     string    `db:"updated"`
}

// ValidationErrors ...
func (a *ArticleComment) ValidationErrors(err error) []string {
	var errMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Field() {
		case "Text":
				message = "コメントを入力してください。"
		}
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}
	return errMessages
}

//QuestionComment ...
type QuestionComment struct {
	ID          int       `db:"id" json:"id"`
	Userid      int       `db:"userid" form:"userid" json:"userid"`
	Questionid   int      `db:"questionid" json:"questionid"`
	UserName    string    `db:"name" json:"name"`
	Text        string    `db:"text" form:"text" json:"text" validate:"required"`
	Like   Like
	Created     string    `db:"created"`
	Updated     string    `db:"updated"`
}

// ValidationErrors ...
func (a *QuestionComment) ValidationErrors(err error) []string {
	var errMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Field() {
		case "Text":
				message = "コメントを入力してください。"
		}
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}