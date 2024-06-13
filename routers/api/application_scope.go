package api

import (
	"IAM/entities"
	"IAM/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApplicationScopeController interface {
	AssignScope(ctx *gin.Context)
}

type applicationScopeController struct {
	service.ApplicationScopeService
}

func InitializeApplicationScopeController(applicationScope service.ApplicationScopeService) ApplicationScopeController {
	return &applicationScopeController{applicationScope}
}

func (c *applicationScopeController) AssignScope(ctx *gin.Context) {
	var request entities.ApplicationScope
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	result, err := c.ApplicationScopeService.AssignScope(&request)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusCreated, result)
}
