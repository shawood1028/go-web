package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"go-web/model"
	"html/template"
	"io/fs"
	"net/http"
	"strconv"
)

//go:embed templates/* static/*
var f embed.FS

// RegisterRoute 配置路由信息，注册单个路由
func RegisterRoute() *gin.Engine {
	r := gin.Default()
	// 初始化默认静态资源
	fe, _ := fs.Sub(f, "static")
	r.StaticFS("/static", http.FS(fe))
	// 加载模板
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(f, "templates/*.html")))
	r.GET("/", indexHandler)
	r.GET("/about", aboutHandler)
	r.GET("/contact", contactHandler)
	r.GET("/post", postHandler)
	r.GET("/articles", articlesHandler)
	r.GET("/article/:id", articleHandler)
	return r
}

// 主页
func indexHandler(ctx *gin.Context) {
	data := gin.H{
		"msg": "this is index",
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func contactHandler(ctx *gin.Context) {
	data := gin.H{
		"msg": "this is index",
	}
	ctx.HTML(http.StatusOK, "contact.html", data)
}

func aboutHandler(ctx *gin.Context) {
	aboutInfo1 := "一个喜欢胡思乱想的理想主义者。"
	aboutInfo2 := "兼职程序员。"
	data := gin.H{
		"msg1": aboutInfo1,
		"msg2": aboutInfo2,
	}
	ctx.HTML(http.StatusOK, "about.html", data)
}

func postHandler(ctx *gin.Context) {
	data := gin.H{
		"msg": "this is index",
	}
	ctx.HTML(http.StatusOK, "post.html", data)
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
