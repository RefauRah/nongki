package response

import "nongki/internal/domain"

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (r *RegisterResponse) UserDomainToRegisterResponse(data *domain.User, token string) *RegisterResponse {
	r.Name = data.Name
	r.Email = data.Email
	r.Token = token

	return r
}
