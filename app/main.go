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

	// Restricted group
	r := e.Group("/restricted")

	config := middleware.JWTConfig{
		Claims:     &handler.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	// Article
	r.GET("/Articles", handler.ArticleIndex)
	r.GET("/Articles/new", handler.ArticleNew)
	r.GET("/Articles/:id", handler.ArticleShow)
	r.GET("/Articles/:id/edit", handler.ArticleEdit)
	r.POST("/Articles", handler.ArticleCreate)
	r.PATCH("/Articles/:id", handler.ArticleUpdate)  // 更新
	r.DELETE("/Articles/:id", handler.ArticleDelete)
  // ArticleComment
	r.GET("/Articles/:id/Comments",handler.ArticleCommentIndex)
	r.POST("/Articles/:id/Comments",handler.ArticleCommentCreate)
	r.DELETE("/Articles/:id/Comments/:cid", handler.ArticleCommentDelete)
  // Question
	r.GET("/Questions", handler.QuestionIndex)
	r.GET("/Questions/new", handler.QuestionNew)
	r.GET("/Questions/:id", handler.QuestionShow)
	r.GET("/Questions/:id/edit", handler.QuestionEdit)
	r.POST("/Questions", handler.QuestionCreate)
	r.PATCH("/Questions/:id", handler.QuestionUpdate)
	r.DELETE("/Questions/:id", handler.QuestionDelete)
  // QuestionComment
	r.GET("/Questions/:id/Comments",handler.QuestionCommentIndex)
	r.POST("/Questions/:id/Comments",handler.QuestionCommentCreate)
	r.DELETE("/Questions/:id/Comments/:cid", handler.QuestionCommentDelete)

	// User
	r.GET("/Users/:id",handler.UserShow)
	r.GET("/Users/:id/edit", handler.UserEdit)
	r.PATCH("/Users/:id", handler.UserUpdate)

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

	// アプリケーションインスタンスを返却
	return e
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	cv.validator.RegisterValidation("mail_check_regexp",model.Mail_check_regexp)
	return cv.validator.Struct(i)
}
