package controller

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (c *UserControllerImpl) UserRegister(ctx *gin.Context) {
	input := web.RegisterUserInput{}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Register failed, your input is wrong",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"error": helper.ValidationErrorsFormatter(err)}))
		return
	}

	registerUser, err := c.UserService.RegisterUser(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Register failed, server is error",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"error": err.Error()}))
		return
	}

	apiResponse := helper.APIResponse(
		"register user is successfully",
		200,
		"success",
		helper.RegisterUserResponseFormatter(registerUser),
	)

	ctx.JSON(200, &apiResponse)
}
