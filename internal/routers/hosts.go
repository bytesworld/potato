package routers

import (
	"gitee.com/bytesworld/tomato/internal/controllers/host"
	"gitee.com/bytesworld/tomato/internal/middleware"
	"gitee.com/bytesworld/tomato/internal/service"
	"github.com/gin-gonic/gin"
)

func loadHostTRouter(router *gin.RouterGroup) {
	router.Use(middleware.JwtAurh(service.AppGuardName))
	router.GET("host/", host.Host)
}
