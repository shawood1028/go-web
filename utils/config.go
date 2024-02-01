package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// Init InitConfig 读取配置文件
func Init() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	workDir, _ := os.Getwd()
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath(workDir + "/etc")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("配置文件初始化")
	return
}
