package config

import (
	"IAM/repository"
	"IAM/routers/api"
	"IAM/service"
)

var applicationRepository repository.ApplicationRepository
var applicationService service.ApplicationService
var applicationController api.ApplicationController

func InitializeInjector() (api.ApplicationController, api.ScopeController) {
	// ApplicationController
	applicationRepository = repository.InitApplicationRepository(DB)
	applicationService = service.InitApplicationService(applicationRepository)
	applicationController = api.InitApplicationController(applicationService)

	scopeRepository := repository.InitScopeRepository(DB)
	scopeService := service.InitScopeService(scopeRepository)
	scopeController := api.InitScopeController(scopeService)

	return applicationController, scopeController
}
