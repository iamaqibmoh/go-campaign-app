package test

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/helper"
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

func TestUserLogin(t *testing.T) {
	userService := newUserService()
	input := web.LoginUserInput{
		Email:    "mario@test.com",
		Password: "123",
	}
	loginUser, err := userService.LoginUser(input)
	helper.PanicIfError(err)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	assert.Equal(t, "mario@test.com", loginUser.Email)
	assert.Equal(t, "FrontEnd", loginUser.Occupation)
}

func TestCheckEmail(t *testing.T) {
	userService := newUserService()
	emailAvailability, err := userService.CheckEmailAvailability(web.CheckEmailInput{Email: "mario@test.com"})
	helper.PanicIfError(err)

	assert.False(t, emailAvailability)
}

func TestUploadAvatar(t *testing.T) {
	userService := newUserService()
	uploadAvatar, err := userService.UploadAvatar(1, "images/1-avatar.jpg")
	helper.PanicIfError(err)

	assert.Equal(t, "images/1-avatar.jpg", uploadAvatar.Avatar)
}
