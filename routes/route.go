package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/model"
	"net/http"
	"strconv"
)

// RegisterRoute 配置路由信息，注册单个路由
func RegisterRoute() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", indexHandler)
	r.GET("/articles", articlesHandler)
	r.GET("/article/:id", articleHandler)
	r.GET("/datas", datasHandler)
	return r
}

// 主页
func indexHandler(ctx *gin.Context) {
	articles := model.GetLastFiveArticles()
	data := gin.H{
		"msg":      "this is index",
		"articles": articles,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

// 获取所有文章
// TODO 分页
func articlesHandler(ctx *gin.Context) {
	articles := model.GetAllArticles()
	data := gin.H{
		"articles": articles,
	}
	ctx.HTML(http.StatusOK, "articles.html", data)
}

// 获取部分文章
func articleHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, _ := strconv.Atoi(id)
	article := model.GetOneArticle(aid)
	data := gin.H{
		"aid":     aid,
		"article": article,
	}
	ctx.HTML(http.StatusOK, "article.html", data)
}

// 数据中心
func datasHandler(ctx *gin.Context) {
	data := "datas"
	ctx.HTML(http.StatusOK, "datas.html", data)
}
