package response

import "nongki/internal/domain"

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (r *LoginResponse) UserDomainToLoginResponse(data *domain.User, token string) *LoginResponse {
	r.Name = data.Name
	r.Email = data.Email
	r.Token = token

	return r
}
