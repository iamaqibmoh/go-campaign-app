package formatter

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

func UserResponseFormatter(user domain.User, token string) web.UserResponse {
	return web.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		ImageURL:   user.Avatar,
	}
}
