package host

import (
	"gitee.com/bytesworld/tomato/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Host(c *gin.Context) {
	logger.Logger.Info("login")
	var msg struct {
		Msg   string
		Hosts []string
	}
	msg.Msg = "hosts"
	msg.Hosts = []string{"aa", "zz"}
	c.JSON(http.StatusOK, msg)
}
