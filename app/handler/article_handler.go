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
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	data := map[string]interface{}{
		"user": myUser,
	}

	return c.JSON(http.StatusOK, data)
}

type ArticleCreateOutput struct {
  Article          *model.Article
  Message          string
  ValidationErrors []string
}

func ArticleCreate(c echo.Context) error {
  var article model.Article
  var out ArticleCreateOutput
  if err := c.Bind(&article); err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusBadRequest, out)
	}

  if err := c.Validate(&article); err != nil {
    c.Logger().Error(err.Error())
    out.ValidationErrors = article.ValidationErrors(err)
    return c.JSON(http.StatusUnprocessableEntity, out) //422
  }
  res, err := repository.ArticleCreate(&article)
  if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError, out)
  }
  id, _ := res.LastInsertId()
  article.ID = int(id)
  out.Article = &article
  return c.JSON(http.StatusOK, out)
}

func ArticleIndex(c echo.Context) error {

	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	articles, err := repository.ArticleListByCursor(0,"new")

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"記事の一覧データを取得中にエラー発生")
	}

	likeRanking,err := repository.KBCRankingTop10("like")
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"likeRankingを取得中にエラー発生")
	}

	postRanking,err := repository.KBCRankingTop10("post")
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"likeRankingを取得中にエラー発生")
	}

	data := map[string]interface{}{
		"user": myUser,
		"Articles": articles,
		"LikeRanking": likeRanking,
		"PostRanking": postRanking,
	}
	return c.JSON(http.StatusOK, data)
}

func ArticleShow(c echo.Context) error {
	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := repository.ArticleGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
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

	var commentLike model.Like
	for i, comment := range comments {
		count, err := repository.GetArticleCommentLike(userId,comment.ID);
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
		commentLike.LikeCount,err = repository.GetArticleCommentLikeCount(comment.ID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"likeデータのカウント数取得中にエラー発生") //500
		}
		comments[i].Like = commentLike
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
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := repository.ArticleGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"記事情報取得中にエラー発生")
	}

	data := map[string]interface{}{
		"user": myUser,
		"Article": article,
	}
	return c.JSON(http.StatusOK, data)
}

type ArticleUpdateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

// ArticleUpdate ...
func ArticleUpdate(c echo.Context) error {
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

func ArticleDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.ArticleDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "記事削除中にエラー発生")
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted.", id))
}

func ArticleIndexOrder(c echo.Context) error {

	order := c.QueryParam("order")
	articles, err := repository.ArticleListByCursor(0,order)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"記事の一覧データを取得中にエラー発生")
	}

	data := map[string]interface{}{
		"Articles": articles,
	}
	return c.JSON(http.StatusOK, data)
}