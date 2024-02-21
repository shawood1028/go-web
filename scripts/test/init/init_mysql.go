package main

import (
	"fmt"
	"go-web/model"
	"go-web/tools"
)

func main() {
	fmt.Println("hello init_mysql")
	db := tools.GetMysqlDB()
	tx := db.Where("title like ?", "%test%").Delete(&model.Article{})
	if tx.Error != nil {
		fmt.Println("删除失败： ", tx.Error)
		return
	} else {
		fmt.Println("删除成功")
	}
}
