package handler

import (
	"net/http"
	"strconv"
  "path/filepath"
	"app/model"
	"app/repository"
	"github.com/labstack/echo/v4"
)

// UserImgUpdate ...
func UserImgUpdate(c echo.Context) error {
	var (
		awsS3 *model.AwsS3
		filetype string
		imgUser model.UserImgUpdateClass
	)
	imgUser.ID = userIDFromToken(c)
	ids := c.Param("id") //string
	id, _ := strconv.Atoi(ids) //int
	if imgUser.ID != id {
		return c.JSON(http.StatusBadRequest, "他人の情報は編集できません。") //400
	}

	file,err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, "") //400
	}
	defer src.Close()

	filetype = filepath.Ext(file.Filename)
	awsS3 = model.NewAwsS3()

	// multipart.File と os.File は同じように扱える
	imgUser.ImgPath, err = awsS3.UploadTest(src,ids,filetype)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, imgUser.ImgPath) //400
	}

	if err := repository.UserImgUpdate(&imgUser); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "変更失敗")
	}

	data := map[string]interface{}{
		"ImgPath": imgUser.ImgPath,
	}
	return c.JSON(http.StatusOK, data)
}

// UserImgDelete ...
func UserImgDelete(c echo.Context) error {
	userID := userIDFromToken(c)
	id, _ := strconv.Atoi(c.Param("id"))
	if userID != id {
		return c.JSON(http.StatusBadRequest, "他人の情報は編集できません。") //400
	}
	myUser,err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	if err := repository.UserImgDelete(myUser); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "削除失敗")
	}
	data := map[string]interface{}{
		"ImgPath": myUser.ImgPath,
	}

	return c.JSON(http.StatusOK, data)
}
