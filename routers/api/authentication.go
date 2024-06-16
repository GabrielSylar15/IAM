package api

import (
	"IAM/dto"
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
}

func InitAuthenticationController(authenticationService service.AuthenticationService) *authenticationController {
	return &authenticationController{authenticationService}
}

func (auth *authenticationController) GetToken(ctx *gin.Context) {
	clientId := ctx.GetHeader("Username")
	var request dto.TokenRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildInternalErrorResponse())
		return
	}
	request.ClientId = clientId
	token, err := auth.AuthenticationService.GetToken(request)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, token)
}

//func GetToken(c *gin.Context) {
//	clientId, _, _ := c.Request.BasicAuth()
//	claims := utils.Claims{
//		ClientId:       clientId,
//		Type:           "direct",
//		Scp:            []string{"test.read:direct"},
//		StandardClaims: jwt.StandardClaims{},
//	}
//	token, err := utils.GenerateToke(&claims, 3*time.Hour, "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIAdDQx8UIEmWymoqfzsizTMaIxMUv6UcBh3k8uFVGGWpoAoGCCqGSM49\nAwEHoUQDQgAEl5ah/KBtLXUwQHrLGJD1kEj3yCViE+3qhDrSevx2cAg4g43sAxF1\npaekJiPxoJunzmU+LG5ULccdhP+zSNi71Q==\n-----END EC PRIVATE KEY-----\n")
//	if err != nil {
//		res := utils.BuildSuccessResponse("Token generation failed!")
//		c.JSON(http.StatusForbidden, res)
//	}
//	res := utils.BuildSuccessResponse(token)
//	c.JSON(http.StatusOK, res)
//}
