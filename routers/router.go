package routers

import (
	"IAM/middleware"
	"IAM/routers/api"
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(applicationController api.ApplicationController,
	scopeController api.ScopeController,
	applicationScopeController api.ApplicationScopeController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.BasicAuth(),
		gin.CustomRecovery(middleware.ErrorHandler))
	{
		apiv1.GET("/hello", api.Hello)
		apiv1.POST("/token", api.GetToken)
		apiv1.POST("/application", applicationController.AddApplication)
		apiv1.GET("/application/:id", applicationController.GetApplication)
		apiv1.POST("/scope", scopeController.CreateScope)
		apiv1.GET("/scope/:client_id", scopeController.GetScope)
		apiv1.POST("/application/scope", applicationScopeController.AssignScope)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.BuildErrorResponse("Invalid resouce!"))
	})
	return r
}
