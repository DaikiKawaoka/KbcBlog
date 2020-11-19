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



// Article Comment Like

func ArticleCommentLike(c echo.Context) error {
	userId := userIDFromToken(c)
	articleCommentId, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetArticleCommentLike(userId,articleCommentId)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteArticleCommentLike(userId, articleCommentId)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "ARTICLE COMMENT LIKE DELETE 成功")
	}else{
		var like model.ArticleCommentLike
		like.Userid = userId
		like.ArticleCommentid = articleCommentId
		// Like作成
		err := repository.CreateArticleCommentLike(&like)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "ARTICLE COOMENT LIKE CREATE 成功")
	}
}



// Question Comment Like

func QuestionCommentLike(c echo.Context) error {
	userId := userIDFromToken(c)
	questionCommentId, _ := strconv.Atoi(c.Param("id"))
	count, err := repository.GetQuestionCommentLike(userId,questionCommentId)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// Like削除
		err := repository.DeleteQuestionCommentLike(userId, questionCommentId)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "QUESTION COMMENT LIKE DELETE 成功")
	}else{
		var like model.QuestionCommentLike
		like.Userid = userId
		like.QuestionCommentid = questionCommentId
		// Like作成
		err := repository.CreateQuestionCommentLike(&like)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "QUESTION COOMENT LIKE CREATE 成功")
	}
}