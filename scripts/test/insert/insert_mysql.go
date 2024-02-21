package main

import (
	"fmt"
	"go-web/model"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		res := model.AddArticle(model.Article{
			ID:          uint(rand.Int()),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       "test title " + num,
			Description: "test desp " + num,
			Content:     "test content " + num,
		})
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}
