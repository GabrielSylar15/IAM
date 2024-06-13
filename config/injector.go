package config

import (
	"IAM/repository"
	"IAM/routers/api"
	"IAM/service"
)

func InitializeInjector() (api.ApplicationController, api.ScopeController, api.ApplicationScopeController) {
	// ApplicationController
	applicationRepository := repository.InitApplicationRepository(DB)
	applicationService := service.InitApplicationService(applicationRepository)
	applicationController := api.InitApplicationController(applicationService)

	applicationScopeRepository := repository.InitApplicationScopeRepository(DB)
	applicationScopeService := service.InitializeApplicationScopeService(applicationScopeRepository)
	applicationScopeController := api.InitializeApplicationScopeController(applicationScopeService)

	scopeRepository := repository.InitScopeRepository(DB)
	scopeService := service.InitScopeService(scopeRepository, applicationScopeRepository, applicationRepository)
	scopeController := api.InitScopeController(scopeService)

	return applicationController, scopeController, applicationScopeController
}
