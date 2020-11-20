package handler

import (
	"net/http"
	"strconv"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

func Follow(c echo.Context) error {
	followerId := userIDFromToken(c)             // フォローした人
	followedId, _ := strconv.Atoi(c.Param("id")) //フォローされた人
	count, err := repository.CheckFollow(followerId,followedId)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"followデータの取得中にエラー発生") //500
	}
	if count > 0 {
		// UnFollow
		err := repository.UnFollow(followerId, followedId)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "UnFollow中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "UnFollow 成功")
	}else{
		var follow model.Following
		follow.FollowerId = followerId
		follow.FollowedId = followedId
		// Follow
		err := repository.Follow(&follow)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "Follow中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "Follow 成功")
	}
}