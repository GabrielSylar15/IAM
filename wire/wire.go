//go:build !wireinject
// +build !wireinject

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
package wire

import (
	"IAM/repository"
	"IAM/routers/api"
	"IAM/service"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// using go run -mod=mod github.com/google/wire/cmd/wire to generate

func InitializeApplicationController(db *gorm.DB) api.ApplicationController {
	wire.Build(
		api.NewApplicationController,
		service.NewApplicationService,
		repository.NewApplicationRepository,
	)
	return nil
}
