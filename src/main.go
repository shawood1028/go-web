package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web/model"
	"go-web/routes"
	"go-web/tools"
)

func main() {
	r := routes.RegisterRoute()
	db := tools.GetMysqlDB()
	// 创建article表
	err := db.AutoMigrate(&model.Article{})

	if err != nil {
		fmt.Printf("创建article表失败: %s", err)
		return
	}

	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	fmt.Println(viper.Get("mysql.DbHost"))
}
