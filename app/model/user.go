package model

import (
	"gopkg.in/go-playground/validator.v9"
	"time"
  "regexp"
)

type User struct {
	ID          int       `db:"id" json:"id"`
	KBC_mail    string    `db:"mail" json:"KBC_mail" validate:"required,mail_check_regexp"`
	Password    string    `json:"password" validate:"min=8,max=50"`
	PassHash    string    `db:"passhash"`
	MailName    string    `db:"mailname" json:"mailname"`
	Name        string    `db:"name" json:"name" validate:"required,min=4,max=30"`
	Created     time.Time `db:"created"`
	Updated     time.Time `db:"updated"`
}

// ValidationErrors ...
func (a *User) ValidationErrors(err error) []string {


	// メッセージを格納するスライスを宣言します。
	var errMessages []string

	// 複数のエラーが発生する場合があるのでループ処理を行います。
	for _, err := range err.(validator.ValidationErrors) {
		// メッセージを格納する変数を宣言します。
		var message string

		// エラーになったフィールドを特定します。
		switch err.Field() {
		case "KBC_mail":
			switch err.Tag() {
			case "required":
				message = "メールアドレスを入力してください。"
			case "mail_check_regexp":
				message = "河原学園のメールアドレスを入力してください。"
			}
		case "Password":
			// エラーになったバリデーションルールを特定します。
			switch err.Tag() {
			case "required":
				message = "パスワードは8文字以上です。"
			case "max":
				message = "タイトルは最大50文字です。"
			}
		case "Name":
			switch err.Tag() {
			case "min":
				message = "名前は4文字以上です。"
			case "max":
				message = "名前は最大30文字です。"
			}
		}
		// メッセージをスライスに追加します。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}

func mail_check_regexp(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
	reg := `[\w\-\._]+@(stu.)?kawahara.ac.jp`
	str := fl.Field().String()
	return regexp.MustCompile(reg).Match([]byte(str))
}