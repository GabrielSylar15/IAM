package routers

import (
	"IAM/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/hello", api.Hello)
	}
	return r
}
