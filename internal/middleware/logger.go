package middleware

import (
	"fmt"
	"gitee.com/bytesworld/tomato/internal"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerHander() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		durationTime := fmt.Sprintf("%6v", endTime.Sub(startTime))
		reqMethod := c.Request.Method
		reqUrl := c.Request.RequestURI
		status := c.Writer.Status()
		clientIp := c.ClientIP()
		internal.Logger.WithFields(logrus.Fields{
			"start_time": startTime,
			"duration":   durationTime,
			"method":     reqMethod,
			"url":        reqUrl,
			"status":     status,
			"client_ip":  clientIp,
		}).Info("access")

	}

}
