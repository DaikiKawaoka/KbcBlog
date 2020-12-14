package model

import (
	"gopkg.in/go-playground/validator.v9"
	"golang.org/x/crypto/bcrypt"
)

// Password ...
type Password struct {
	ID                      int       `db:"id" json:"id"`
	CurrentPassword         string    `json:"currentPassword"`
	NewPassword             string    `json:"newpassword" validate:"required,min=8,max=50"`
	PasswordConfirmation    string    `json:"passwordConfirmation"`
}

// NewPasswordHash パスワードハッシュを作る
func (p *Password) NewPasswordHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.NewPassword = string(hash)
	return nil
}

// CurrentPasswordHash ...
func (p *Password) CurrentPasswordHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.CurrentPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.CurrentPassword = string(hash)
	return nil
}

// CurrentPasswordCheck 現在のパスワードがあっているかチェック
func (p *Password) CurrentPasswordCheck(dbcurrentP string) bool{
	if(p.CurrentPassword == dbcurrentP){
		return true
	}
	return false
}

// NewPasswordCheck ...
func (p *Password) NewPasswordCheck() bool{
	if(p.NewPassword == p.PasswordConfirmation){
		return true
	}
	return false
}

// ValidationErrors ...
func (p *Password) ValidationErrors(err error) []string {
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
			case "passwordCheck":
				message = "パスワードとパスワード確認が異なります。"
			}
		}
		// メッセージをスライスに追加します。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}
	//password-check
	if p.NewPasswordCheck() != true{
		message := "パスワードとパスワード確認が異なります。"
		errMessages = append(errMessages, message)
	}
	return errMessages
}