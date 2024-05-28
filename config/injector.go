package config

import (
	"IAM/repository"
	"IAM/routers/api"
	"IAM/service"
)

var applicationRepository repository.ApplicationRepository
var applicationService service.ApplicationService
var applicationController api.ApplicationController

func InitializeInjector() api.ApplicationController {
	// ApplicationController
	applicationRepository = repository.InitApplicationRepository(DB)
	applicationService = service.InitApplicationService(applicationRepository)
	applicationController = api.InitApplicationController(applicationService)
	
	return applicationController
}
