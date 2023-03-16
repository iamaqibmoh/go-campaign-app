package helper

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

func APIResponse(message string, code int, status string, data interface{}) web.Response {
	return web.Response{
		Meta: web.Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}

func RegisterUserResponseFormatter(user domain.User) web.RegisterUserResponse {
	return web.RegisterUserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      "tokeniniceritanya",
	}
}
