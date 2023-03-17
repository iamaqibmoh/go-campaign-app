package formatter

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

func UserResponseFormatter(user domain.User, token string) web.RegisterUserResponse {
	return web.RegisterUserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
}
