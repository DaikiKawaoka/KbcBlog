package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

func TestArticleNew(t *testing.T) {
		e := echo.New()
    req := httptest.NewRequest("GET", "/guestLogin", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/Users")
    e.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "Hello, Echo World!!", rec.Body.String())
}