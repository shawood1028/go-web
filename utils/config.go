package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// InitConfig 读取配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath(workDir + "/etc")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}
