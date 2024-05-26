package config

import (
	"IAM/repository"
	"IAM/routers/api"
	"IAM/service"
	"github.com/google/wire"
)

func InitInjector() (*api.ApplicationController, error) {
	wire.Build(api.NewApplicationController, service.NewApplicationService, repository.NewApplicationRepository, DB)
	return nil, nil
}
