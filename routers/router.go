package routers

import (
	"IAM/middleware"
	"IAM/routers/api"
	"IAM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(applicationController api.ApplicationController,
	scopeController api.ScopeController,
	applicationScopeController api.ApplicationScopeController,
	authenticationController api.AuthenticationController) *gin.Engine {
	r := gin.New()
	//r.Use(func(c *gin.Context) {
	//	rid := uuid.New().String()
	//	c.Request.Header.Set("Request-Id", rid)
	//	log.WithField("requestID", rid).Info("New request")
	//	c.Next()
	//})
	//r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
	//	Formatter: func(param gin.LogFormatterParams) string {
	//		// Ghi log bằng Logrus
	//		log.WithFields(log.Fields{
	//			"requestID": param.Request.Header.Get("Request-Id"), // Lấy RID từ header nếu đã có
	//			"status":    param.StatusCode,
	//			"method":    param.Method,
	//			"path":      param.Path,
	//			"clientIP":  param.ClientIP,
	//			"latency":   param.Latency,
	//			"userAgent": param.Request.UserAgent(),
	//			"referer":   param.Request.Referer(),
	//		}).Info("Request handled")
	//
	//		// Trả về chuỗi rỗng vì log đã được xử lý bởi Logrus
	//		return ""
	//	},
	//	Output: gin.DefaultWriter, // Sử dụng output của Logrus
	//}))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.Use(middleware.BasicAuth(), gin.CustomRecovery(middleware.ErrorHandler))

		authGroup := apiv1.Group("/auth")
		{
			authGroup.POST("/clients/token", authenticationController.GetToken)
			authGroup.GET("/clients/jwk", authenticationController.GetJWK)
		}

		applicationGroup := apiv1.Group("/application")
		{
			applicationGroup.POST("/", applicationController.AddApplication)
			applicationGroup.GET("/:id", applicationController.GetApplication)
		}

		scopeGroup := apiv1.Group("/scope")
		{
			scopeGroup.POST("/", scopeController.CreateScope)
			scopeGroup.GET("/:client_id", scopeController.GetScope)
		}

		applicationScopeGroup := apiv1.Group("/appication/scope")
		{
			applicationScopeGroup.POST("/", applicationScopeController.AssignScope)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.BuildErrorResponse("Invalid resouce!"))
	})
	return r
}
