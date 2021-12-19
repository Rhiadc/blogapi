package utils

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh-token"`
}
