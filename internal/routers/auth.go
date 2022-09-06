package routers

import (
	"gitee.com/bytesworld/tomato/internal/controllers/auth"
	"github.com/gin-gonic/gin"
)

func loadAuthRouter(router *gin.RouterGroup) {
	router.GET("auth/", auth.Auth)
	router.GET("user/", auth.GetUsers)
	router.POST("user/", auth.CreateUser)
	router.GET("user/:id/", auth.GetUser)
	router.PUT("user/:id/", auth.UpdateUser)
	router.DELETE("user/:id/", auth.DeleteUser)
}
