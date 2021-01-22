package handler

import (
	"net/http"
	"app/repository"
	"github.com/labstack/echo/v4"
)

// About ...
func About(c echo.Context) error {
	userID := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}
	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"通知数取得中にエアー発生")
	}
	data := map[string]interface{}{
		"user": myUser,
		"NotificationCount": notificationCount,
	}
	return c.JSON(http.StatusOK, data)
}