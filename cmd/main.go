package main

import (
	"fmt"
	"gitee.com/bytesworld/tomato/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	cmd,_:= os.Getwd()
	fmt.Println(cmd)
	fmt.Println(configs.LoadConfig())
	r.Run(":"+configs.AppObj.Config.App.Port)
}
