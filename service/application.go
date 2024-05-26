package service

import (
	"IAM/entities"
	"IAM/repository"
)

type ApplicationService interface {
	CreateService(form entities.Application) error
}

type applicationService struct {
	repository repository.ApplicationRepository
}

func NewApplicationService(repository repository.ApplicationRepository) ApplicationService {
	return &applicationService{repository}
}

func (s *applicationService) CreateService(entities.Application) error {
	return nil
}
