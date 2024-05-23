package api

import (
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	res := utils.BuildSuccessResponse("Hello, World!")
	c.JSON(http.StatusOK, res)
}
