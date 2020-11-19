package handler

import (
	"net/http"
	"strconv"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// like
// Article Like

func ArticleLike(c echo.Context) error {
	userId := userIDFromToken(c)
	articleId, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetArticleLike(userId,articleId)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteArticleLike(userId, articleId)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "ARTICLE LIKE DELETE 成功")
	}else{
		var like model.ArticleLike
		like.Userid = userId
		like.Articleid = articleId
		// Like作成
		err := repository.CreateArticleLike(&like)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "ARTICLE LIKE CREATE 成功")
	}
}

// Question Like

func QuestionLike(c echo.Context) error {
	userId := userIDFromToken(c)
	questionId, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetQuestionLike(userId,questionId)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteQuestionLike(userId, questionId)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "QUESTION LIKE DELETE 成功")
	}else{
		var like model.QuestionLike
		like.Userid = userId
		like.Questionid = questionId
		// Like作成
		err := repository.CreateQuestionLike(&like)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "QUESTION LIKE CREATE 成功")
	}
}