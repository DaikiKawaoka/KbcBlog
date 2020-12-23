package handler

import (
	"net/http"
	"strconv"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// Follow フォロー、アンフォロー処理
func Follow(c echo.Context) error {
	followerID := userIDFromToken(c)             // フォローした人
	followedID, _ := strconv.Atoi(c.Param("id")) //フォローされた人
	count, err := repository.CheckFollow(followerID,followedID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"followデータの取得中にエラー発生") //500
	}

	if count > 0 {
		// UnFollow
		err := repository.UnFollow(followerID, followedID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "UnFollow中にエラー発生") //500
		}
		return c.JSON(http.StatusOK, "UnFollow 成功")
	}
		var follow model.Following
		follow.FollowerID = followerID
		follow.FollowedID = followedID
		// Follow
		err2 := repository.Follow(&follow)
		if err2 != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "Follow中にエラー発生") //500
		}


	// 通知作成
	var notification model.CreateNotification
	if err := c.Bind(&notification); err != nil {
			return err
	}
	if notification.Visitedid != notification.Visiterid{
		if err := repository.NotificationCreate(&notification); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "通知作成失敗")
		}
	}

	return c.JSON(http.StatusOK, "Follow 成功")
}