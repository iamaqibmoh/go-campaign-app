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

	//campaign image endpoint dependency
	campaignImageRepository := repository.NewCampaignImageRepository(app.DBConnection())
	campaignImageService := service.NewCampaignImageService(campaignImageRepository, campaignRepository)
	campaignImageController := controller.NewCampaignImageController(campaignImageService)

	router := app.Router()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	//user endpoint
	api.POST("/users", userController.RegisterUser)
	api.POST("/sessions", userController.LoginUser)
	api.POST("/email-checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(userAuth, userService), userController.UploadAvatar)

	//campaigns endpoint
	api.GET("/campaigns", campaignController.FindCampaigns)
	api.POST("/campaigns", middleware.AuthMiddleware(userAuth, userService), campaignController.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(userAuth, userService), campaignController.UpdateCampaign)
	api.GET("/campaigns/:id", campaignController.FindCampaignByID)

	//campaign image endpoint
	api.POST("/campaign-images", middleware.AuthMiddleware(userAuth, userService), campaignImageController.CreateCampaignImage)

	router.Run(":2802")
}
