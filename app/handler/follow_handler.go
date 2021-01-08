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

	if count >= 1 {
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


// Following ...
func Following(c echo.Context) error {
	userID := userIDFromToken(c)

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repository.UserGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザデータを取得中にエラー発生") //500
	}

	// フォローリスト取得
	following, err := repository.GetFollowerInfoList(user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "フォローリスト取得中にエラー発生") //500
	}


	for _, f := range following {
		count, err := repository.CheckFollow(userID, f.ID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "followcheck中にエラー発生") //500
		}
		// ユーザがいいねしているか検証
		if count > 0 {
			f.IsFollowing = true
		} else {
			f.IsFollowing = false
		}
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Following":           following,
	}

	return c.JSON(http.StatusOK, data)
}

// Followers ...
func Followers(c echo.Context) error {
	userID := userIDFromToken(c)

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repository.UserGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザデータを取得中にエラー発生") //500
	}

	// フォロワーリスト
	followers, err := repository.GetFollowedInfoList(user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "フォロワーリスト取得中にエラー発生") //500
	}


	for _, f := range followers {
		count, err := repository.CheckFollow(userID, f.ID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "followcheck中にエラー発生") //500
		}
		// ユーザがいいねしているか検証
		if count > 0 {
			f.IsFollowing = true
		} else {
			f.IsFollowing = false
		}
	}
	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Followers":           followers,
	}

	return c.JSON(http.StatusOK, data)
}