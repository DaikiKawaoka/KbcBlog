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

// ArticleCreateOutput ...
type ArticleCreateOutput struct {
  Article          *model.Article
  Message          string
  ValidationErrors []string
}

// ArticleCreate ...
func ArticleCreate(c echo.Context) error {
  // 送信されてくるフォームの内容を格納する構造体を宣言します。
  var article model.Article

  // レスポンスとして返却する構造体を宣言します。
  var out ArticleCreateOutput

  // フォームの内容を構造体に埋め込みます。
  if err := c.Bind(&article); err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // リクエストの解釈に失敗した場合は 400 エラーを返却します。
    return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーションチェックを実行します。
  if err := c.Validate(&article); err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // エラー内容を検査してカスタムエラーメッセージを取得します。
    out.ValidationErrors = article.ValidationErrors(err)

    // 解釈できたパラメータが許可されていない値の場合は 422 エラーを返却します。
    return c.JSON(http.StatusUnprocessableEntity, out)
  }

  // repository を呼び出して保存処理を実行します。
  res, err := repository.ArticleCreate(&article)
  if err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
    return c.JSON(http.StatusInternalServerError, out)
  }

  // SQL 実行結果から作成されたレコードの ID を取得します。
  id, _ := res.LastInsertId()

  // 構造体に ID をセットします。
  article.ID = int(id)

  // レスポンスの構造体に保存した記事のデータを格納します。
  out.Article = &article

  // 処理成功時はステータスコード 200 でレスポンスを返却します。
  return c.JSON(http.StatusOK, out)
}

// ArticleIndex ...
// func ArticleIndex(c echo.Context) error {
// 	// 記事データの一覧を取得する
// 	articles, err := repository.ArticleList()
// 	if err != nil {
// 		log.Println(err.Error())
// 		return c.NoContent(http.StatusInternalServerError)
// 	}
// 	data := map[string]interface{}{
// 		"Articles": articles, // 記事データをテンプレートエンジンに渡す
// 	}
// 	return c.JSON(http.StatusOK, data)
// }

// ArticleIndex ...
func ArticleIndex(c echo.Context) error {
	// リポジトリの処理を呼び出して記事の一覧データを取得します。
	articles, err := repository.ArticleListByCursor(0)

	// エラーが発生した場合
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.NoContent(http.StatusInternalServerError)
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Articles": articles,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却します
	return c.JSON(http.StatusOK, data)
}

// ArticleNew ...
func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return c.JSON(http.StatusOK, data)
}

// ArticleShow ...
func ArticleShow(c echo.Context) error {
	// パスパラメータから記事 ID を取得します。
	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("id"))

	// 記事データを取得します。
	article, err := repository.ArticleGetByID(id)

	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// ステータスコード 500 でレスポンスを返却します。
		return c.NoContent(http.StatusInternalServerError)
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Article": article,
	}

	return c.JSON(http.StatusOK, data)
}

// ArticleEdit ...
func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return c.JSON(http.StatusOK, data)
}