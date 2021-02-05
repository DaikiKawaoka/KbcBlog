package handler

import (
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

// QuestionIndex 記事一覧ページ
func QuestionIndex(c echo.Context) error {

	userID := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	var scope model.Scope
	scope.Order = c.QueryParam("order")//並び順
	scope.Tag = c.QueryParam("tag")//絞り込みタグ
	scope.Text = c.QueryParam("searchText")//絞り込みテキスト
	scope.FriendsOnly,_ = strconv.ParseBool(c.QueryParam("friendsOnly")) // フレンドのみ? true or false
	rankingType,_ := strconv.Atoi(c.QueryParam("rankingType")) // ランダムな数字 0 or 1

	questions, err := repository.QuestionListByCursor(0,&scope,userID)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"質問の一覧データを取得中にエラー発生")
	}

	var Ranking []*model.RankingUser

	// 質問数ランキング
	Ranking,err = repository.KBCQuestionRankingTop10(rankingType,"period")
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"Rankingを取得中にエラー発生")
	}

	// 取得できた最後の記事の ID をカーソルとして設定します。
	var cursor int
	if len(questions) != 0 {
		cursor = questions[len(questions)-1].ID
	}

	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"通知数取得中にエアー発生")
	}

	data := map[string]interface{}{
		"user": myUser,
		"Questions": questions,
		"Cursor":   cursor,
		"Ranking": Ranking,
		"NotificationCount": notificationCount,
	}
	return c.JSON(http.StatusOK, data)
}




// // QuestionIndex 質問一覧ページ
// func QuestionIndex(c echo.Context) error {

// 	userID := userIDFromToken(c)
// 	myUser,err := repository.GetMyUser(userID)
// 	if err != nil {
// 		c.Logger().Error(err.Error())
// 		return c.JSON(http.StatusInternalServerError,"userが存在しません") //500
// 	}
// 	questions, err := repository.QuestionListByCursor(0)
// 	if err != nil {
// 		c.Logger().Error(err.Error())
// 		return c.JSON(http.StatusInternalServerError,"質問の一覧データを取得中にエラー発生") //500
// 	}
// 	data := map[string]interface{}{
// 		"user": myUser,
// 		"Questions": questions,
// 	}
// 	return c.JSON(http.StatusOK, data)
// }

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

	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"通知数取得中にエアー発生")
	}

	data := map[string]interface{}{
		"user": myUser,
		"Question": question,
		"Like": like,
		"Comments": comments,
		"NotificationCount": notificationCount,
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

	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"通知数取得中にエアー発生")
	}

	data := map[string]interface{}{
		"user": myUser,
		"NotificationCount": notificationCount,
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
	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"通知数取得中にエアー発生")
	}
	data := map[string]interface{}{
		"user": myUser,
		"Question": question,
		"NotificationCount": notificationCount,
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

// QuestionIndexOrder 質問一覧のスクロール時に次の10記事を返す
func QuestionIndexOrder(c echo.Context) error {

	cursor, _ := strconv.Atoi(c.QueryParam("cursor"))
	userID := userIDFromToken(c)
	var scope model.Scope
	scope.Order = c.QueryParam("order")//並び順
	scope.Tag = c.QueryParam("tag")//絞り込みタグ
	scope.Text = c.QueryParam("searchText")//絞り込みテキスト
	scope.FriendsOnly,_ = strconv.ParseBool(c.QueryParam("friendsOnly")) // フレンドのみ? true or false

	// articles, err := repository.ArticleListByCursor(0,order,tag,text,friendsOnly,userId)
	questions, err := repository.QuestionListByCursor(cursor,&scope,userID)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"記事の一覧データ取得中にエラー発生")
	}

	if scope.Order == "new" {
		if len(questions) != 0 {
			cursor = questions[len(questions)-1].ID
		}
	}else{
		if len(questions) != 0 {
			cursor = cursor + 10
		}
	}

	data := map[string]interface{}{
		"Questions": questions,
		"Cursor":   cursor,
	}
	return c.JSON(http.StatusOK, data)
}

// QuestionMarkdown 記事詳細ページ
func QuestionMarkdown(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	question, err := repository.QuestionGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"記事データを取得中にエラー発生")
	}
	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Question": question,
	}
	return c.JSON(http.StatusOK, data)
}