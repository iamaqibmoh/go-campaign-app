package main

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/auth"
	"bwa-campaign-app/controller"
	"bwa-campaign-app/middleware"
	"bwa-campaign-app/repository"
	"bwa-campaign-app/service"
)

func main() {
	//user endpoint dependency
	userRepository := repository.NewUserRepository(app.DBConnection())
	userService := service.NewUserService(userRepository)
	userAuth := auth.NewJWTAuth()
	userController := controller.NewUserController(userService, userAuth)

	//campaign endpoint dependency
	campaignRepository := repository.NewCampaignRepository(app.DBConnection())
	campaignService := service.NewCampaignService(campaignRepository)
	campaignController := controller.NewCampaignController(campaignService)

	router := app.Router()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	api.POST("/users", userController.RegisterUser)
	api.POST("/sessions", userController.LoginUser)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(userAuth, userService), userController.UploadAvatar)

	//campaigns endpoint
	api.GET("/campaigns", campaignController.FindCampaigns)
	api.POST("/campaigns", middleware.AuthMiddleware(userAuth, userService), campaignController.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(userAuth, userService), campaignController.UpdateCampaign)
	api.GET("/campaigns/:id", campaignController.FindCampaignByID)

	router.Run(":2802")
}
