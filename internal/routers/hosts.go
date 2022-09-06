package routers

import (
	"gitee.com/bytesworld/tomato/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func host(c *gin.Context) {
	logger.Logger.Info("login")
	var msg struct {
		Msg   string
		Hosts []string
	}
	msg.Msg = "hosts"
	msg.Hosts = []string{"aa", "zz"}
	c.JSON(http.StatusOK, msg)
}
func loadHostTRouter(router *gin.RouterGroup) {
	router.GET("host/", host)
}
