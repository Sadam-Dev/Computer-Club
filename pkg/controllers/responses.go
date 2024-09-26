package controllers

type DefaultResponse struct {
	Message string `json:"message"`
}

func newDefaultResponse(message string) DefaultResponse {
	return DefaultResponse{
		Message: message,
	}
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}
