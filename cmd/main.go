package main

import (
	"gitee.com/bytesworld/tomato/configs"
	"gitee.com/bytesworld/tomato/internal"
	"gitee.com/bytesworld/tomato/internal/middleware"
	"gitee.com/bytesworld/tomato/internal/routers"
)

func main() {
	// 初始化数据库
	configs.AppObj.DB = internal.InitDb()
	// 程序关闭前，释放数据库连接
	defer func() {
		if configs.AppObj.DB != nil {
			db, _ := configs.AppObj.DB.DB()
			db.Close()
		}
	}()

	r := routers.SetetupRouter()
	//v1 := r.Group("/v1")
	//fmt.Println(v1)
	r.Use(middleware.LoggerHander())
	//r.GET("/ping", func(c *gin.Context) {
	//	logger.Logger.Info("weidong")
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})r.GET("/ping", func(c *gin.Context) {
	//	logger.Logger.Info("weidong")
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})r.GET("/ping", func(c *gin.Context) {
	//	logger.Logger.Info("weidong")
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run(":" + configs.AppObj.Config.App.Port)
}
