package service

import (
	"IAM/entities"
	"IAM/repository"
)

type scopeService struct {
	scopeRepository repository.ScopeRepository
}

type ScopeService interface {
	CreateScope(scope *entities.Scope) (entities.Scope, error)
	GetScope(clientID string) ([]*entities.Scope, error)
}

func InitScopeService(scopeRepository repository.ScopeRepository) ScopeService {
	return &scopeService{scopeRepository}
}

func (c *scopeService) CreateScope(scope *entities.Scope) (entities.Scope, error) {
	return c.scopeRepository.CreateScope(scope)
}

func (c *scopeService) GetScope(clientID string) ([]*entities.Scope, error) {
	return c.scopeRepository.GetScope(clientID)
}
