package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

type UserService interface {
	RegisterUser(input web.RegisterUserInput) (domain.User, error)
}
