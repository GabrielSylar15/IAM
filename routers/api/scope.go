package api

import (
	"IAM/entities"
	"IAM/service"
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type scopeController struct {
	scopeService service.ScopeService
}

type ScopeController interface {
	CreateScope(c *gin.Context)
	GetScope(c *gin.Context)
}

func InitScopeController(scopeService service.ScopeService) ScopeController {
	return &scopeController{scopeService}
}

func (c *scopeController) CreateScope(ctx *gin.Context) {
	var request entities.Scope
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildInternalErrorResponse())
		return
	}
	entity, err := c.scopeService.CreateScope(&request)
	if err != nil {
		panic("Error creating scope")
	}
	ctx.JSON(http.StatusOK, utils.BuildSuccessResponse(entity))
}

func (c *scopeController) GetScope(ctx *gin.Context) {
	result, err := c.scopeService.GetScope(ctx.Param("client_id"))
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(http.StatusOK, utils.BuildSuccessResponse(result))
}
