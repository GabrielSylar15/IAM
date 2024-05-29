package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GenerateECDSAPrivateKey() (string, *ecdsa.PrivateKey, error) {
	// Tạo khóa riêng tư với đường cong P-256
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", nil, fmt.Errorf("error generating ECDSA key: %w", err)
	}

	// Chuyển đổi khóa riêng tư sang định dạng DER
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("error marshaling private key: %w", err)
	}

	// Mã hóa khóa riêng tư thành định dạng PEM
	privateKeyBlock := pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKeyPEM := pem.EncodeToMemory(&privateKeyBlock)

	// Chuyển đổi PEM thành chuỗi
	privateKeyString := string(privateKeyPEM)
	return privateKeyString, privateKey, nil
}

func GenerateECDSAPublicKey(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", fmt.Errorf("error marshaling public key: %w", err)
	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyPEM := pem.EncodeToMemory(&publicKeyBlock)

	publicKeyString := string(publicKeyPEM)
	return publicKeyString, nil
}
