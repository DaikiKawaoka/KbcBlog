package model

import (
	"gopkg.in/go-playground/validator.v9"
)

// Question ...
type Question struct {
	ID      int       `db:"id" json:"id"`
	Userid  int       `db:"userid" json:"userid"`
	UserName string   `db:"name" json:"name"`
	Title   string    `db:"title" json:"title" validate:"required,max=50"`
	Body    string    `db:"body" json:"body" validate:"required"`
	LikeCount int     `db:"likecount" json:"likecount"`
	Created string `db:"created"`
	Updated string `db:"updated"`
}

// AnswerQuestion ...
type AnswerQuestion struct {
	ID      int       `db:"id" json:"id"`
	Userid  int       `db:"userid" json:"userid"`
	UserName string   `db:"name" json:"name"`
	Title   string    `db:"title" json:"question_title"`
	Text    string    `db:"text" json:"comment_text"`
	Created string    `db:"created" json:"comment_created"`
}

// ValidationErrors ...
func (q *Question) ValidationErrors(err error) []string {
	// メッセージを格納するスライスを宣言します。
	var errMessages []string

	// 複数のエラーが発生する場合があるのでループ処理を行います。
	for _, err := range err.(validator.ValidationErrors) {
		// メッセージを格納する変数を宣言します。
		var message string

		// エラーになったフィールドを特定します。
		switch err.Field() {
		case "Title":
			// エラーになったバリデーションルールを特定します。
			switch err.Tag() {
			case "required":
				message = "タイトルを入力してください。"
			case "max":
				message = "タイトルは最大50文字です。"
			}
		case "Body":
			message = "本文を入力してください。"
		}

		// メッセージをスライスに追加します。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}