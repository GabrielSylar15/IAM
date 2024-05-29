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
	CreateService(form entities.Application) (ApplicationResponseStruct, error)
}

type applicationService struct {
	repository repository.ApplicationRepository
}

func InitApplicationService(repository repository.ApplicationRepository) ApplicationService {
	return &applicationService{repository}
}

func (s *applicationService) CreateService(request entities.Application) (ApplicationResponseStruct, error) {
	entity := BuildApplicationEntity()
	entity.ServiceName = request.ServiceName
	return ApplicationResponseStruct{
		ServiceName:  entity.ServiceName,
		KeyId:        entity.KeyID,
		ClientId:     entity.ClientID,
		ClientSecret: entity.ClientSecret,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}, nil
}

func BuildApplicationEntity() *entities.Application {
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
		ServiceName:  "",
		PublicKey:    publicKeyStr,
		PrivateKey:   privateKeyStr,
		CreatedBy:    0,
		KeyID:        keyId,
	}
}
