package main

import (
	"IAM/config"
	"IAM/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// RequestIDLogger mở rộng từ logrus.Logger để tự động thêm requestID
type RequestIDLogger struct {
	*logrus.Logger
}

// Middleware để thiết lập requestID và lưu vào context
func LoggerMiddleware(logger *RequestIDLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tạo và đặt requestID vào context
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		// Gọi next handler
		c.Next()
	}
}

// Phương thức để ghi log với requestID tự động
func (r *RequestIDLogger) InfoWithRequestID(requestID string, args ...interface{}) {
	r.WithField("requestID", requestID).Info(args...)
}

func main() {
	logger := &RequestIDLogger{Logger: logrus.New()}
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	config.InitConnection()
	defer config.DB.Close() // close when application was closed

	applicationController,
		scopeController,
		applicationScopeController,
		authenticationController := config.InitializeInjector()

	gin.SetMode("release")

	routersInit := routers.InitRouter(applicationController,
		scopeController,
		applicationScopeController,
		authenticationController)
	readTimeout := time.Duration(600) * time.Second
	writeTimeout := time.Duration(600) * time.Second
	endPoint := fmt.Sprintf(":%d", 8080)
	maxHeaderBytes := 1 << 20
	routersInit.Use(LoggerMiddleware(logger))

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logger.Info("start http server listening %s", endPoint)

	server.ListenAndServe()
}

// ginLogAdapter là một adapter để sử dụng Logrus với Gin
type ginLogAdapter struct {
	log *logrus.Logger
}

// Implement interface gin.Logger
func (g *ginLogAdapter) Println(v ...interface{}) {
	g.log.Info(v...)
}

// Để tránh xuất log lặp lại
func (g *ginLogAdapter) Printf(format string, v ...interface{}) {
	g.log.Infof(format, v...)
}
