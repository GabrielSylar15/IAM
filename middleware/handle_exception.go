package middleware

import (
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleException() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			// log, handle, etc.
			log.Println(err)
		}

		c.JSON(http.StatusNotFound, utils.BuildErrorResponse("An errors has been occurs!"))
	}
}
