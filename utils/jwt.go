package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"os"
	"time"
)

type Claims struct {
	ClientId string   `json:"client_id"`
	Type     string   `json:"type"`
	Scp      []string `json:"scp"`
	jwt.StandardClaims
}

func loadPrivateKey(jwtScret string) *ecdsa.PrivateKey {
	// Load private key from file
	privateKeyBytes := []byte(jwtScret)

	// Parse PEM block
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		panic("failed to parse PEM block")
	}

	// Parse ECDSA private key
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return privateKey
}

func GenerateToken(claim *Claims, duration time.Duration, jwtScret string) (string, error) {
	privateKey := loadPrivateKey(jwtScret)

	now := time.Now()
	expire := now.Add(duration)
	claim.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expire.Unix(),
		Issuer:    "iam",
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	token, err := tokenClaims.SignedString(privateKey)
	return token, err
}

func LoadPublicKey(path string) (*ecdsa.PublicKey, error) {
	pubKeyPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key PEM file: %v", err)
	}

	block, _ := pem.Decode(pubKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DER encoded public key: %v", err)
	}

	ecPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not ECDSA public key")
	}

	return ecPubKey, nil
}

func ConvertToJWK(pubKey *ecdsa.PublicKey, kid string, alg string) (jwk.Key, error) {
	jwkKey, err := jwk.New(pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWK: %v", err)
	}

	jwkKey.Set(jwk.KeyIDKey, kid)
	jwkKey.Set(jwk.AlgorithmKey, alg)
	jwkKey.Set(jwk.KeyUsageKey, "sig")
	return jwkKey, nil
}
