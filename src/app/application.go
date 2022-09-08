package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mugnainiguillermo/bookstore_oauth-api/src/http"
	"github.com/mugnainiguillermo/bookstore_oauth-api/src/repository/db"
	"github.com/mugnainiguillermo/bookstore_oauth-api/src/repository/rest"
	"github.com/mugnainiguillermo/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/ping", http.Ping)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.SetTrustedProxies(nil)
	router.Run("localhost:9001")
}
