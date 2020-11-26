package model

import (
	"gopkg.in/go-playground/validator.v9"
	// "regexp"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	ID                      int       `db:"id" json:"id"`
	CurrentPassword         string    `json:"currentPassword"`
	NewPassword             string    `json:"newpassword" validate:"required,min=8,max=50"`
	PasswordConfirmation    string    `json:"passwordConfirmation"`
}

// パスワードハッシュを作る
func (u *Password) NewPasswordHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.NewPassword = string(hash)
	return nil
}
func (u *Password) CurrentPasswordHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.CurrentPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.CurrentPassword = string(hash)
	return nil
}

// 現在のパスワードがあっているかチェック
func (u *Password) CurrentPassword_check(dbcurrentP string) bool{
	if(u.CurrentPassword == dbcurrentP){
		return true
	}
	return false
}

func (u *Password) NewPassword_check() bool{
	if(u.NewPassword == u.PasswordConfirmation){
		return true
	}
	return false
}


func (a *Password) ValidationErrors(err error) []string {
	// メッセージを格納するスライスを宣言します。
	var errMessages []string

	// 複数のエラーが発生する場合があるのでループ処理を行います。
	for _, err := range err.(validator.ValidationErrors) {
		// メッセージを格納する変数を宣言します。
		var message string

		// エラーになったフィールドを特定します。
		switch err.Field() {
		case "NewPassword":
			// エラーになったバリデーションルールを特定します。
			switch err.Tag() {
			case "required":
				message = "パスワードを入力してください。"
			case "min":
				message = "パスワードは8文字以上です。"
			case "max":
				message = "パスワードは最大50文字です。"
			case "password_check":
				message = "パスワードとパスワード確認が異なります。"
			}
		}
		// メッセージをスライスに追加します。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}
	//password-check
	if a.NewPassword_check() != true{
		message := "パスワードとパスワード確認が異なります。"
		errMessages = append(errMessages, message)
	}
	return errMessages
}