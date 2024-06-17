package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer

		// Process Request
		c.Next()

		// Stop timer

		entry := log.WithFields(log.Fields{
			"client_ip":  "util.GetClientIP(c)",
			"duration":   "duration",
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"user_id":    "",
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
