package main

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/controller"
	"bwa-campaign-app/repository"
	"bwa-campaign-app/service"
)

func main() {
	userRepository := repository.NewUserRepository(app.DBConnection())
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := app.Router()
	api := router.Group("/api/v1")
	api.POST("/users", userController.RegisterUser)
	api.POST("/sessions", userController.LoginUser)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", userController.UploadAvatar)

	router.Run(":2802")
}
