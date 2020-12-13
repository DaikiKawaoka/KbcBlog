package handler

import (
	"net/http"
	"strconv"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

//  article

// ArticleCommentCreateOutput 記事コメント作成リターン型
type ArticleCommentCreateOutput struct {
  Comment          *model.ArticleComment
  Message          string
  ValidationErrors []string
}

// ArticleCommentCreate 記事コメント作成処理
func ArticleCommentCreate(c echo.Context) error {
  var comment model.ArticleComment
	var out ArticleCommentCreateOutput

  if err := c.Bind(&comment); err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusBadRequest, out) //400
	}
  if err := c.Validate(&comment); err != nil {
    c.Logger().Error(err.Error())
    out.ValidationErrors = comment.ValidationErrors(err)
    return c.JSON(http.StatusUnprocessableEntity, out) //422
  }
  res, err := repository.ArticleCommentCreate(&comment)
  if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError, out) //500
	}
  id, _ := res.LastInsertId()
	comment.ID = int(id)
  out.Comment = &comment
  return c.JSON(http.StatusOK, out)
}

// ArticleCommentIndex 記事コメントリスト取得処理
func ArticleCommentIndex(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	comments, err := repository.ArticleCommentListByCursor(id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"コメント一覧を取得中にエラー発生") //500
	}
	data := map[string]interface{}{
		"Comments": comments,
	}
	return c.JSON(http.StatusOK, data)
}

// ArticleCommentDelete 記事コメント削除処理
func ArticleCommentDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("cid"))
	if err := repository.ArticleCommentDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "コメント削除中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "コメントを削除しました。")
}



//  question

// QuestionCommentCreateOutput 質問コメント作成リターン型
type QuestionCommentCreateOutput struct {
  Comment          *model.QuestionComment
  Message          string
  ValidationErrors []string
}

// QuestionCommentCreate 質問コメント作成処理
func QuestionCommentCreate(c echo.Context) error {
  var comment model.QuestionComment
	var out QuestionCommentCreateOutput

  if err := c.Bind(&comment); err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusBadRequest, out) //400
	}
  if err := c.Validate(&comment); err != nil {
    c.Logger().Error(err.Error())
    out.ValidationErrors = comment.ValidationErrors(err)
    return c.JSON(http.StatusUnprocessableEntity, out) //422
  }
  res, err := repository.QuestionCommentCreate(&comment)
  if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError, out) //500
	}
  id, _ := res.LastInsertId()
  comment.ID = int(id)
  out.Comment = &comment
  return c.JSON(http.StatusOK, out)
}

// QuestionCommentIndex 質問コメントリスト取得処理
func QuestionCommentIndex(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	comments, err := repository.QuestionCommentListByCursor(id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"コメント一覧を取得中にエラー発生") //500
	}
	data := map[string]interface{}{
		"Comments": comments,
	}
	return c.JSON(http.StatusOK, data)
}

// QuestionCommentDelete 質問コメント削除処理
func QuestionCommentDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("cid"))
	if err := repository.QuestionCommentDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "コメント削除中にエラー発生") //500
	}
	return c.JSON(http.StatusOK, "コメントを削除しました。")
}