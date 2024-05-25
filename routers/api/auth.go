package api

import (
	"IAM/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Hello(c *gin.Context) {
	res := utils.BuildSuccessResponse("Hello, World!")
	c.JSON(http.StatusOK, res)
}

func GetToken(c *gin.Context) {
	clientId, _, _ := c.Request.BasicAuth()
	claims := utils.Claims{
		ClientId:       clientId,
		Type:           "direct",
		Scp:            []string{"test.read:direct"},
		StandardClaims: jwt.StandardClaims{},
	}
	token, err := utils.GenerateToke(&claims, 3*time.Hour, "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIAdDQx8UIEmWymoqfzsizTMaIxMUv6UcBh3k8uFVGGWpoAoGCCqGSM49\nAwEHoUQDQgAEl5ah/KBtLXUwQHrLGJD1kEj3yCViE+3qhDrSevx2cAg4g43sAxF1\npaekJiPxoJunzmU+LG5ULccdhP+zSNi71Q==\n-----END EC PRIVATE KEY-----\n")
	if err != nil {
		res := utils.BuildSuccessResponse("Token generation failed!")
		c.JSON(http.StatusForbidden, res)
	}
	res := utils.BuildSuccessResponse(token)
	c.JSON(http.StatusOK, res)
}
