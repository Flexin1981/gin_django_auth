package handlers

import (
	"github.com/Flexin1981/gin_django_auth/datalayer"
	"github.com/Flexin1981/gin_django_auth/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DjangoLoginHandler(c *gin.Context) {
	var inputJson Login

	if err := c.BindJSON(&inputJson); err != nil {
		c.JSON(http.StatusBadRequest, `{}`)
	}
	authService := datalayer.NewAuthUserService()

	user, err := authService.GetByUsername(inputJson.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, `{}`)
		return
	}

	if user.Verify(inputJson.Password) {
		sessionService := datalayer.NewSessionService()
		id, err := sessionService.Create(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, `{}`)
			return
		}
		c.SetCookie(middleware.DjangoSessionCookie, id, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, `{}`)
		return
	}

	c.JSON(http.StatusInternalServerError, `{}`)
	return
}
