package model

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name" validate:"required,min=4,max=20"`
	KBC_mail    string    `db:"mail" json:"KBC_mail" validate:"required,mail_check_regexp"`
	Password    string    `json:"password" validate:"required,min=8,max=50"`
	Pass_cfm    string    `json:"password_confirmation"`
	PassHash    string    `db:"passhash"`
	MailName    string    `db:"mailname" json:"mailname"`
}

type User struct {
	ID          int       `db:"id" json:"id"`
	KBC_mail    string    `db:"mail" json:"KBC_mail"`
	MailName    string    `db:"mailname" json:"mailname"`
	Name        string    `db:"name" json:"name"`
}

type LoginUser struct {
	ID          int       `db:"id" json:"id"`
	KBC_mail    string    `db:"mail" json:"KBC_mail"`
	PassHash    string    `db:"passhash"`
}

// CreateUserからUserを作成
func (u *User) SetupUser(cu CreateUser, id int) () {
	u.ID = cu.ID
	u.MailName = cu.MailName
	u.Name = cu.Name
}


// パスワードハッシュを作る
func (u *CreateUser) PasswordHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	u.PassHash = string(hash)
	return string(hash), err
}

// MailNameを作る
func (u *CreateUser) CreateMailName(){
	mail := u.KBC_mail
	w := strings.Index(mail, "@")
	u.MailName = mail[0:w]
}

// ValidationErrors ...
func (a *CreateUser) ValidationErrors(err error) []string {


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
				message = "名前は最大20文字です。"
			}
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
	if a.Password_check() != true{
		message := "パスワードとパスワード確認が異なります。"
		errMessages = append(errMessages, message)
	}

	return errMessages
}

func Mail_check_regexp(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
	reg := `[\w\-\._]+@(stu.)?kawahara.ac.jp`
	str := fl.Field().String()
	return regexp.MustCompile(reg).Match([]byte(str))
}

func (u *CreateUser) Password_check() bool{
	if(u.Password == u.Pass_cfm){
		return true
	}
	return false
}