package main

import (
	"fmt"
	"go-web/model"
	"go-web/tools"
	"math/rand"
	"time"
)

func main() {
	// 配置初始化
	db := tools.GetMysqlDB()
	article := model.Article{
		ID:          uint(rand.Int()),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Content:     "test content",
		Title:       "test title",
		Description: "test desc",
	}
	addRes := model.AddArticle(article)
	if addRes.Error != nil {
		_ = fmt.Sprintf("%v", addRes)
	} else {
		fmt.Println("添加数据成功")
	}
	var art []model.Article
	err := db.First(art)
	if err != nil {
		fmt.Println("查询失败: ", err)
	}
	fmt.Printf("查询结果: %v", art)
}
