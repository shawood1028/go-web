package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	MySQL MySQLConfig
}

type AppConfig struct {
	HttpPort int
}

type MySQLConfig struct {
	DbUser     string
	DbPassWord string
	DbHost     string
	DbPort     string
	DbName     string
}

// InitConfig 读取配置文件
func InitConfig() Config {
	var config Config
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./etc/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	// TODO handle err
	viper.Unmarshal(&config)
	fmt.Println("config: ", config)
	return config
}
