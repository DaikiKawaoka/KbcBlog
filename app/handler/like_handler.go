package handler

import (
	"net/http"
	"strconv"

	"app/repository"
	// "app/model"
	"github.com/labstack/echo/v4"
)

// like

func ArticleLike(c echo.Context) error {

	userId := userIDFromToken(c)
	// articleId取得
	articleId, _ := strconv.Atoi(c.Param("id"))

	like, err := repository.GetArticleLike(userId,articleId)
	if err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())
    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
    return c.JSON(http.StatusInternalServerError,"likeデータの取得中にエラー発生")
	}
	if like != nil {
		// Like削除
		err := repository.DeleteArticleLike(like.Userid,like.Articleid)
		if err != nil {
			c.Logger().Error(err.Error())
			// 500 エラー
			return c.JSON(http.StatusInternalServerError, "likeデータ削除中にエラー発生")
		}
	}else{
		// Like作成
		err := repository.CreateArticleLike(like)
		if err != nil {
			c.Logger().Error(err.Error())
			// 500 エラー
			return c.JSON(http.StatusInternalServerError, "likeデータ作成中にエラー発生")
		}
	}

  // 処理成功時はステータスコード 200 でレスポンスを返却します。
  return c.JSON(http.StatusOK, "LIKE change 成功")
}