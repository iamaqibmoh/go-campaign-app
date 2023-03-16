package test

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
	"bwa-campaign-app/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newUserService() service.UserService {
	userRepository := repository.NewUserRepository(app.DBConnection())
	userService := service.NewUserService(userRepository)

	return userService
}

func TestUserServiceRegisterUser(t *testing.T) {
	userService := newUserService()
	registerUser, _ := userService.RegisterUser(web.RegisterUserInput{
		Name:       "Otong",
		Occupation: "Backend Dev",
		Email:      "otong@test.com",
		Password:   "123",
	})

	assert.Equal(t, "otong@test.com", registerUser.Email)
}
