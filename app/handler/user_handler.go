package handler

import (
	// "fmt"
	// "log"
	"net/http"
	"strconv"
	// "time"
	// "github.com/nfnt/resize"
	// "image"
	_"image/gif"
	// "image/jpeg"
	// _"image/png"
	// "os"
	"encoding/base64"
	// "encoding/json"
	// "fmt"
	"io/ioutil"

	"app/model"
	"app/repository"
	"github.com/labstack/echo/v4"
)

// UserCreateOutput ...
type UserCreateOutput struct {
	User             *model.User
	Message          string
	ValidationErrors []string
}

// UserCreate ...
func UserCreate(c echo.Context) error {
	var createuser model.CreateUser
	var user model.User
	var out UserCreateOutput

	if err := c.Bind(&createuser); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, out) //400
	}

	if err := c.Validate(&createuser); err != nil {
		c.Logger().Error(err.Error())
		out.ValidationErrors = createuser.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out) //422
	}

	res, err := repository.UserCreate(&createuser)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, out) //500
	}
	id, _ := res.LastInsertId()
	user.SetupUser(createuser, int(id))
	out.User = &user
	return c.JSON(http.StatusOK, out)
}

// UserShow ...
func UserShow(c echo.Context) error {
	userID := userIDFromToken(c)
	myUser, err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "userが存在しません")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repository.UserGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザデータを取得中にエラー発生") //500
	}

	if err := repository.PostCount(user); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザの投稿数取得中にエラー発生") //500
	}

	articles, err := repository.GetUserArticle(0, user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザの記事一覧データを取得中にエラー発生") //500
	}

	questions, err := repository.GetUserQuestion(0, user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザの質問一覧データを取得中にエラー発生") //500
	}

	answerQuestions, err := repository.GetAnswerQuestionList(0, user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザの回答一覧データを取得中にエラー発生") //500
	}

	// Follow情報を取得
	var follow model.Follow
	if userID != user.ID {
		count, err := repository.CheckFollow(userID, user.ID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "Follw情報取得中にエラー発生") //500
		}
		// ユーザがfollwしているか検証
		if count > 0 {
			follow.IsFollow = true
		}
	}
	// フォロー数取得
	follow.FollowerCount, err = repository.GetFollowerCount(user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "フォロー数取得中にエラー発生") //500
	}
	// フォロワー数取得
	follow.FollowedCount, err = repository.GetFollowedCount(user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "フォロワー数取得中にエラー発生") //500
	}

	// GetArticleTagsユーザ全記事のタグ配列取得
	tags, err := repository.GetArticleTags(user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "ユーザのタグ配列取得中にエラー発生") //500
	}

	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError, "通知数取得中にエアー発生")
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"MyUser":            myUser,
		"User":              user,
		"Articles":          articles,
		"Questions":         questions,
		"AnswerQuestions":   answerQuestions,
		"Follow":            follow,
		"Tags":              tags,
		"NotificationCount": notificationCount,
		// "likes" : likes,
	}

	return c.JSON(http.StatusOK, data)
}

// UserEdit ...
func UserEdit(c echo.Context) error {
	userID := userIDFromToken(c)
	myUser, err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "userが存在しません")
	}
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := repository.UserGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "User情報取得中にエラー発生") //500
	}

	if myUser.ID != user.ID {
		return c.JSON(http.StatusBadRequest, "他人のプロフィールは編集できません。")
	}

	notificationCount, err := repository.GetNotificationCount(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError, "通知数取得中にエアー発生")
	}

	data := map[string]interface{}{
		"MyUser":            myUser,
		"User":              user,
		"NotificationCount": notificationCount,
	}
	return c.JSON(http.StatusOK, data)
}

// UserUpdateOutput ...
type UserUpdateOutput struct {
	User             *model.User
	Message          string
	ValidationErrors []string
}

// UserUpdate ...
func UserUpdate(c echo.Context) error {
	userID := userIDFromToken(c)
	myUser, err := repository.GetMyUser(userID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "userが存在しません")
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if myUser.ID != id {
		return c.JSON(http.StatusBadRequest, "他人の情報は編集できません。") //400
	}

	var user model.User
	var out UserUpdateOutput
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, out) //400
	}
	// 入力値のチェック
	if err := c.Validate(&user); err != nil {
		out.ValidationErrors = user.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out)
	}
	user.ID = id
	if err := repository.UserUpdate(&user); err != nil {
		out.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, out)
	}
	out.User = &user
	return c.JSON(http.StatusOK, out)
}




// UserImgUpdate ...
func UserImgUpdate(c echo.Context) error {
	var imgUser model.UserImgUpdateClass
	imgUser.ID = userIDFromToken(c)

	id, _ := strconv.Atoi(c.Param("id"))
	if imgUser.ID != id {
		return c.JSON(http.StatusBadRequest, "他人の情報は編集できません。") //400
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	/*ファイルをBase64エンコード*/
	binary, _ := ioutil.ReadAll(src)
	imgUser.ImgData64 = base64.StdEncoding.EncodeToString(binary)

	if err := repository.UserImgUpdate(&imgUser); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "変更失敗")
	}

	var user model.User
	user.ImgData64.Valid = true
	user.ImgData64.String = imgUser.ImgData64

	data := map[string]interface{}{
		"ImgData64": user.ImgData64,
	}

	return c.JSON(http.StatusOK, data)
}

// UserImgDelete ...
func UserImgDelete(c echo.Context) error {
	var imgUser model.UserImgUpdateClass
	imgUser.ID = userIDFromToken(c)

	id, _ := strconv.Atoi(c.Param("id"))
	if imgUser.ID != id {
		return c.JSON(http.StatusBadRequest, "他人の情報は編集できません。") //400
	}

	if err := repository.UserImgDelete(&imgUser); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "削除失敗")
	}

	var user model.User
	user.ImgData64.Valid = false
	user.ImgData64.String = ""

	data := map[string]interface{}{
		"ImgData64": user.ImgData64,
	}

	return c.JSON(http.StatusOK, data)
}
