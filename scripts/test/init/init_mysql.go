package main

import (
	"fmt"
	"go-web/model"
	"go-web/tools"
)

func main() {
	fmt.Println("hello init_mysql")
	db := tools.GetMysqlDB()
	tx := db.Unscoped().Where("title like ?", "%test%").Delete(&model.Article{})

	if tx.Error != nil {
		fmt.Println("删除失败： ", tx.Error)
		return
	} else {
		fmt.Println("删除成功")
	}

	tx2 := db.Exec("ALTER TABLE articles AUTO_INCREMENT = 1;")
	if tx2.Error != nil {
		fmt.Println("重置id失败： ", tx.Error)
		return
	} else {
		fmt.Println("重置id成功")
	}
}
