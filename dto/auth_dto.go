package dto

type TokenRequest struct {
	Scopes   []string
	ClientId string
}

type TokenResponse struct {
}
