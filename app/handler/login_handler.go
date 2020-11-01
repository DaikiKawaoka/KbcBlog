package handler

import (
	// "log"
	"net/http"
	// "strconv"
	"time"

	"app/repository"
	// "app/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	KBC_mail    string    `json:"KBC_mail"`
	Password    string    `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Message          string
  ValidationErrors []string
}

func Login(c echo.Context) error {

	var loginReq LoginRequest

  if err := c.Bind(&loginReq); err != nil {
    c.Logger().Error(err.Error())
    // リクエストの解釈に失敗した場合は 400 エラーを返却します。
    return c.JSON(http.StatusBadRequest, err)
	}

	// ユーザ認証

	loginUser, err := repository.GetLoginPassHash(loginReq.KBC_mail)

	if  err != nil{
		// エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())
		// サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
		message := "ユーザが存在しません。"
    return c.JSON(http.StatusInternalServerError, message)
	}

	//パスワードとパスワードハッシュが一致しているか検証
	if err := passwordVerify(loginUser.PassHash, loginReq.Password); err != nil{
    c.Logger().Error(err.Error())
		message := "パスワードが違います。"
    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
    return c.JSON(http.StatusInternalServerError, message)
	}

	// Generate encoded token and send it as response.
	var loginRes LoginResponse
	t, err := CreateToken(loginUser.ID)
	if err != nil {
		return err
	}
	loginRes.Token = t

	// response
	return c.JSON(http.StatusOK,loginRes)
}

// 登録したパスワードハッシュが入力パスワードにマッチするかどうかを調べる関数
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func CreateToken(userID int) (string, error) {
	// tokenの作成
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsの設定
	token.Claims = jwt.MapClaims{
			"user": userID,
			"exp":  time.Now().Add(time.Hour * 1).Unix(), // 有効期限を指定
	}

	// 署名
	var secretKey = "secret" // 任意の文字列
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
			return "", err
	}
	return tokenString, nil
}