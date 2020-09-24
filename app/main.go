package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/", articleIndex)
	e.GET("/Articles/new", articleNew)
	e.GET("/Articles/:id", articleShow)
	e.GET("/Articles/:id/edit", articleEdit)
	// Webサーバーをポート番号 8082 で起動する
	e.Logger.Fatal(e.Start(":8082"))
	// e.Startの中はdocker-composeのgoコンテナで設定したportsを指定してください。
}

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article Index",
		"Now":     time.Now(),
	}
	return c.JSON(http.StatusOK, data)
}
func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return c.JSON(http.StatusOK, data)

}

func articleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return c.JSON(http.StatusOK, data)

}

func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return c.JSON(http.StatusOK, data)

}

func createMux() *echo.Echo {
	// アプリケーションインスタンスを生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	//アプリケーションのどこかで予期せずにpanicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーする
	e.Use(middleware.Recover())
	//アクセスログのようなリクエスト単位のログを出力してくれます。
	e.Use(middleware.Logger())
	//gzip圧縮スキームを使用してHTTP応答を圧縮します。
	e.Use(middleware.Gzip())

	// アプリケーションインスタンスを返却
	return e
}
