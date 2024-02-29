package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web/model"
	"go-web/tools"
)

func main() {
	// 设置环境
	gin.SetMode(viper.GetString("server.Mode"))
	r := RegisterRoute()
	db := tools.GetMysqlDB()
	// 创建article表
	err := db.AutoMigrate(&model.Article{})

	if err != nil {
		fmt.Printf("创建article表失败: %s", err)
		return
	}

	port := viper.GetString("server.HttpPort")
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	fmt.Println(viper.Get("mysql.DbHost"))
}
