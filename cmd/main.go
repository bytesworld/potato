package main

//import (
//	"context"
//	"gitee.com/bytesworld/tomato/configs"
//	"gitee.com/bytesworld/tomato/internal"
//	"gitee.com/bytesworld/tomato/internal/logger"
//	"gitee.com/bytesworld/tomato/internal/middleware"
//	"gitee.com/bytesworld/tomato/internal/routers"
//	"net/http"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//)
//
//func main() {
//	// 初始化数据库
//	configs.AppObj.DB = internal.DB
//	// 程序关闭前，释放数据库连接
//	defer func() {
//		if configs.AppObj.DB != nil {
//			db, _ := configs.AppObj.DB.DB()
//			db.Close()
//		}
//	}()
//
//	r := routers.SetetupRouter()
//	r.Use(middleware.LoggerHander())
//	// 原有启动服务
//	//r.Run(":" + configs.AppObj.Config.App.Port)
//
//	// 修改为收到信号后5s再停止
//	srv := &http.Server{
//		Addr:    ":" + configs.AppObj.Config.App.Port,
//		Handler: r,
//	}
//
//	go func() {
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			logger.Logger.Warn("listen: %s\n", err)
//		}
//	}()
//
//	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
//	quit := make(chan os.Signal)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//	logger.Logger.Info("Shutdown Server ...")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		logger.Logger.Fatal("Server Shutdown:", err)
//	}
//	logger.Logger.Println("Server exiting")
//}
import (
	"fmt"
	"gitee.com/bytesworld/tomato/internal/service/firewall"
)
func main()  {
	pcap:=firewall.NewPcap()
	fmt.Println(pcap.GetDevs())
}