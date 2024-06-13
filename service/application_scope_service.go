package service

import (
	"IAM/entities"
	"IAM/repository"
)

type applicationScopeService struct {
	repository.ApplicationScopeRepository
}

type ApplicationScopeService interface {
	AssignScope(applicationScope *entities.ApplicationScope) (*entities.ApplicationScope, error)
}

func InitializeApplicationScopeService(applicationScopeRepository repository.ApplicationScopeRepository) ApplicationScopeService {
	return &applicationScopeService{applicationScopeRepository}
}

func (service *applicationScopeService) AssignScope(applicationScope *entities.ApplicationScope) (*entities.ApplicationScope, error) {
	return service.ApplicationScopeRepository.AssignScope(applicationScope)
}
