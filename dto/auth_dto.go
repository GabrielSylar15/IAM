package dto

type TokenRequest struct {
	Scopes   []string `json:"scps"`
	ClientId string   `json:"client_id"`
}

type TokenResponse struct {
}
