package service

import (
	"IAM/entities"
	"IAM/repository"
	"IAM/utils"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type ApplicationResponseStruct struct {
	ServiceName  string    `json:"service_name"`
	KeyId        string    `json:"key_id""`
	ClientId     string    `json:"client_id""`
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at""`
	UpdatedAt    time.Time `json:"updated_at""`
}

type ApplicationService interface {
	CreateApplication(form entities.Application) ApplicationResponseStruct
	GetApplication(id int64) ApplicationResponseStruct
}

type applicationService struct {
	repository repository.ApplicationRepository
}

func InitApplicationService(repository repository.ApplicationRepository) ApplicationService {
	return &applicationService{repository}
}

func (s *applicationService) CreateApplication(request entities.Application) ApplicationResponseStruct {
	entity := BuildApplicationEntity(request)
	entity.ServiceName = request.ServiceName
	entity.CreatedBy = request.CreatedBy
	_, err := s.repository.CreateApplication(entity)
	if err != nil {
		panic(err)
	}

	return ApplicationResponseStruct{
		ServiceName:  entity.ServiceName,
		KeyId:        entity.KeyID,
		ClientId:     entity.ClientID,
		ClientSecret: entity.ClientSecret,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

func BuildApplicationEntity(request entities.Application) *entities.Application {
	clientId := strings.ReplaceAll(strings.ToUpper(uuid.New().String()), "-", "")
	clientSecret := strings.ReplaceAll(uuid.New().String(), "-", "")
	keyId := clientId + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	privateKeyStr, privateKey, error := utils.GenerateECDSAPrivateKey()
	publicKeyStr, error := utils.GenerateECDSAPublicKey(privateKey)
	if error != nil {
		panic(error)
	}
	return &entities.Application{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		ServiceName:  request.ServiceName,
		PublicKey:    publicKeyStr,
		PrivateKey:   privateKeyStr,
		KeyID:        keyId,
	}
}

func (s *applicationService) GetApplication(id int64) ApplicationResponseStruct {
	entity, err := s.repository.GetApplication(id)
	if err != nil {
		panic(err)
	}
	return ApplicationResponseStruct{
		ServiceName:  entity.ServiceName,
		KeyId:        entity.KeyID,
		ClientId:     entity.ClientID,
		ClientSecret: entity.ClientSecret,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
