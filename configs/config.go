package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Configuration struct {
	App  App    `mapstructure:"APP" json:"APP" yaml:"APP"`
	Db   Db     `mapstructure:"DATABASE" json:"DATABASE" yaml:"DATABASE"`
	Log  Log    `mapstructure:"LOG" json:"LOG" yaml:"LOG"`
	Mode string `mapstructure:"RUN_MODE" json:"RUN_MODE" yaml:"RUN_MODE"`
	Jwt Jwt `yaml:"JWT"`
}

const (
	configName = "potato"
	configType = "yaml"
)

// 设置默认配置读取路径
var ConfigPath = []string{
	"~/.config/potato/",
	"~/.config/",
	"~/config/",
	"/etc/potato/",
	"./configs/",
}

type Application struct {
	Config      Configuration
	ConfigViper *viper.Viper
	Log         *logrus.Logger
	DB          *gorm.DB
}

var AppObj = new(Application)

func init() {
	LoadConfig()
}

func LoadConfig() *viper.Viper {
	v := viper.New()
	for _, path := range ConfigPath {
		v.AddConfigPath(path)
	}
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	//fmt.Println()
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	v.AutomaticEnv()
	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&AppObj.Config); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&AppObj.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
