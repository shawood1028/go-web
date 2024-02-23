package main

import (
	"fmt"
	"go-web/model"
	"strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		res := model.AddArticle(model.Article{
			Title:       "test title " + num,
			Description: "test desp " + num,
			Content:     "test content " + num,
		})
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}
