package internal

import (
	"gitee.com/bytesworld/tomato/configs"
	"gitee.com/bytesworld/tomato/internal/logger"
	"gitee.com/bytesworld/tomato/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	logger.Logger = logger.initLog()
}

func InitDb() *gorm.DB {
	dbConfig := configs.AppObj.Config.Db
	logger.Logger.Infof("db config is %+v",dbConfig)
	db, err := gorm.Open(sqlite.Open(dbConfig.Name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//// Migrate the schema
	db.AutoMigrate(&models.User{})
	logger.Logger.Info("完成数据库初始化")
	return db

}
