package api

import (
	"IAM/dto"
	"IAM/log"
	"IAM/service"
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authenticationController struct {
	service.AuthenticationService
}

type AuthenticationController interface {
	GetToken(ctx *gin.Context)
	GetJWK(ctx *gin.Context)
}

func InitAuthenticationController(authenticationService service.AuthenticationService) *authenticationController {
	return &authenticationController{authenticationService}
}

// TODO: handle panic error message
func (auth *authenticationController) GetToken(ctx *gin.Context) {
	clientId := ctx.GetHeader("Username")
	var request dto.TokenRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildInternalErrorResponse())
		return
	}
	request.ClientId = clientId
	response, err := auth.AuthenticationService.GetToken(request)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, utils.BuildSuccessResponse(response))
}

func (auth *authenticationController) GetJWK(ctx *gin.Context) {
	clientId := ctx.GetHeader("X-User-name")
	log.Info(ctx.Request.Context(), "%s get jwk", clientId)
	response, err := auth.AuthenticationService.GetJWK(ctx.Request.Context(), clientId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, utils.BuildSuccessResponse(response))
}
