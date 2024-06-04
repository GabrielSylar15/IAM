package api

import (
	"IAM/service"
	"github.com/gin-gonic/gin"
)

type scopeController struct {
	scopeService *service.ScopeService
}

type ScopeController interface {
	CreateScope(c *gin.Context)
}

func InitScopeController(scopeService *service.ScopeService) ScopeController {
	return &scopeController{scopeService}
}

func (c *scopeController) CreateScope(ctx *gin.Context) {
	
}
