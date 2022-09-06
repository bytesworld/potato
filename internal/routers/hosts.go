package routers

import (
	"gitee.com/bytesworld/tomato/internal/controllers/host"
	"github.com/gin-gonic/gin"
)

func loadHostTRouter(router *gin.RouterGroup) {
	router.GET("host/", host.Host)
}
