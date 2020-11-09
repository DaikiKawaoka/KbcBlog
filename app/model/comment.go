package model

import (
	"gopkg.in/go-playground/validator.v9"
)

// Article ...
type Comment struct {
	ID          int       `db:"id" json:"id"`
	Userid      int       `db:"userid" form:"userid" json:"userid"`
	Articleid   int       `db:"articleid" json:"articleid"`
	UserName    string    `db:"name" json:"name"`
	Text        string    `db:"text" form:"text" json:"text" validate:"required"`
	Created     string `db:"created"`
	Updated     string `db:"updated"`
}

// ValidationErrors ...
func (a *Comment) ValidationErrors(err error) []string {
	// メッセージを格納するスライスを宣言します。
	var errMessages []string

	// 複数のエラーが発生する場合があるのでループ処理を行います。
	for _, err := range err.(validator.ValidationErrors) {
		// メッセージを格納する変数を宣言します。
		var message string

		// エラーになったフィールドを特定します。
		switch err.Field() {
		case "Text":
				message = "コメントを入力してください。"
		}
		// メッセージをスライスに追加します。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}