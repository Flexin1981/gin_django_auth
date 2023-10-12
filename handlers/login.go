package handlers

import (
	"log"
	"github.com/Flexin1981/gin_django_auth/datalayer"
	"github.com/Flexin1981/gin_django_auth/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/meehow/go-django-hashers"
)

func DjangoLoginHandler(c *gin.Context) {
	var inputJson Login

	if err := c.BindJSON(&inputJson); err != nil {
		c.JSON(http.StatusBadRequest, ``)
	}
	authService := datalayer.NewAuthUserService()

	user, err := authService.GetByUsername(inputJson.Username)
	if err != nil {
		log.Printf(`unauthorized user login attempt %v`, err)
		c.JSON(http.StatusUnauthorized, ``)
		return
	}

	ok, err := hashers.CheckPassword(inputJson.Password, user.Password)
	if err != nil {
		log.Printf(`password checker errored %v`, err)
		c.JSON(http.StatusInternalServerError, ``)
	}

	if ok {
		sessionService := datalayer.NewSessionService()
		session, err := sessionService.Create(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ``)
			return
		}

		c.SetCookie(middleware.DjangoSessionCookie, session.SessionKey, 3600, "/", "", false, true)
		c.JSON(http.StatusOK, `{}`)
		return
	}

	log.Printf(`unauthorized user login attempt %v`, err)
	c.JSON(http.StatusUnauthorized, ``)
	return
}
