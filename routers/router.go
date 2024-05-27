package routers

import (
	"IAM/middleware"
	"IAM/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(controller api.ApplicationController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.BasicAuth())
	{
		apiv1.GET("/hello", api.Hello)
		apiv1.POST("/token", api.GetToken)
		apiv1.POST("/add", controller.AddApplication)
	}
	return r
}
