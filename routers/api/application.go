package api

import (
	"IAM/entities"
	"IAM/service"
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApplicationController interface {
	AddApplication(c *gin.Context)
}

type applicationController struct {
	applicationService service.ApplicationService
}

func InitApplicationController(applicationService service.ApplicationService) ApplicationController {
	return &applicationController{applicationService}
}

func (c *applicationController) AddApplication(ctx *gin.Context) {
	var request entities.Application
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildInternalErrorResponse())
		return
	}

	result, err := c.applicationService.CreateService(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildInternalErrorResponse())
		return
	}
	ctx.JSON(http.StatusCreated, result)
}
