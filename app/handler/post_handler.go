package handler

import (
	"net/http"
	"strconv"
	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// GetFavoritePosts ...
func GetFavoritePosts(c echo.Context) error {
	userID := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repository.UserGetByID(id)

	articles, err := repository.GetFavoriteArticle(0, user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザの記事一覧データを取得中にエラー発生") //500
	}

	var articleCursor int
	if len(articles) != 0 {
		articleCursor = articles[len(articles)-1].LikeID
	}

	questions, err := repository.GetFavoriteQuestion(0, user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザの質問一覧データを取得中にエラー発生") //500
	}

	answerQuestions, err := repository.GetFavoriteAnswerQuestion(0, user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザの回答一覧データを取得中にエラー発生") //500
	}

	// Follow情報を取得
	var follow model.Follow
	if userID != user.ID {
		count, err := repository.CheckFollow(userID, user.ID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "Follow情報取得中にエラー発生") //500
		}
		// ユーザがfollowしているか検証
		if count > 0 {
			follow.IsFollow = true
		}
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"MyUser":            myUser,
		"User":              user,
		"Articles":          articles,
		"ArticleCursor":     articleCursor,
		"Follow":            follow,
		"Questions":         questions,
		"AnswerQuestions":   answerQuestions,
	}

	return c.JSON(http.StatusOK, data)
}

// GetFavoriteArticles 記事一覧のスクロール時に次の10記事を返す
func GetFavoriteArticles(c echo.Context) error {

	cursor, _ := strconv.Atoi(c.QueryParam("articleCursor"))
	userid, _ := strconv.Atoi(c.QueryParam("userid"))

	articles, err := repository.GetFavoriteArticle(cursor,userid)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"記事の一覧データ取得中にエラー発生")
	}

	if len(articles) != 0 {
		cursor = articles[len(articles)-1].LikeID
	}

	data := map[string]interface{}{
		"Articles": articles,
		"Cursor":   cursor,
	}
	return c.JSON(http.StatusOK, data)
}