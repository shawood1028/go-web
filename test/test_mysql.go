package main

import (
	"fmt"
	"go-web/model"
	"go-web/tools"
	"time"
)

func main() {
	// 配置初始化
	tools.Init()
	tools.Init()
	// 测试数据
	tt := time.Now().Format("2006-01-02 15:04:05")
	article := model.Article{ArticleId: '1', Title: "test1", Description: "ttt", Content: "test content", CreateTime: tt, UpdateTime: tt}
	fmt.Print("article数据为: ", article)
	db := tools.GetMysqlDB()
	result := db.Create(&article)
	fmt.Println(result.Error)
}
