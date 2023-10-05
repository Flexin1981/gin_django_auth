package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// Ensure that the "sessionid" cookie is present in the incoming request
// the middleware should return an unauthorized response if not.
func TestMiddlewareReturnsUnauthorisedIfSessionCookieNotPresent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))

	LoginRequired(c)

	if w.Code != 401 {
		t.Error(w.Code)
	}
}

// Ensure that the "sessionid" cookie is not blank in the incoming request
// the middleware should return an unauthorized response if not.
func TestMiddlewareReturnsUnauthorisedIfSessionCookieBlank(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
	c.Request.AddCookie(
		&http.Cookie{
			Name:     "sessionid",
			Value:    "",
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})

	LoginRequired(c)

	if w.Code != 401 {
		t.Error(w.Code)
	}
}

// Ensure that the "sessionid" cookie is not blank in the incoming request
// the middleware should return an unauthorized response if not.
func TestMiddlewareReturnsUnauthorisedIfSessionCookieNotInDataBase(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_ = os.Setenv("GINDJANGOAUTHTEST", "true")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
	c.Request.AddCookie(
		&http.Cookie{
			Name:    "sessionid",
			Value:   "error",
			Path:    "/",
			Expires: time.Now().Add(24 * time.Hour),
		})
	LoginRequired(c)

	if w.Code != 401 {
		t.Error(w.Code)
	}
}
