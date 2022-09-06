package auth

import (
	"gitee.com/bytesworld/tomato/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	logger.Logger.Info("login")
	c.JSON(http.StatusOK, gin.H{
		"message": "auth",
	})
}
