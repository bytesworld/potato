package routers

import (
	"gitee.com/bytesworld/tomato/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func auth(c *gin.Context) {
	logger.Logger.Info("login")
	c.JSON(http.StatusOK, gin.H{
		"message": "auth",
	})
}

func loadAuthRouter(router *gin.RouterGroup) {
	router.GET("auth/", auth)
}
