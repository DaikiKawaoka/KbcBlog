package handler

import (
	// "log"
	"net/http"
	"strconv"
	"time"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

type QuestionCreateOutput struct {
  Question          *model.Question
  Message          string
  ValidationErrors []string
}

func QuestionCreate(c echo.Context) error {

  var question model.Question

  var out QuestionCreateOutput

  if err := c.Bind(&question); err != nil {
    c.Logger().Error(err.Error())
    // リクエストの解釈に失敗した場合は 400 エラーを返却します。
    return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーションチェックを実行します。
  if err := c.Validate(&question); err != nil {

    c.Logger().Error(err.Error())
    out.ValidationErrors = question.ValidationErrors(err)
    // 解釈できたパラメータが許可されていない値の場合は 422 エラーを返却します。
    return c.JSON(http.StatusUnprocessableEntity, out)
  }

  // repository を呼び出して保存処理を実行します。
  res, err := repository.QuestionCreate(&question)
  if err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
    return c.JSON(http.StatusInternalServerError, out)
  }

  // SQL 実行結果から作成されたレコードの ID を取得します。
  id, _ := res.LastInsertId()

  // 構造体に ID をセットします。
  question.ID = int(id)

  // レスポンスの構造体に保存した記事のデータを格納します。
  out.Question = &question

  // 処理成功時はステータスコード 200 でレスポンスを返却します。
  return c.JSON(http.StatusOK, out)
}

func QuestionIndex(c echo.Context) error {

	userId := userIDFromToken(c)

	myUser,err := repository.GetMyUser(userId)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	// リポジトリの処理を呼び出して記事の一覧データを取得します。
	questions, err := repository.QuestionListByCursor(0)

	// エラーが発生した場合
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"記事の一覧データを取得中にエラー発生")
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"user": myUser,
		"Questions": questions,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却します
	return c.JSON(http.StatusOK, data)
}

func QuestionShow(c echo.Context) error {

	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	// パスパラメータから記事 ID を取得します。
	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("id"))

	// 記事データを取得します。
	question, err := repository.QuestionGetByID(id)

	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// ステータスコード 500 でレスポンスを返却します。
		return c.NoContent(http.StatusInternalServerError)
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"user": myUser,
		"Question": question,
	}

	return c.JSON(http.StatusOK, data)
}

func QuestionEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Question Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return c.JSON(http.StatusOK, data)
}

func QuestionNew(c echo.Context) error {
	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	data := map[string]interface{}{
		"user": myUser,
	}

	return c.JSON(http.StatusOK, data)
}