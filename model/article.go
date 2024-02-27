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
	result := db.Create(&article)
	return result
}

// GetOneArticle 获取指定id的文章
func GetOneArticle(aid int) Article {
	db := tools.GetMysqlDB()
	var article Article
	res := db.Find(&article, "id = ?", aid)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	fmt.Println(article)
	return article
}

// GetLastFiveArticles 获取五篇最近文章
func GetLastFiveArticles() []Article {
	db := tools.GetMysqlDB()
	var articles []Article
	res := db.Limit(5).Order("created_at desc").Find(&articles).Scan(&articles)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return articles
}

// GetAllArticles 获取所有文章（暂时没有分页）
func GetAllArticles() []Article {
	db := tools.GetMysqlDB()
	var articles []Article
	res := db.Order("created_at desc").Find(&articles).Scan(&articles)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return articles
}
