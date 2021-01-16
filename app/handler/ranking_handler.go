package handler

import (
	"net/http"
	"strconv"
	"app/repository"
	"github.com/labstack/echo/v4"
)


// GetArticleRanking ...
func GetArticleRanking(c echo.Context) error {

	rankingType,_ := strconv.Atoi(c.QueryParam("rankingType"))
	period := c.QueryParam("period")//絞り込みタグ

	Ranking,err := repository.KBCArticleRankingTop10(rankingType,period)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"Rankingを取得中にエラー発生")
	}
	data := map[string]interface{}{
		"Ranking": Ranking,
	}
	return c.JSON(http.StatusOK, data)
}

// GetQuestionRanking ...
func GetQuestionRanking(c echo.Context) error {

	rankingType,_ := strconv.Atoi(c.QueryParam("rankingType"))
	period := c.QueryParam("period")//絞り込みタグ
	Ranking,err := repository.KBCQuestionRankingTop10(rankingType,period)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"likeRankingを取得中にエラー発生")
	}
	data := map[string]interface{}{
		"Ranking": Ranking,
	}
	return c.JSON(http.StatusOK, data)
}