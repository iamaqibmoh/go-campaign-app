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
	//users endpoint dependency
	userRepository := repository.NewUserRepository(app.DBConnection())
	userService := service.NewUserService(userRepository)
	userAuth := auth.NewJWTAuth()
	userController := controller.NewUserController(userService, userAuth)

	//campaigns endpoint dependency
	campaignRepository := repository.NewCampaignRepository(app.DBConnection())
	campaignService := service.NewCampaignService(campaignRepository)
	campaignController := controller.NewCampaignController(campaignService)

	//campaign images endpoint dependency
	campaignImageRepository := repository.NewCampaignImageRepository(app.DBConnection())
	campaignImageService := service.NewCampaignImageService(campaignImageRepository, campaignRepository)
	campaignImageController := controller.NewCampaignImageController(campaignImageService)

	//transactions dependency
	transactionsController := controller.NewTransactionController(
		service.NewTransactionsService(
			repository.NewTransactionsRepository(app.DBConnection()),
			campaignRepository),
	)

	router := app.Router()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	//users endpoint
	api.POST("/users", userController.RegisterUser)
	api.POST("/sessions", userController.LoginUser)
	api.POST("/email-checkers", userController.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(userAuth, userService), userController.UploadAvatar)

	//campaigns endpoint
	api.GET("/campaigns", campaignController.FindCampaigns)
	api.POST("/campaigns", middleware.AuthMiddleware(userAuth, userService), campaignController.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(userAuth, userService), campaignController.UpdateCampaign)
	api.GET("/campaigns/:id", campaignController.FindCampaignByID)

	//campaign images endpoint
	api.POST("/campaign-images", middleware.AuthMiddleware(userAuth, userService), campaignImageController.CreateCampaignImage)

	//transactions endpoint
	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(userAuth, userService), transactionsController.GetByCampaignID)

	router.Run(":2802")
}
