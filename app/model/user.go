package model

import (
	"database/sql"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser ...
type CreateUser struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name" validate:"required,min=4,max=10"`
	KBCMail    string    `db:"mail" json:"KBC_mail" validate:"required,"`
	Password    string    `json:"password" validate:"required,min=8,max=50"`
	PassCfm    string    `json:"password_confirmation"`
	PassHash    string    `db:"passhash"`
}

// User ...
type User struct {
	ID          int       `db:"id" json:"id"`
	KBCMail    string    `db:"mail" json:"KBC_mail"`
	Name        string    `db:"name" json:"name" validate:"required,min=4,max=10"`
	Comment     sql.NullString   `db:"comment" json:"comment" validate:"max=150"`
	PostCount int      `db:"PostCount" json:"postCount"` //投稿数
	Github      sql.NullString    `db:"github" json:"github"`
	Website     sql.NullString    `db:"website" json:"website"`
	Languages   sql.NullString    `db:"languages" json:"languages"`
}

// LoginUser ...
type LoginUser struct {
	ID          int       `db:"id" json:"id"`
	KBCMail    string    `db:"mail" json:"KBC_mail"`
	PassHash    string    `db:"passhash"`
}

// RankingUser ...
type RankingUser struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Comment     sql.NullString   `db:"comment" json:"comment"`
	Count int     `db:"count" json:"count"`
}

// SetupUser CreateUserからUserを作成
func (u *User) SetupUser(cu CreateUser, id int) () {
	u.ID = cu.ID
	u.Name = cu.Name
}

// PasswordHash パスワードハッシュを作る
func (u *CreateUser) PasswordHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	u.PassHash = string(hash)
	return string(hash), err
}

// ValidationErrors ...
func (u *CreateUser) ValidationErrors(err error) []string {
	// メッセージを格納するスライスを宣言します。
	var errMessages []string

	// 複数のエラーが発生する場合があるのでループ処理を行います。
	for _, err := range err.(validator.ValidationErrors) {
		// メッセージを格納する変数を宣言します。
		var message string

		// エラーになったフィールドを特定します。
		switch err.Field() {
		case "Name":
			switch err.Tag() {
			case "required":
				message = "名前を入力してください。"
			case "min":
				message = "名前は4文字以上です。"
			case "max":
				message = "名前は最大10文字です。"
			}
		case "KBCMail":
			switch err.Tag() {
			case "required":
				message = "メールアドレスを入力してください。"
			case "MailCheckRegexp":
				message = "河原学園のメールアドレスを入力してください。"
			}
		case "Password":
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
	if u.PasswordCheck() != true{
		message := "パスワードとパスワード確認が異なります。"
		errMessages = append(errMessages, message)
	}
	return errMessages
}

// ValidationErrors ...
func (u *User) ValidationErrors(err error) []string {
	var errMessages []string
	// 複数のエラーが発生する場合があるのでループ処理を行います。
	for _, err := range err.(validator.ValidationErrors) {
		// メッセージを格納する変数を宣言します。
		var message string

		// エラーになったフィールドを特定します。
		switch err.Field() {
		case "Name":
			switch err.Tag() {
			case "required":
				message = "名前を入力してください。"
			case "min":
				message = "名前は4文字以上です。"
			case "max":
				message = "名前は最大10文字です。"
			}
		case "Comment":
			message = "コメントは最大150文字です。"
		}
		// メッセージをスライスに追加します。
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}
	return errMessages
}

// MailCheckRegexp ...
func MailCheckRegexp(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
	reg := `[\w\-\._]+@(stu.)?kawahara.ac.jp`
	str := fl.Field().String()
	return regexp.MustCompile(reg).Match([]byte(str))
}

// PasswordCheck ...
func (u *CreateUser) PasswordCheck() bool{
	if(u.Password == u.PassCfm){
		return true
	}
	return false
}