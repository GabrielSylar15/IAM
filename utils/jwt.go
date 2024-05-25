package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type Claims struct {
	ClientId string   `json:"client_id"`
	Type     string   `json:"type"`
	Scp      []string `json:"scp"`
	jwt.StandardClaims
}

var privateKey *ecdsa.PrivateKey

func loadPrivateKey() error {
	// Load private key from file
	privateKeyBytes, err := ioutil.ReadFile("E:\\Golang\\private_key.pem")
	if err != nil {
		return err
	}

	// Parse PEM block
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return fmt.Errorf("Error decoding PEM block")
	}

	// Parse ECDSA private key
	privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	return nil
}

func GenerateToke(claim *Claims, duration time.Duration, jwtScret interface{}) (string, error) {
	if privateKey == nil {
		if err := loadPrivateKey(); err != nil {
			return "", err
		}
	}

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
