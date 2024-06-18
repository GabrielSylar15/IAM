package dto

type TokenRequest struct {
	Scopes   []string `json:"scps"`
	ClientId string   `json:"client_id"`
}

type TokenResponse struct {
	TokenType   string   `json:"token_type"`
	AccessToken string   `json:"access_token"`
	Scopes      []string `json:"scope"`
	ExpiresIn   int64    `json:"expires_in"`
}
