package handler

import (
	// "log"
	"net/http"
	"strconv"
	// "time"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// UserCreateOutput ...
type UserCreateOutput struct {
  User          *model.User
  Message          string
  ValidationErrors []string
}


func UserCreate(c echo.Context) error {
  // 送信されてくるフォームの内容を格納する構造体を宣言します。
	var createuser model.CreateUser
	var user model.User

  // レスポンスとして返却する構造体を宣言します。
  var out UserCreateOutput

  // フォームの内容を構造体に埋め込みます。
  if err := c.Bind(&createuser); err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // リクエストの解釈に失敗した場合は 400 エラーを返却します。
    return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーションチェックを実行します。
  if err := c.Validate(&createuser); err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // エラー内容を検査してカスタムエラーメッセージを取得します。
    out.ValidationErrors = createuser.ValidationErrors(err)

    // 解釈できたパラメータが許可されていない値の場合は 422 エラーを返却します。
    return c.JSON(http.StatusUnprocessableEntity, out)
  }

  // repository を呼び出して保存処理を実行します。
  res, err := repository.UserCreate(&createuser)
  if err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
    return c.JSON(http.StatusInternalServerError, out)
  }

  // SQL 実行結果から作成されたレコードの ID を取得します。
  id, _ := res.LastInsertId()

  // 構造体に ID をセットします。
  // user.ID = int(id)
  // Userをセットアップする
  user.SetupUser(createuser,int(id))

  // レスポンスの構造体に保存した記事のデータを格納します。
  out.User = &user

  // 処理成功時はステータスコード 200 でレスポンスを返却します。
  return c.JSON(http.StatusOK,out)
}


func UserShow(c echo.Context) error {
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
	user, err := repository.UserGetByID(id)

	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())
		// ステータスコード 500 でレスポンスを返却します。
		return c.JSON(http.StatusInternalServerError,"ユーザデータを取得中にエラー発生")
	}

  // リポジトリの処理を呼び出して記事の一覧データを取得します。
	articles, err := repository.GetUserArticle(0,user.ID)

	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"ユーザの記事一覧データを取得中にエラー発生")
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
    "MyUser": myUser,
    "User":user,
		"Articles": articles,
	}

	return c.JSON(http.StatusOK, data)
}