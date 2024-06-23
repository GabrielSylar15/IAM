package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"strings"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = strings.ReplaceAll(uuid.New().String(), "-", "")[:13]
		}
		ctx := context.WithValue(c.Request.Context(), "rid", requestID)
		c.Request = c.Request.WithContext(ctx)

		log.WithFields(logrus.Fields{
			"rid": requestID,
			"st":  c.Writer.Status(),
			"mt":  c.Request.Method,
			"p":   c.Request.URL.Path,
		}).Info("Handled request")
		c.Next()
	}
}

//func JSONLogMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		rid, _ := c.Get("r-id")
//		if rid == nil {
//			rid = uuid.New().String()
//		}
//		ctx := ctx.WithRequestID(c.Request.Context(), rid)
//		c.Request = c.Request.WithContext(ctx)
//		entry := log.WithFields(log.Fields{
//			"client_ip":  "192.168.1.1",
//			"duration":   "duration",
//			"method":     c.Request.Method,
//			"path":       c.Request.RequestURI,
//			"status":     c.Writer.Status(),
//			"user_id":    "",
//			"referrer":   c.Request.Referer(),
//			"request_id": rid,
//		})
//
//		if c.Writer.Status() >= 500 {
//			entry.Error(c.Errors.String())
//		} else {
//			entry.Info("")
//		}
//		c.Next()
//	}
//}
