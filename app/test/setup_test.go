package test

import (
	"log"
	"os"
	"fmt"
	"app/handler"
	"app/model"
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
	"github.com/joho/godotenv" //環境変数
)

func connectDB() *sqlx.DB {
	err := godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))
    if err != nil {
			// e.Logger.Fatal(err)
		}
	var (
		DBPASSWORDROOT = os.Getenv("DB_PASSWORD_ROOT_TEST")
		DBUSERNAME = os.Getenv("DB_USERNAME_TEST")
		DBDATABASE = os.Getenv("DB_DATABASE_TEST")
	)
	dsn := DBUSERNAME+":"+DBPASSWORDROOT+"@tcp(127.0.0.1:3307)/"+DBDATABASE+"?parseTime=true&autocommit=0&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		// e.Logger.Fatal(err)
		log.Println("err1")
	}
	if err := db.Ping(); err != nil {
		log.Println("err2")
	}
	log.Println("db connection succeeded test")
	return db
}

func createMux() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	// e.Use(middleware.CORS())
	e.POST("/Users",handler.UserCreate)
	e.POST("/Login",handler.Login)
	e.GET("/guestLogin",handler.GuestLogin)
	// Restricted group
	r := e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:     &handler.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	//Static
	r.GET("/About",handler.About)
	// r.GET("/Help",handler.Help)
	// Article
	r.GET("/Articles", handler.ArticleIndex)
	r.GET("/Articles/scope", handler.ArticleIndexOrder)
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
	r.GET("/Questions/scope", handler.QuestionIndexOrder)
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
	// UserIcon
	r.PATCH("/Users/:id/img", handler.UserImgUpdate)
	r.POST("/Users/:id/img", handler.UserImgUpdate)
	r.DELETE("/Users/:id/img", handler.UserImgDelete)
	// Password
	r.PATCH("/Users/:id/password/edit",handler.UserPasswordEdit)
	// Like
	r.POST("/Articles/:id/Likes", handler.CreateArticleLike)
	r.POST("/Questions/:id/Likes", handler.CreateQuestionLike)
	r.POST("/Articles/Comments/:id/Likes",handler.CreateArticleCommentLike)
	r.POST("/Questions/Comments/:id/Likes",handler.CreateQuestionCommentLike)
	r.DELETE("/Articles/:id/Likes", handler.DeleteArticleLike)
	r.DELETE("/Questions/:id/Likes", handler.DeleteQuestionLike)
	r.DELETE("/Articles/Comments/:id/Likes",handler.DeleteArticleCommentLike)
	r.DELETE("/Questions/Comments/:id/Likes",handler.DeleteQuestionCommentLike)
	// Follow
	r.GET("/Users/:id/Following",handler.Following)
	r.GET("/Users/:id/Followers",handler.Followers)
	r.POST("/Users/:id/Follow",handler.Follow)
	r.DELETE("/Users/:id/Follow",handler.UnFollow)
	// Notificatino
	r.GET("/Notifications",handler.Notifications)
	r.GET("/Notifications/scope",handler.NotificationOrder)
	r.DELETE("/Notifications", handler.NotificationDelete)
	//Ranking
	r.GET("/Article/Ranking",handler.GetArticleRanking)
	r.GET("/Question/Ranking",handler.GetQuestionRanking)

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
  _ = cv.validator.RegisterValidation("MailCheckRegexp",model.MailCheckRegexp)
	return cv.validator.Struct(i)
}
