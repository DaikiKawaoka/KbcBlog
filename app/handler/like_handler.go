package handler

import (
	"net/http"
	"strconv"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// like

// ArticleLike 記事いいね処理
func ArticleLike(c echo.Context) error {
	userID := userIDFromToken(c)
	articleID, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetArticleLike(userID,articleID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteArticleLike(userID, articleID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "ARTICLE LIKE DELETE 成功")
	}
	var like model.ArticleLike
	like.Userid = userID
	like.Articleid = articleID
	// Like作成
	err2 := repository.CreateArticleLike(&like)
	if err2 != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "ARTICLE LIKE CREATE 成功")
}

// QuestionLike 質問いいね処理
func QuestionLike(c echo.Context) error {
	userID := userIDFromToken(c)
	questionID, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetQuestionLike(userID,questionID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteQuestionLike(userID, questionID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "QUESTION LIKE DELETE 成功")
	}
	var like model.QuestionLike
	like.Userid = userID
	like.Questionid = questionID
	// Like作成
	err2 := repository.CreateQuestionLike(&like)
	if err2 != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "QUESTION LIKE CREATE 成功")
}



// ArticleCommentLike 記事コメントいいね処理
func ArticleCommentLike(c echo.Context) error {
	userID := userIDFromToken(c)
	articleCommentID, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetArticleCommentLike(userID,articleCommentID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteArticleCommentLike(userID, articleCommentID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "ARTICLE COMMENT LIKE DELETE 成功")
	}
	var like model.ArticleCommentLike
	like.Userid = userID
	like.ArticleCommentid = articleCommentID
	// Like作成
	err2 := repository.CreateArticleCommentLike(&like)
	if err2 != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "ARTICLE COOMENT LIKE CREATE 成功")
}



// QuestionCommentLike 質問コメントいいね処理
func QuestionCommentLike(c echo.Context) error {
	userID := userIDFromToken(c)
	questionCommentID, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetQuestionCommentLike(userID,questionCommentID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteQuestionCommentLike(userID, questionCommentID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "QUESTION COMMENT LIKE DELETE 成功")
	}
	var like model.QuestionCommentLike
	like.Userid = userID
	like.QuestionCommentid = questionCommentID
	// Like作成
	err2 := repository.CreateQuestionCommentLike(&like)
	if err2 != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "QUESTION COOMENT LIKE CREATE 成功")
}