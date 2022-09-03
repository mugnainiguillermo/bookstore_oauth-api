package http

import (
	"github.com/gin-gonic/gin"
	http2 "net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http2.StatusOK, map[string]string{"msg": "pong"})
}
