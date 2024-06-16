package service

import (
	"IAM/dto"
	"IAM/entities"
	"IAM/repository"
	"IAM/utils"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

type authenticationService struct {
	repository.ApplicationRepository
	repository.ScopeRepository
	repository.ApplicationScopeRepository
}

type AuthenticationService interface {
	GetToken(request dto.TokenRequest) (string, error)
}

func InitAuthenticationService(applicationRepository repository.ApplicationRepository,
	scopeRepository repository.ScopeRepository,
	applicationScopeRepository repository.ApplicationScopeRepository) AuthenticationService {
	return &authenticationService{
		ApplicationRepository:      applicationRepository,
		ScopeRepository:            scopeRepository,
		ApplicationScopeRepository: applicationScopeRepository,
	}
}

func (s *authenticationService) GetToken(request dto.TokenRequest) (string, error) {
	application, _ := s.ApplicationRepository.GetApplicationByClientID(request.ClientId)
	privateKey := application.PrivateKey
	applicationScope, _ := s.ScopeRepository.GetScope(request.ClientId)
	allScps := utils.Map(applicationScope, func(s *entities.Scope) string {
		return s.Scp
	})

	var scps []string
	for _, item := range request.Scopes {
		if utils.Contains(allScps, item) {
			scps = append(scps, item)
		}
	}

	claims := utils.Claims{
		ClientId:       application.ClientID,
		Type:           "",
		Scp:            scps,
		StandardClaims: jwt.StandardClaims{},
	}
	duration, _ := strconv.ParseInt(os.Getenv("DURATION_VALID_TOKEN"), 10, 64)
	return utils.GenerateTokenByPrivateKey(&claims, time.Duration(duration)*time.Hour, privateKey)
}
