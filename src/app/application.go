package app

import (
	"bookstore_oauth-api/src/domain/access_token"
	http2 "bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http2.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/ping", http2.Ping)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.SetTrustedProxies(nil)
	router.Run("localhost:9001")
}
