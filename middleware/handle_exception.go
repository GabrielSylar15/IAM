package middleware

import (
	"IAM/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	httpResponse := utils.BuildErrorResponse("Intenal server error")
	c.AbortWithStatusJSON(500, httpResponse)
}
