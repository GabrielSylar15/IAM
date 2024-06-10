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
		ctx.JSON(http.StatusInternalServerError, utils.BuildInternalErrorResponse())
	}
	ctx.JSON(http.StatusOK, utils.BuildSuccessResponse(entity))
}
