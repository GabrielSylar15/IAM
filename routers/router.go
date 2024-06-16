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
	applicationScopeController api.ApplicationScopeController,
	authenticationController api.AuthenticationController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.Use(middleware.BasicAuth(), gin.CustomRecovery(middleware.ErrorHandler))

		authGroup := apiv1.Group("/auth")
		{
			authGroup.POST("/clients/token", authenticationController.GetToken)
		}

		applicationGroup := apiv1.Group("/application")
		{
			applicationGroup.POST("/", applicationController.AddApplication)
			applicationGroup.GET("/:id", applicationController.GetApplication)
		}

		scopeGroup := apiv1.Group("/scope")
		{
			scopeGroup.POST("/", scopeController.CreateScope)
			scopeGroup.GET("/:client_id", scopeController.GetScope)
		}

		applicationScopeGroup := apiv1.Group("/appication/scope")
		{
			applicationScopeGroup.POST("/", applicationScopeController.AssignScope)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.BuildErrorResponse("Invalid resouce!"))
	})
	return r
}
