package middleware

import "github.com/gin-gonic/gin"

const (
	DjangoSessionCookieNotFound  string = "django session cookie not found in the incoming request"
	DjangoSessionCookieIsBlank   string = "django session cookie is blank in the incoming request"
	DjangoSessionCookieIsExpired string = "django session cookie is expired in the incoming request"
)

var (
	UnauthorizedJson = gin.H{"error": "unauthorized"}
)
