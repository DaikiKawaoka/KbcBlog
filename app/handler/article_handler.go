package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)


func ArticleNew(c echo.Context) error {
	userId := userIDFromToken(c)

	myUser,err := repository.GetMyUser(userId)

	// エラーが発生した場合
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	data := map[string]interface{}{
		"user": myUser,
	}

	return c.JSON(http.StatusOK, data)
}

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
func ArticleIndex(c echo.Context) error {

	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)

	// エラーが発生した場合
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	// リポジトリの処理を呼び出して記事の一覧データを取得します。
	articles, err := repository.ArticleListByCursor(0)

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
		"Articles": articles,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却します
	return c.JSON(http.StatusOK, data)
}

// ArticleShow ...
func ArticleShow(c echo.Context) error {
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
	article, err := repository.ArticleGetByID(id)

	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())
		// ステータスコード 500 でレスポンスを返却します。
		return c.JSON(http.StatusInternalServerError,"記事データを取得中にエラー発生")
	}

	// 記事データのいいねを取得する
	var like model.Like
	count, err := repository.GetArticleLike(userId,id)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生") //500
	}
	// ユーザがいいねしているか検証
	if count > 0{
		like.IsLike = true
	}

	like.LikeCount,err = repository.GetArticleLikeCount(id)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"likeデータのカウント数取得中にエラー発生") //500
	}

	// 記事データへのコメント一覧データを取得します。
	comments, err := repository.ArticleCommentListByCursor(article.ID)

	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"記事のコメント一覧データを取得中にエラー発生")
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"user": myUser,
		"Article": article,
		"Like": like,
		"Comments": comments,
	}

	return c.JSON(http.StatusOK, data)
}

// ArticleEdit ...
func ArticleEdit(c echo.Context) error {
	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}
	// パスパラメータから記事 ID を取得します。
	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("id"))

	// 編集フォームの初期値として表示するために記事データを取得します。
	article, err := repository.ArticleGetByID(id)
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// ステータスコード 500 でレスポンスを返却します。
		return c.JSON(http.StatusInternalServerError,"記事情報取得中にエラー発生")
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"user": myUser,
		"Article": article,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却します。
	return c.JSON(http.StatusOK, data)
}

// ArticleUpdateOutput ...
type ArticleUpdateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

// ArticleUpdate ...
func ArticleUpdate(c echo.Context) error {
	// リクエスト送信元のパスを取得します。
	ref := c.Request().Referer()

	// リクエスト送信元のパスから記事 ID を抽出します。
	refID := strings.Split(ref, "/")[4]

	// リクエスト URL のパスパラメータから記事 ID を抽出します。
	reqID := c.Param("id")

	// 編集画面で表示している記事と更新しようとしている記事が異なる場合は、
	// 更新処理をせずに 400 エラーを返却します。
	if reqID != refID {
		return c.JSON(http.StatusBadRequest, "パス違い")
	}

	// フォームで送信される記事データを格納する構造体を宣言します。
	var article model.Article

	// レスポンスするデータの構造体を宣言します。
	var out ArticleUpdateOutput

	// フォームで送信されたデータを変数に格納します。
	if err := c.Bind(&article); err != nil {
		// リクエストのパラメータの解釈に失敗した場合は 400 エラーを返却します。
		return c.JSON(http.StatusBadRequest, out)
	}

	// 入力値のチェック（バリデーションチェック）を行います。
	if err := c.Validate(&article); err != nil {
		// エラー内容をレスポンスのフィールドに格納します。
		out.ValidationErrors = article.ValidationErrors(err)

		// 解釈できたパラメータが不正な値の場合は 422 エラーを返却します。
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	// 文字列型の ID を数値型にキャストします。
	articleID, _ := strconv.Atoi(reqID)

	// フォームデータを格納した構造体に ID をセットします。
	article.ID = articleID

	// 記事を更新する処理を呼び出します。
	_, err := repository.ArticleUpdate(&article)

	if err != nil {
		// レスポンスの構造体にエラー内容をセットします。
		out.Message = err.Error()

		// リクエスト自体は正しいにも関わらずサーバー側で処理が失敗した場合は 500 エラーを返却します。
		return c.JSON(http.StatusInternalServerError, out)
	}

	// レスポンスの構造体に記事データをセットします。
	out.Article = &article

	// 処理成功時はステータスコード 200 でレスポンスを返却します。
	return c.JSON(http.StatusOK, out)
}

// ArticleDelete ...
func ArticleDelete(c echo.Context) error {
	// パスパラメータから記事 ID を取得します。
	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("id"))

	// repository の記事削除処理を呼び出します。
	if err := repository.ArticleDelete(id); err != nil {
		// サーバーのログにエラー内容を出力します。
		c.Logger().Error(err.Error())

		// サーバーサイドでエラーが発生した場合は 500 エラーを返却します。
		return c.JSON(http.StatusInternalServerError, "記事削除中にエラー発生")
	}

	// 成功時はステータスコード 200 を返却します。
	return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted.", id))
}