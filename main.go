package main

import (
	"IAM/config"
	"IAM/log"
	"IAM/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	config.ConfigLogging()
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

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Info(nil, "start http server listening %s")

	server.ListenAndServe()
}
