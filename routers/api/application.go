package api

import (
	"IAM/entities"
	"IAM/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApplicationController interface {
	AddApplication(c *gin.Context)
}

type applicationController struct {
	applicationService service.ApplicationService
}

func NewApplicationController(applicationService service.ApplicationService) *ApplicationController {
	return &applicationController{applicationService}
}

func (c *ApplicationController) AddApplication(ctx *gin.Context) {
	var request entities.Application
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.applicationService.CreateService(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, request)
}
