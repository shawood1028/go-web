package model

type Article struct {
	Id          int    `json:"id"`
	ArticleId   int    `json:"articleId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}
