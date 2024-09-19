package response

import "nongki/internal/domain"

type LoginResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (r *LoginResponse) UserDomainToLoginResponse(data *domain.User, token *map[string]string) *LoginResponse {
	var resulToken map[string]string

	if token != nil {
		resulToken = *token
	}

	r.Name = data.Name
	r.Email = data.Email
	r.AccessToken = resulToken["access_token"]
	r.RefreshToken = resulToken["refresh_token"]

	return r
}
