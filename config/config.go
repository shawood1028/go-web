package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	// TODO 读取环境变量
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	// 基于项目根路径搜索
	viper.AddConfigPath("./etc")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
