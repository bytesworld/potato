package internal

import (
	"gitee.com/bytesworld/tomato/configs"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

var Logger *logrus.Logger

func init() {
	Logger = initLog()
}

func initLog() *logrus.Logger {
	log := logrus.New()
	logFileName := configs.AppObj.Config.Log.FileName
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: configs.AppObj.Config.Log.Format,
	}
	logPath := configs.AppObj.Config.Log.Path
	logName := path.Join(logPath, logFileName)
	var f *os.File
	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
	log.SetOutput(f)
	//log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:configs.AppObj.Config.Log.Format,
		FullTimestamp:true,
	})

	switch configs.AppObj.Config.Log.LogLevel {
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "WARN":
		log.SetLevel(logrus.WarnLevel)

	}
	return log
}
