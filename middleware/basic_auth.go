package middleware

import (
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && username == "admin" && password == "password" {
			c.Request.Header.Set("username", username)
			c.Next()
		} else {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("Unauthentication, please try again!"))
		}
	}
}
