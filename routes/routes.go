package routes

import (
	"github.com/gin-gonic/gin"
	"go-superset/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("api/auth/login", controller.Login)
	return r
}
