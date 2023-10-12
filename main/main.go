package main

import (
	"net/http"
	"github.com/Flexin1981/gin_django_auth/handlers"
	"github.com/Flexin1981/gin_django_auth/middleware"
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  }

func main() {
  r := gin.Default()
  r.POST("/login", handlers.DjangoLoginHandler)
  r.GET("/ping", middleware.LoginRequired, handler)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}