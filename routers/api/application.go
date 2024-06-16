package api

import (
	"IAM/entities"
	"IAM/service"
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ApplicationController interface {
	AddApplication(c *gin.Context)
	GetApplication(c *gin.Context)
}

type applicationController struct {
	applicationService service.ApplicationService
}

func InitApplicationController(applicationService service.ApplicationService) ApplicationController {
	return &applicationController{applicationService}
}

// TODO: encrypt and decrypt with salt private key
func (c *applicationController) AddApplication(ctx *gin.Context) {
	var request entities.Application
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildInternalErrorResponse())
		return
	}

	result := c.applicationService.CreateApplication(request)
	ctx.JSON(http.StatusCreated, result)
}

func (c *applicationController) GetApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	entity := c.applicationService.GetApplication(id)
	ctx.JSON(http.StatusOK, utils.BuildSuccessResponse(entity))
}
