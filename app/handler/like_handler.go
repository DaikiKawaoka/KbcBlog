package handler

import (
	"net/http"
	"strconv"
	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// like

// CreateArticleLike 記事いいね処理
func CreateArticleLike(c echo.Context) error {
	userID := userIDFromToken(c)
	articleID, _ := strconv.Atoi(c.Param("id"))
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}

	var like model.ArticleLike
	like.Userid = userID
	like.Articleid = articleID
	// Like作成
	err := repository.CreateArticleLike(&like)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}

	// 通知作成
	var notification model.CreateNotification
	visitedid,err := repository.ArticleGetByID2(articleID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError, "ユーザID取得中にエラー") //500
	}
	notification.Articleid = articleID
	notification.Visiterid = userID
	notification.Visitedid = visitedid
	notification.Action = "alike"
	if notification.Visitedid != notification.Visiterid{
		if err := repository.NotificationCreate(&notification); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "通知作成失敗")
		}
	}

	return c.JSON(http.StatusOK, "ARTICLE LIKE CREATE 成功")
}

// DeleteArticleLike 記事いいね処理
func DeleteArticleLike(c echo.Context) error {
	userID := userIDFromToken(c)
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}
	articleID, _ := strconv.Atoi(c.Param("id"))
	// Like削除
	err := repository.DeleteArticleLike(userID, articleID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "ARTICLE LIKE DELETE 成功")
}

// CreateQuestionLike 質問いいね処理
func CreateQuestionLike(c echo.Context) error {
	userID := userIDFromToken(c)
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}
	questionID, _ := strconv.Atoi(c.Param("id"))
	var like model.QuestionLike
	like.Userid = userID
	like.Questionid = questionID
	// Like作成
	err := repository.CreateQuestionLike(&like)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}

		// 通知作成
	var notification model.CreateNotification
	visitedid,err := repository.QuestionGetByID2(questionID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError, "ユーザID取得中にエラー") //500
	}
	notification.Questionid = questionID
	notification.Visiterid = userID
	notification.Visitedid = visitedid
	notification.Action = "qlike"
	if notification.Visitedid != notification.Visiterid{
		if err := repository.NotificationCreate(&notification); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "通知作成失敗")
		}
	}

	return c.JSON(http.StatusOK, "QUESTION LIKE CREATE 成功")
}

// DeleteQuestionLike 質問いいね処理
func DeleteQuestionLike(c echo.Context) error {
	userID := userIDFromToken(c)
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}
	questionID, _ := strconv.Atoi(c.Param("id"))
	// Like削除
	err := repository.DeleteQuestionLike(userID, questionID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "QUESTION LIKE DELETE 成功")
}



// CreateArticleCommentLike 記事コメントいいね処理
func CreateArticleCommentLike(c echo.Context) error {
	userID := userIDFromToken(c)
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}
	articleCommentID, _ := strconv.Atoi(c.Param("id"))
	var like model.ArticleCommentLike
	like.Userid = userID
	like.ArticleCommentid = articleCommentID
	// Like作成
	err := repository.CreateArticleCommentLike(&like)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}

	// 通知作成
	var notification model.CreateNotification
	if err := c.Bind(&notification); err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusBadRequest, "代入失敗")
	}
	notification.Visiterid = userID
	notification.ACommentid = articleCommentID
	notification.Action = "aclike"
	if notification.Visitedid != notification.Visiterid{
		if err := repository.NotificationCreate(&notification); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "通知作成失敗")
		}
	}

	return c.JSON(http.StatusOK, "ARTICLE COOMENT LIKE CREATE 成功")
}

// DeleteArticleCommentLike 記事コメントいいね処理
func DeleteArticleCommentLike(c echo.Context) error {
	userID := userIDFromToken(c)
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}
	articleCommentID, _ := strconv.Atoi(c.Param("id"))
	// Like削除
	err := repository.DeleteArticleCommentLike(userID, articleCommentID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "ARTICLE COMMENT LIKE DELETE 成功")
}



// CreateQuestionCommentLike 質問コメントいいね処理
func CreateQuestionCommentLike(c echo.Context) error {
	userID := userIDFromToken(c)
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}
	questionCommentID, _ := strconv.Atoi(c.Param("id"))
	var like model.QuestionCommentLike
	like.Userid = userID
	like.QuestionCommentid = questionCommentID
	// Like作成
	err := repository.CreateQuestionCommentLike(&like)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生") //500
	}

	// 通知作成
	var notification model.CreateNotification
	if err := c.Bind(&notification); err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusBadRequest, "代入失敗")
	}
	notification.Visiterid = userID
	notification.QCommentid = questionCommentID
	notification.Action = "qclike"
	if notification.Visitedid != notification.Visiterid{
		if err := repository.NotificationCreate(&notification); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "通知作成失敗")
		}
	}

	return c.JSON(http.StatusOK, "QUESTION COOMENT LIKE CREATE 成功")
}


// DeleteQuestionCommentLike 質問コメントいいね処理
func DeleteQuestionCommentLike(c echo.Context) error {
	userID := userIDFromToken(c)
	if userID == 920437694{
		return c.JSON(http.StatusInternalServerError,"ゲストユーザーではこの処理はできません。") //500
	}
	questionCommentID, _ := strconv.Atoi(c.Param("id"))
	// Like削除
	err := repository.DeleteQuestionCommentLike(userID, questionCommentID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "QUESTION COMMENT LIKE DELETE 成功")
}