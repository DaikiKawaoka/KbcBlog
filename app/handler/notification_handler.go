package handler

import (
	"net/http"
	"strconv"
	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// Notifications 通知ページ
func Notifications(c echo.Context) error {
	userID := userIDFromToken(c)
	notifications,err := repository.GetNotification(userID,0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,"通知取得中にエラー") //500
	}

	err = repository.NotificationChecked(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,"通知カウントリセット中にエラー") //500
	}

	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"通知数取得中にエアー発生")
	}

	data := map[string]interface{}{
				"Notifications": notifications,
				"NotificationCount": notificationCount,
			}
	return c.JSON(http.StatusOK, data)
}


// NotificationCreate 新規通知処理
func NotificationCreate(notification model.CreateNotification) error {
  err := repository.NotificationCreate(&notification)
  if err != nil {
    return err
  }
  return nil
}

// NotificationDelete 通知削除処理
func NotificationDelete(c echo.Context) error {
	userID := userIDFromToken(c)
	if err := repository.NotificationDelete(userID); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "通知削除中にエラー発生")
	}
	return c.JSON(http.StatusOK, "notification delete")
}


// NotificationOrder 記事一覧のスクロール時に次の4記事を返す
func NotificationOrder(c echo.Context) error {
	userID := userIDFromToken(c)
	notificationCursor, _ := strconv.Atoi(c.QueryParam("notificationCursor"))

	notifications, err := repository.GetNotification(userID,notificationCursor)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"通知取得中にエラー")
	}
	notificationCursor += 4

	data := map[string]interface{}{
		"Notifications": notifications,
		"NotificationCursor":   notificationCursor,
	}
	return c.JSON(http.StatusOK, data)
}