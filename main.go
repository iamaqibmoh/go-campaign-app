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
	userCMSController := viewsController.NewUserCMSController(userService)
	campaignCMSController := viewsController.NewCampaignCMSController(campaignService, campaignImageService, userService)
	transactionCMSController := viewsController.NewTransactionCMSController(transactionsService)
	sessionCMSController := viewsController.NewSessionCMSController(userService)

	router := app.Router()
	api := router.Group("/api/v1")

	//users endpoint
	api.POST("/users", userController.RegisterUser)
	api.GET("/users/fetch", middleware.AuthMiddleware(userAuth, userService), userController.FetchUser)
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

	//users cms controller
	router.GET("/users", middleware.WebCMSAuthMiddleware(), userCMSController.Index)
	router.GET("/users/new", middleware.WebCMSAuthMiddleware(), userCMSController.Create)
	router.POST("/users", middleware.WebCMSAuthMiddleware(), userCMSController.PostCreate)
	router.GET("/users/edit/:id", middleware.WebCMSAuthMiddleware(), userCMSController.Update)
	router.POST("/users/update/:id", middleware.WebCMSAuthMiddleware(), userCMSController.PostUpdate)
	router.GET("/users/avatar/:id", middleware.WebCMSAuthMiddleware(), userCMSController.UpdateAvatar)
	router.POST("/users/avatar/:id", middleware.WebCMSAuthMiddleware(), userCMSController.PostUpdateAvatar)

	//campaigns cms controller
	router.GET("/campaigns", middleware.WebCMSAuthMiddleware(), campaignCMSController.Index)
	router.GET("/campaigns/show/:id", middleware.WebCMSAuthMiddleware(), campaignCMSController.ShowDetail)
	router.GET("/campaigns/new", middleware.WebCMSAuthMiddleware(), campaignCMSController.Create)
	router.POST("/campaigns", middleware.WebCMSAuthMiddleware(), campaignCMSController.PostCreate)
	router.GET("/campaigns/image/:id", middleware.WebCMSAuthMiddleware(), campaignCMSController.UploadImage)
	router.POST("/campaigns/image/:id", middleware.WebCMSAuthMiddleware(), campaignCMSController.PostUploadImage)
	router.GET("/campaigns/edit/:id", middleware.WebCMSAuthMiddleware(), campaignCMSController.Update)
	router.POST("/campaigns/update/:id", middleware.WebCMSAuthMiddleware(), campaignCMSController.PostUpdate)

	//transactions cms controller
	router.GET("/transactions", middleware.WebCMSAuthMiddleware(), transactionCMSController.ShowAll)

	//sessions cms controller
	router.GET("/login", sessionCMSController.NewSession)
	router.POST("/sessions", sessionCMSController.Create)
	router.GET("/logout", middleware.WebCMSAuthMiddleware(), sessionCMSController.Destroy)

	router.Run(":2802")
}
