package service

import (
	"IAM/dto"
	"IAM/entities"
	"IAM/repository"
	"IAM/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	log "github.com/sirupsen/logrus"
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
	GetToken(request dto.TokenRequest) (dto.TokenResponse, error)
	GetJWK(clientId string) (jwk.Key, error)
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

func (s *authenticationService) GetToken(request dto.TokenRequest) (dto.TokenResponse, error) {
	application, _ := s.ApplicationRepository.GetApplicationByClientID(request.ClientId)
	privateKey := application.PrivateKey
	applicationScope, _ := s.ScopeRepository.GetScope(request.ClientId)
	allScps := utils.Map(applicationScope, func(s *entities.Scope) string {
		return s.Scp
	})

	scps := make([]string, 0)
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
	duration, err := strconv.ParseInt(os.Getenv("DURATION_VALID_TOKEN"), 10, 64)
	token, _ := utils.GenerateToken(&claims, time.Duration(duration)*time.Hour, privateKey)
	response := dto.TokenResponse{
		TokenType:   "Bearer",
		AccessToken: token,
		Scopes:      scps,
		ExpiresIn:   duration * 60 * 60,
	}
	return response, err
}

func (s *authenticationService) GetJWK(clientId string) (jwk.Key, error) {
	log.Info("abcd")
	application, _ := s.ApplicationRepository.GetApplicationByClientID(clientId)
	publicKey := application.PublicKey
	return utils.ConvertToJWK(publicKey, application.KeyID, "ES256")
}
