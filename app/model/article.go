package model

import (
	"gopkg.in/go-playground/validator.v9"
)

// Article ...
type Article struct {
	ID      int       `db:"id" json:"id"`
	Userid  int       `db:"userid" json:"userid"`
	UserName string   `db:"name" json:"name"`
	KBCMail    string    `db:"mail" json:"KBC_mail"`
	ImgPath      string    `db:"imgpath" json:"imgpath"`
	Sex         int       `db:"sex" json:"sex"`
	Title   string    `db:"title" json:"title" validate:"required,max=50"`
	Body    string    `db:"body" json:"body" validate:"required"`
	Tag    string     `db:"tag" json:"tag" validate:"required"`
	LikeCount int     `db:"likecount" json:"likecount"`
	Created string `db:"created"`
	Updated string `db:"updated"`
}

// FavoriteArticle ...
type FavoriteArticle struct {
	ID      int       `db:"t_id" json:"id"`
	Userid  int       `db:"userid" json:"userid"`
	UserName string   `db:"name" json:"name"`
	ImgPath      string    `db:"imgpath" json:"imgpath"`
	Title   string    `db:"title" json:"title" validate:"required,max=50"`
	Tag    string     `db:"tag" json:"tag" validate:"required"`
	LikeCount int     `db:"likecount" json:"likecount"`
	LikeID   int      `db:"likeid" json:"likeid"`
	Updated string `db:"updated"`
}

// ValidationErrors ...
func (a *Article) ValidationErrors(err error) []string {
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
		case "Tag":
			message = "タグが1つも選択されていません。"
		}

		// メッセージをスライスに追加します。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}
