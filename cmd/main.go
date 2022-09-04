package main

import (
	"fmt"
	"gitee.com/bytesworld/tomato/configs"
	"gitee.com/bytesworld/tomato/internal"
	"gitee.com/bytesworld/tomato/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println(configs.LoadConfig())
	r := gin.Default()
	r.Use(middleware.LoggerHander())
	r.GET("/ping", func(c *gin.Context) {
		internal.Logger.Info("weidong")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + configs.AppObj.Config.App.Port)
}
