package service

import (
	"IAM/entities"
	"IAM/repository"
)

type scopeService struct {
	scopeRepository            repository.ScopeRepository
	applicationScopeRepository repository.ApplicationScopeRepository
	applicationRepository      repository.ApplicationRepository
}

type ScopeService interface {
	CreateScope(scope *entities.Scope) (entities.Scope, error)
	GetScope(clientID string) ([]*entities.Scope, error)
}

func InitScopeService(scopeRepository repository.ScopeRepository,
	applicationScopeRepository repository.ApplicationScopeRepository,
	applicationRepository repository.ApplicationRepository,
) ScopeService {
	return &scopeService{scopeRepository, applicationScopeRepository, applicationRepository}
}

func (c *scopeService) CreateScope(scope *entities.Scope) (entities.Scope, error) {
	result, err := c.scopeRepository.CreateScope(scope)

	application, _ := c.applicationRepository.GetApplicationByClientID(scope.OwnerClient)
	applcationScope := entities.ApplicationScope{
		ServiceID: application.ID,
		ScopeID:   result.ID,
		ClientID:  result.OwnerClient,
		CreatedBy: result.CreatedBy,
	}
	c.applicationScopeRepository.AssignScope(&applcationScope)
	return result, err
}

func (c *scopeService) GetScope(clientID string) ([]*entities.Scope, error) {
	return c.scopeRepository.GetScope(clientID)
}
