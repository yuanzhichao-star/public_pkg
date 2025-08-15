package inits

import (
	"github.com/spf13/viper"
	"github.com/yuanzhichao-star/public_pkg/config"
	"log"
)

func InitViper() {
	viper.SetConfigFile("../config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置文件失败")
	}
	if err := viper.Unmarshal(&config.AppCong); err != nil {
		panic("解析配置文件失败")
	}
	log.Println("viper init success")
}
