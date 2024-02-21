package model

import (
	"fmt"
	"go-web/tools"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func AddArticle(article Article) *gorm.DB {
	db := tools.GetMysqlDB()
	result := db.Create(article)
	return result
}

func GetOneArticle() *gorm.DB {
	db := tools.GetMysqlDB()
	var art Article
	res := db.First(&art)
	fmt.Println(res)
	return nil
}

// GetLastFiveArticles 获取五篇最近文章
func GetLastFiveArticles() []Article {
	db := tools.GetMysqlDB()
	var articles []Article
	res := db.Limit(5).Order("created_at desc").Find(&articles).Scan(&articles)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	fmt.Println(articles)
	return articles
}
