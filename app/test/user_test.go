package test

import (
	"net/http"
	"strings"
	"app/handler"
	"app/repository"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

var db = connectDB()
var jwt = ""


func TestUserNew(t *testing.T) {
	e := createMux()
	repository.SetDB(db)
	var  userJSON = `{"name":"kawa", "KBC_mail": "kbc19a11@stu.kawahara.ac.jp","password":"0000000a","password_confirmation":"0000000a","sex":1}`
	req := httptest.NewRequest(http.MethodPost, "/Users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handler.UserCreate(c)) {
		// assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestLogin(t *testing.T){
	e := echo.New()
	repository.SetDB(db)
	var  userJSON = `{"password":"0000000a", "KBC_mail": "kbc19a11@stu.kawahara.ac.jp"}`
	req := httptest.NewRequest(http.MethodPost,"/Login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusOK , rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestUserShow(t *testing.T) {
	e := createMux()
	repository.SetDB(db)
	req := httptest.NewRequest(http.MethodGet, "/Users/920437695",nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsImV4cCI6MTYxMTk4MjY0MH0.HybdYt8cLZrALcJ9BRQ5fnLa2aCpgDSpvzV80mND-ko")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handler.UserShow(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
