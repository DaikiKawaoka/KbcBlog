package handler

import (
	"net/http"
	"strconv"
	"strings"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)


// ArticleUpdate ...
func UserPasswordEdit(c echo.Context) error {
	// リクエスト送信元のパスを取得
	ref := c.Request().Referer()
	// リクエスト送信元のパスから記事 ID を抽出。
	refID := strings.Split(ref, "/")[4]
	// リクエスト URL のパスパラメータから記事 ID を抽出します。
	reqID := c.Param("id")

	// 編集画面で表示している記事と更新しようとしている記事が異なる場合は、
	// 更新処理をせずに 400 エラーを返却します。
	if reqID != refID {
		return c.JSON(http.StatusBadRequest, "パス違い")
	}
	var article model.Article
	var out ArticleUpdateOutput

	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, out)
	}

	if err := c.Validate(&article); err != nil {
		out.ValidationErrors = article.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out) //422
	}

	// 文字列型の ID を数値型にキャスト
	articleID, _ := strconv.Atoi(reqID)
	article.ID = articleID
	_, err := repository.ArticleUpdate(&article)

	if err != nil {
		out.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, out)
	}
	out.Article = &article
	return c.JSON(http.StatusOK, out)
}