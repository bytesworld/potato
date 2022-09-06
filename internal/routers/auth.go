package routers

import (
	"gitee.com/bytesworld/tomato/internal/apps/auth"
	"github.com/gin-gonic/gin"
)

func loadAuthRouter(router *gin.RouterGroup) {
	router.GET("auth/", auth.Auth)
	router.GET("add/", auth.AddUser)
}
