package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/model"
	"net/http"
)

// RegisterRoute 配置路由信息，注册单个路由
func RegisterRoute() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", indexHandler)
	return r
}

func indexHandler(ctx *gin.Context) {
	articles := model.GetLastFiveArticles()
	data := gin.H{
		"msg":      "this is index",
		"articles": articles,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
