package main

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/auth"
	"bwa-campaign-app/controller"
	"bwa-campaign-app/middleware"
	"bwa-campaign-app/repository"
	"bwa-campaign-app/service"
	viewsController "bwa-campaign-app/web/controller"
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
	transactionsRepository := repository.NewTransactionsRepository(app.DBConnection())
	midtransService := service.NewMidtransService(transactionsRepository, campaignRepository)
	transactionsService := service.NewTransactionsService(transactionsRepository, campaignRepository, midtransService)
	transactionsController := controller.NewTransactionsController(transactionsService, midtransService)

	//cms
	userViewsController := viewsController.NewUserViewsController(userService)
	campaignCMSController := viewsController.NewCampaignCMSController(campaignService, campaignImageService, userService)

	router := app.Router()
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
	api.GET("/transactions", middleware.AuthMiddleware(userAuth, userService), transactionsController.GetByUserID)
	api.POST("/transactions", middleware.AuthMiddleware(userAuth, userService), transactionsController.CreateTransaction)
	api.POST("/transactions/notification", transactionsController.GetMidtransNotification)

	//user cms controller
	router.GET("/users", userViewsController.Index)
	router.GET("/users/new", userViewsController.Create)
	router.POST("/users", userViewsController.PostCreate)
	router.GET("/users/edit/:id", userViewsController.Update)
	router.POST("/users/update/:id", userViewsController.PostUpdate)
	router.GET("/users/avatar/:id", userViewsController.UpdateAvatar)
	router.POST("/users/avatar/:id", userViewsController.PostUpdateAvatar)

	//campaign cms controller
	router.GET("/campaigns", campaignCMSController.Index)
	router.GET("/campaigns/show/:id", campaignCMSController.ShowDetail)
	router.GET("/campaigns/new", campaignCMSController.Create)
	router.POST("/campaigns", campaignCMSController.PostCreate)
	router.GET("/campaigns/image/:id", campaignCMSController.UploadImage)
	router.POST("/campaigns/image/:id", campaignCMSController.PostUploadImage)
	router.GET("/campaigns/edit/:id", campaignCMSController.Update)
	router.POST("/campaigns/update/:id", campaignCMSController.PostUpdate)

	router.Run(":2802")
}
