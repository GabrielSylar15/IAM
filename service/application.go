package service

import (
	"IAM/entities"
	"IAM/repository"
	"time"
)

type ApplicationResponseStruct struct {
	ServiceName string    `json:"service_name"`
	KeyId       string    `json:"key_id""`
	ClientId    string    `json:"client_id""`
	CreatedAt   time.Time `json:"created_at""`
	UpdatedAt   time.Time `json:"updated_at""`
}

type ApplicationService interface {
	CreateService(form entities.Application) (ApplicationResponseStruct, error)
}

type applicationService struct {
	repository repository.ApplicationRepository
}

func InitApplicationService(repository repository.ApplicationRepository) ApplicationService {
	return &applicationService{repository}
}

func (s *applicationService) CreateService(entities.Application) (ApplicationResponseStruct, error) {
	return ApplicationResponseStruct{
		ServiceName: "GAM",
		KeyId:       "018ABCDAAGPOMNAVB_1505",
		ClientId:    "018ABCDAAGPOMNAVB",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}, nil
}
