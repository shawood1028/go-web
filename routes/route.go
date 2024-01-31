package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterRoute 配置路由信息，注册单个路由
func RegisterRoute() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/article", articleHandler)
	r.GET("/", indexHandler)
	return r
}

func indexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "this is index",
	})
}

func articleHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "article.html", gin.H{
		"msg": "this is article",
	})
}
