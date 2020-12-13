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
  // SQL 実行結果から作成されたレコードの ID を取得します。
  id, _ := res.LastInsertId()

  // 構造体に ID をセットします。
	comment.ID = int(id)

  // レスポンスの構造体に保存した記事のデータを格納します。
  out.Comment = &comment
  return c.JSON(http.StatusOK, out)
}

// ArticleCommentIndex 記事コメントリスト取得処理
func ArticleCommentIndex(c echo.Context) error {

	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("id"))
	comments, err := repository.ArticleCommentListByCursor(id)

	// エラーが発生した場合
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"コメント一覧を取得中にエラー発生")
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Comments": comments,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却します
	return c.JSON(http.StatusOK, data)
}

// ArticleCommentDelete 記事コメント削除処理
func ArticleCommentDelete(c echo.Context) error {
	// パスパラメータから記事 ID を取得します。
	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("cid"))

	// repository の記事削除処理を呼び出します。
	if err := repository.ArticleCommentDelete(id); err != nil {
		// サーバーのログにエラー内容を出力します。
		c.Logger().Error(err.Error())

		// サーバーサイドでエラーが発生した場合は 500 エラーを返却します。
		return c.JSON(http.StatusInternalServerError, "コメント削除中にエラー発生")
	}

	// 成功時はステータスコード 200 を返却します。
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
    // リクエストの解釈に失敗した場合は 400 エラーを返却します。
    return c.JSON(http.StatusBadRequest, out)
	}
	// バリデーションチェックを実行します。
  if err := c.Validate(&comment); err != nil {
    c.Logger().Error(err.Error())
    out.ValidationErrors = comment.ValidationErrors(err)
    // 解釈できたパラメータが許可されていない値の場合は 422 エラーを返却します。
    return c.JSON(http.StatusUnprocessableEntity, out)
  }
  // repository を呼び出して保存処理を実行します。
  res, err := repository.QuestionCommentCreate(&comment)
  if err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())
    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
    return c.JSON(http.StatusInternalServerError, out)
	}
  // SQL 実行結果から作成されたレコードの ID を取得します。
  id, _ := res.LastInsertId()

  // 構造体に ID をセットします。
  comment.ID = int(id)

  // レスポンスの構造体に保存した記事のデータを格納します。
  out.Comment = &comment

  // 処理成功時はステータスコード 200 でレスポンスを返却します。
  return c.JSON(http.StatusOK, out)
}

// QuestionCommentIndex 質問コメントリスト取得処理
func QuestionCommentIndex(c echo.Context) error {

	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("id"))
	comments, err := repository.QuestionCommentListByCursor(id)

	// エラーが発生した場合
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"コメント一覧を取得中にエラー発生")
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Comments": comments,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却します
	return c.JSON(http.StatusOK, data)
}

// QuestionCommentDelete 質問コメント削除処理
func QuestionCommentDelete(c echo.Context) error {
	// パスパラメータから記事 ID を取得します。
	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("cid"))

	// repository の記事削除処理を呼び出します。
	if err := repository.QuestionCommentDelete(id); err != nil {
		// サーバーのログにエラー内容を出力します。
		c.Logger().Error(err.Error())

		// サーバーサイドでエラーが発生した場合は 500 エラーを返却します。
		return c.JSON(http.StatusInternalServerError, "コメント削除中にエラー発生")
	}

	// 成功時はステータスコード 200 を返却します。
	return c.JSON(http.StatusOK, "コメントを削除しました。")
}