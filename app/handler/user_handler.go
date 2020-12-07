package handler

import (
	// "log"
	"net/http"
	"strconv"
	// "time"

	"app/repository"
	"app/model"
	"github.com/labstack/echo/v4"
)

// UserCreateOutput ...
type UserCreateOutput struct {
  User          *model.User
  Message          string
  ValidationErrors []string
}


func UserCreate(c echo.Context) error {
  // 送信されてくるフォームの内容を格納する構造体を宣言します。
	var createuser model.CreateUser
	var user model.User

  // レスポンスとして返却する構造体を宣言します。
  var out UserCreateOutput

  // フォームの内容を構造体に埋め込みます。
  if err := c.Bind(&createuser); err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // リクエストの解釈に失敗した場合は 400 エラーを返却します。
    return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーションチェックを実行します。
  if err := c.Validate(&createuser); err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())
    // エラー内容を検査してカスタムエラーメッセージを取得します。
    out.ValidationErrors = createuser.ValidationErrors(err)
    // 解釈できたパラメータが許可されていない値の場合は 422 エラーを返却します。
    return c.JSON(http.StatusUnprocessableEntity, out)
  }

  // repository を呼び出して保存処理を実行します。
  res, err := repository.UserCreate(&createuser)
  if err != nil {
    // エラーの内容をサーバーのログに出力します。
    c.Logger().Error(err.Error())

    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
    return c.JSON(http.StatusInternalServerError, out)
  }

  // SQL 実行結果から作成されたレコードの ID を取得します。
  id, _ := res.LastInsertId()

  // 構造体に ID をセットします。
  // user.ID = int(id)
  // Userをセットアップする
  user.SetupUser(createuser,int(id))

  // レスポンスの構造体に保存した記事のデータを格納します。
  out.User = &user

  // 処理成功時はステータスコード 200 でレスポンスを返却します。
  return c.JSON(http.StatusOK,out)
}


func UserShow(c echo.Context) error {
	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	// 記事データを取得します。
	user, err := repository.UserGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
		// ステータスコード 500 でレスポンスを返却します。
		return c.JSON(http.StatusInternalServerError,"ユーザデータを取得中にエラー発生")
  }

  if err := repository.PostCount(user); err != nil{
    c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"ユーザの投稿数取得中にエラー発生")
  }

  // リポジトリの処理を呼び出して記事の一覧データを取得します。
	articles, err := repository.GetUserArticle(0,user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.JSON(http.StatusInternalServerError,"ユーザの記事一覧データを取得中にエラー発生")
	}

	questions, err := repository.GetUserQuestion(0,user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"ユーザの質問一覧データを取得中にエラー発生")
	}

	answerQuestions, err := repository.GetAnswerQuestionList(0,user.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"ユーザの回答一覧データを取得中にエラー発生")
	}
	// Follow情報を取得
	var follow model.Follow
	if userId != user.ID{
		count, err := repository.CheckFollow(userId,user.ID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"Follw情報取得中にエラー発生") //500
		}
		// ユーザがfollwしているか検証
		if count > 0{
			follow.IsFollow = true
		}
	}
	// フォロー数取得
	follow.FollowerCount,err = repository.GetFollowerCount(user.ID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"フォロー数取得中にエラー発生") //500
	}
	// フォロワー数取得
	follow.FollowedCount,err = repository.GetFollowedCount(user.ID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"フォロワー数取得中にエラー発生") //500
	}

	// フォローリスト取得
	followers , err := repository.GetFollowerInfoList(user.ID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"フォロワーリスト取得中にエラー発生") //500
	}

  // フォロワーリスト
	followeds , err := repository.GetFollowedInfoList(user.ID)
	if err != nil {
    c.Logger().Error(err.Error())
    return c.JSON(http.StatusInternalServerError,"フォロワーリスト取得中にエラー発生") //500
	}

	for _,f := range followers {
		count, err := repository.CheckFollow(userId,f.ID);
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"followcheck中にエラー発生") //500
		}
		// ユーザがいいねしているか検証
		if count > 0{
			f.IsFollowing = true
		}else{
			f.IsFollowing = false
		}
	}

	for _,f := range followeds {
		count, err := repository.CheckFollow(userId,f.ID);
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError,"followcheck中にエラー発生") //500
		}
		// ユーザがいいねしているか検証
		if count > 0{
			f.IsFollowing = true
		}else{
			f.IsFollowing = false
		}
	}


	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
    "MyUser": myUser,
    "User":user,
		"Articles": articles,
		"Questions": questions,
		"AnswerQuestions" : answerQuestions,
		"Follow": follow,
		"Follows":followers,
		"Followers":followeds,
		// "likes" : likes,
	}

	return c.JSON(http.StatusOK, data)
}

func UserEdit(c echo.Context) error {
	userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
  }

	id, _ := strconv.Atoi(c.Param("id"))

	// 編集フォームの初期値として表示するために記事データを取得します。
	user, err := repository.UserGetByID(id)
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())
		// ステータスコード 500 でレスポンスを返却します。
		return c.JSON(http.StatusInternalServerError,"User情報取得中にエラー発生")
  }


  if myUser.ID != user.ID{
    return c.JSON(http.StatusBadRequest,"他人のプロフィールは編集できません。")
  }

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"MyUser": myUser,
		"User": user,
	}
	return c.JSON(http.StatusOK, data)
}


type UserUpdateOutput struct {
	User         *model.User
	Message          string
	ValidationErrors []string
}

func UserUpdate(c echo.Context) error {

  userId := userIDFromToken(c)
	myUser,err := repository.GetMyUser(userId)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError,"userが存在しません")
  }

	id, _ := strconv.Atoi(c.Param("id"))

	// 更新処理をせずに 400 エラーを返却します。
	if myUser.ID != id {
		return c.JSON(http.StatusBadRequest, "他人の情報は編集できません。")
	}

	// フォームで送信される記事データを格納する構造体を宣言します。
	var user model.User

	// レスポンスするデータの構造体を宣言します。
	var out UserUpdateOutput

	if err := c.Bind(&user); err != nil {
		// リクエストのパラメータの解釈に失敗した場合は 400 エラーを返却します。
		return c.JSON(http.StatusBadRequest, out)
  }

  // user.SetupUserCommnet()

	// 入力値のチェック（バリデーションチェック）を行います。
	if err := c.Validate(&user); err != nil {
		out.ValidationErrors = user.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	// フォームデータを格納した構造体に ID をセットします。
  user.ID = id

  if err := repository.UserUpdate(&user); err != nil{
    out.Message = err.Error()
    return c.JSON(http.StatusInternalServerError, out)
  }

	// レスポンスの構造体に記事データをセットします。
	out.User = &user

	// 処理成功時はステータスコード 200 でレスポンスを返却します。
	return c.JSON(http.StatusOK, out)
}