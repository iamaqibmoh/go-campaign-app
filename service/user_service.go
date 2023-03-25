package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

type UserService interface {
	RegisterUser(input web.RegisterUserInput) (domain.User, error)
	LoginUser(input web.LoginUserInput) (domain.User, error)
	CheckEmailAvailability(input web.CheckEmailInput) (bool, error)
	UploadAvatar(id int, fileLocation string) (domain.User, error)
	FindUserByID(id int) (domain.User, error)
	FindAllUsers() ([]domain.User, error)
}
