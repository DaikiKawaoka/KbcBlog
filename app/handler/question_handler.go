package handler

import (
	// "log"
	"net/http"
	"strconv"
	"strings"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// QuestionCreateOutput 質問作成リターン型
type QuestionCreateOutput struct {
  Question          *model.Question
  Message          string
  ValidationErrors []string
}

// QuestionCreate 質問作成処理
func QuestionCreate(c echo.Context) error {

  var question model.Question
	var out QuestionCreateOutput

  if err := c.Bind(&question); err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusBadRequest, out) //400
	}
  if err := c.Validate(&question); err != nil {
    c.Logger().Error(err.Error())
    out.ValidationErrors = question.ValidationErrors(err)
    return c.JSON(http.StatusUnprocessableEntity, out) //422
  }

  res, err := repository.QuestionCreate(&question)
  if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError, out) //500
  }
  // SQL 実行結果から作成されたレコードの ID を取得します。
  id, _ := res.LastInsertId()
  // 構造体に ID をセットします。
  question.ID = int(id)
  out.Question = &question
  return c.JSON(http.StatusOK, out)
}

// QuestionIndex 質問一覧ページ
func QuestionIndex(c echo.Context) error {

	userID := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません") //500
	}
	questions, err := repository.QuestionListByCursor(0)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"質問の一覧データを取得中にエラー発生") //500
	}
	data := map[string]interface{}{
		"user": myUser,
		"Questions": questions,
	}
	return c.JSON(http.StatusOK, data)
}

// QuestionShow 質問詳細ページ
func QuestionShow(c echo.Context) error {

	userID := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	question, err := repository.QuestionGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// 記事データのいいねを取得する
	var like model.Like
	count, err := repository.GetQuestionLike(userID,id)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	// ユーザがいいねしているか検証
	if count > 0{
		like.IsLike = true
	}

	like.LikeCount,err = repository.GetQuestionLikeCount(id)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータのカウント数取得中にエラー発生") //500
	}

	// 記事データへのコメント一覧データを取得します。
	comments, err := repository.QuestionCommentListByCursor(question.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"質問のコメント一覧データを取得中にエラー発生") //500
	}

	var commentLike model.Like
	for i, comment := range comments {
		count, err := repository.GetQuestionCommentLike(userID,comment.ID);
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
		}
		// ユーザがいいねしているか検証
		if count > 0{
			commentLike.IsLike = true
		}else{
			commentLike.IsLike = false
		}
		commentLike.LikeCount,err = repository.GetQuestionCommentLikeCount(comment.ID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"likeデータのカウント数取得中にエラー発生") //500
		}
		comments[i].Like = commentLike
	}

	data := map[string]interface{}{
		"user": myUser,
		"Question": question,
		"Like": like,
		"Comments": comments,
	}
	return c.JSON(http.StatusOK, data)
}

// QuestionNew 質問作成ページ
func QuestionNew(c echo.Context) error {
	userID := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}
	data := map[string]interface{}{
		"user": myUser,
	}
	return c.JSON(http.StatusOK, data)
}

// QuestionEdit 質問編集ページ
func QuestionEdit(c echo.Context) error {
	userID := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}
	id, _ := strconv.Atoi(c.Param("id"))
	question, err := repository.QuestionGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError) //500
	}
	data := map[string]interface{}{
		"user": myUser,
		"Question": question,
	}
	return c.JSON(http.StatusOK, data)
}

// QuestionUpdateOutput 質問編集リターン型
type QuestionUpdateOutput struct {
	Question         *model.Question
	Message          string
	ValidationErrors []string
}

// QuestionUpdate 質問編集処理
func QuestionUpdate(c echo.Context) error {
	ref := c.Request().Referer()
	refID := strings.Split(ref, "/")[4]
	reqID := c.Param("id")

	if reqID != refID {
		return c.JSON(http.StatusBadRequest, "パス違い") //400
	}

	var question model.Question
	var out QuestionUpdateOutput
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, out) //400
	}

	// 入力値のチェック
	if err := c.Validate(&question); err != nil {
		// エラー内容をレスポンスのフィールドに格納します。
		out.ValidationErrors = question.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out) //422
	}

	questionID, _ := strconv.Atoi(reqID)
	question.ID = questionID
	_, err := repository.QuestionUpdate(&question)

	if err != nil {
		out.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, out) //500
	}
	out.Question = &question
	return c.JSON(http.StatusOK, out)
}

//QuestionDelete 質問削除処理
func QuestionDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.QuestionDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "") //500
	}
	message := "質問を削除しました。"
	return c.JSON(http.StatusOK, message)
}