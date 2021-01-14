package handler

import (
	"net/http"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)


// GetArticleRanking ...
func GetArticleRanking(c echo.Context) error {

	rankingType := c.QueryParam("rankingType")
	period := c.QueryParam("period")//絞り込みタグ

	var Ranking []*model.RankingUser
	var err error
	if rankingType == "" {
		Ranking,err = repository.KBCArticleRankingTop10(rankingType,period)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"likeRankingを取得中にエラー発生")
		}
	}else{
		Ranking,err = repository.KBCArticleRankingTop10(rankingType,period)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"postRankingを取得中にエラー発生")
		}
	}

	data := map[string]interface{}{
		"Ranking": Ranking,
	}
	return c.JSON(http.StatusOK, data)
}