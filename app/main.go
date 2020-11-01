package main

import (
	"log"
	// "net/http"
	// "time"
	"app/handler"
	"app/repository"
	"app/model"
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	e.POST("/Users",handler.UserCreate)
	e.POST("/Login",handler.Login)
	// e.POST("/session",writeCookie)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/Articles", handler.ArticleIndex)
	r.GET("/Articles/new", handler.ArticleNew)
	r.GET("/Articles/:id", handler.ArticleShow)
	r.GET("/Articles/:id/edit", handler.ArticleEdit)
	r.POST("/Articles", handler.ArticleCreate)

	// Webサーバーをポート番号 8082 で起動する
	e.Logger.Fatal(e.Start(":8082"))
	// e.Startの中はdocker-composeのgoコンテナで設定したportsを指定してください。
}

func connectDB() *sqlx.DB {
	dsn := "root:kbc_password@tcp(myapp-mysql:3306)/kbcblog-mysql?parseTime=true&autocommit=0&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

func createMux() *echo.Echo {
	// アプリケーションインスタンスを生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	//アプリケーションのどこかで予期せずにpanicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーする
	e.Use(middleware.Recover())
	//アクセスログのようなリクエスト単位のログを出力してくれます。
	e.Use(middleware.Logger())
	//gzip圧縮スキームを使用してHTTP応答を圧縮します。
	e.Use(middleware.Gzip())

	e.Validator = &CustomValidator{validator: validator.New()}

	// &CustomValidator.validator.RegisterValidation("mail_check_regexp",model.Mail_check_regexp)
	// &CustomValidator.validator.RegisterValidation("password_check",model.Password_check)

	// アプリケーションインスタンスを返却
	return e
}

// CustomValidator ...
// type Jwt struct {
// 	Token        string    `json:"JWT"`
// }

// func writeCookie(c echo.Context) error {
// 	var jwt Jwt

// 	if err := c.Bind(&jwt); err != nil {
//     // エラーの内容をサーバーのログに出力します。
//     c.Logger().Error(err.Error())

//     // リクエストの解釈に失敗した場合は 400 エラーを返却します。
//     return c.JSON(http.StatusBadRequest, "NO write a cookie")
// 	}

// 	cookie := new(http.Cookie)
// 	cookie.Name = "JWT"
// 	cookie.Value = jwt.Token
// 	cookie.Expires = time.Now().Add(1 * time.Hour)
// 	c.SetCookie(cookie)
// 	return c.JSON(http.StatusOK, "write a cookie")
// }

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	cv.validator.RegisterValidation("mail_check_regexp",model.Mail_check_regexp)
	return cv.validator.Struct(i)
}
