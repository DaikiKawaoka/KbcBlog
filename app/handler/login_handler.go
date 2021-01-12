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

// LoginRequest ...
type LoginRequest struct {
	KBCMail    string    `json:"KBC_mail"`
	Password    string    `json:"password"`
}

// LoginResponse ...
type LoginResponse struct {
	Token string `json:"token"`
	Message          string
  ValidationErrors []string
}

// JwtCustomClaims ...
type JwtCustomClaims struct {
	UserID      int
	jwt.StandardClaims
}

// Login ...
func Login(c echo.Context) error {

	var loginReq LoginRequest

  if err := c.Bind(&loginReq); err != nil {
    c.Logger().Error(err.Error())
    // リクエストの解釈に失敗した場合は 400 エラーを返却します。
    return c.JSON(http.StatusBadRequest, err)
	}

	// ユーザ認証

	loginUser, err := repository.GetLoginPassHash(loginReq.KBCMail)
	var errmessages[1] string;

	if  err != nil{
    c.Logger().Error(err.Error())
		message := "ユーザが存在しません。"
		errmessages[0] = message;
    return c.JSON(http.StatusInternalServerError, errmessages) //500
	}

	//パスワードとパスワードハッシュが一致しているか検証
	if err := passwordVerify(loginUser.PassHash, loginReq.Password); err != nil{
    c.Logger().Error(err.Error())
		message := "パスワードが違います。"
		errmessages[0] = message;
    return c.JSON(http.StatusInternalServerError, errmessages)
	}

	// Generate encoded token and send it as response.
	var loginRes LoginResponse
	t, err := CreateToken(loginUser.ID)
	if err != nil {
		return err
	}
	loginRes.Token = t
	return c.JSON(http.StatusOK,loginRes)
}

// passwordVerify 登録したパスワードハッシュが入力パスワードにマッチするかどうかを調べる関数
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

// CreateToken ...
func CreateToken(userID int) (string, error) {

	claims := &JwtCustomClaims{
		userID,
		jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 6).Unix(), // 有効期限を指定
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "",err
		}
	return t, nil
}

func userIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	uid := claims.UserID
	return uid
}


// GuestLogin ...
func GuestLogin(c echo.Context) error {

	guestUser, err := repository.GetGuestUser()
	var errmessages[1] string;
	if  err != nil{
    c.Logger().Error(err.Error())
		message := "ユーザが存在しません。"
		errmessages[0] = message;
    return c.JSON(http.StatusInternalServerError, errmessages) //500
	}

	// Generate encoded token and send it as response.
	var loginRes LoginResponse
	t, err := CreateToken(guestUser.ID)
	if err != nil {
		print("aaaaaaa")
		return err
	}
	loginRes.Token = t
	return c.JSON(http.StatusOK,loginRes)
}