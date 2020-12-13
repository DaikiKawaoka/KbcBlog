package handler

import (
	"net/http"
	"strconv"
	"strings"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// PasswordEditOutput パスワード変更リターン型
type PasswordEditOutput struct {
  Message          string
  ValidationErrors []string
}

// UserPasswordEdit パスワード編集処理
func UserPasswordEdit(c echo.Context) error {
	ref := c.Request().Referer()
	refID := strings.Split(ref, "/")[4]
	reqID := c.Param("id")

	if reqID != refID {
		return c.JSON(http.StatusBadRequest, "パス違い") //400
	}
	var password model.Password
	var out PasswordEditOutput

	if err := c.Bind(&password); err != nil {
		return c.JSON(http.StatusBadRequest, out)
	}

	if err := c.Validate(&password); err != nil {
		out.ValidationErrors = password.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out) //422
	}

	password.ID , _ = strconv.Atoi(reqID)

	dbpasshash,err := repository.GetPassword(password.ID);
	if err != nil {
		out.Message = "パスワード取得中にエラー"
		return c.JSON(http.StatusInternalServerError, out)
	}

	if err := passwordVerify(dbpasshash, password.CurrentPassword); err != nil {
		out.ValidationErrors = append(out.ValidationErrors, "現在のパスワードが違います。")
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	if err := password.NewPasswordHash(); err != nil {
		out.Message = "パスワードhash中にエラー"
		return c.JSON(http.StatusInternalServerError, out)
	}

	if err := repository.PasswordUpdate(&password); err != nil{
    out.Message = "update中にエラー"
    return c.JSON(http.StatusInternalServerError, out)
  }

	return c.JSON(http.StatusOK, out)
}