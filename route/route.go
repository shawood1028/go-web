package route

import (
	"embed"
	"github.com/gin-gonic/gin"
	"go-web/model"
	"net/http"
	"strconv"
)

import (
	"html/template"
)

//go:embed templates/* templates/*/*
var f embed.FS

// RegisterRoute 配置路由信息，注册单个路由
func RegisterRoute() *gin.Engine {
	r := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.html"))
	r.StaticFS("/public", http.FS(f))
	r.SetHTMLTemplate(templ)
	r.GET("/", indexHandler)
	r.GET("/articles", articlesHandler)
	r.GET("/article/:id", articleHandler)
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
