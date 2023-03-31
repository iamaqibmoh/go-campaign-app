package controller

import (
	"bwa-campaign-app/auth"
	"bwa-campaign-app/formatter"
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	service.UserService
	auth.JWTAuth
}

func (c *UserControllerImpl) FetchUser(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(domain.User)
	ctx.JSON(200, helper.APIResponse(
		"Successfully fetch user data", 200, "success", formatter.UserResponseFormatter(user, "")))
}

func (c *UserControllerImpl) UploadAvatar(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to upload avatar",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"is_uploaded": false}))
		return
	}

	userID := ctx.MustGet("currentUser").(domain.User).ID
	path := fmt.Sprintf("images/%d-%s", userID, fileHeader.Filename)

	err = ctx.SaveUploadedFile(fileHeader, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to upload avatar",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"is_uploaded": false, "error": err.Error()}))
		return
	}

	_, err = c.UserService.UploadAvatar(userID, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to upload avatar",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"is_uploaded": false}))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"UpdateAvatar successfully uploaded",
		200,
		"success",
		gin.H{"is_uploaded": true},
	))
}

func (c *UserControllerImpl) CheckEmailAvailability(ctx *gin.Context) {
	input := web.CheckEmailInput{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Email checking failed",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"error": helper.ValidationErrorsFormatter(err)}))
		return
	}

	emailAvailability, err := c.UserService.CheckEmailAvailability(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Email checking failed",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"error": err.Error()}))
		return
	}

	var metaMessage string
	if emailAvailability {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	ctx.JSON(200, helper.APIResponse(
		metaMessage,
		200,
		"success",
		gin.H{"is_available": emailAvailability},
	))
}

func (c *UserControllerImpl) LoginUser(ctx *gin.Context) {
	input := web.LoginUserInput{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Login failed",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"error": helper.ValidationErrorsFormatter(err)}))
		return
	}

	loginUser, err := c.UserService.LoginUser(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Login failed",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"error": err.Error()}))
		return
	}

	token, err := c.JWTAuth.GenerateToken(loginUser.ID)
	helper.PanicIfError(err)

	ctx.JSON(200, helper.APIResponse(
		"You are logged in",
		200,
		"success",
		formatter.UserResponseFormatter(loginUser, token),
	))
}

func (c *UserControllerImpl) RegisterUser(ctx *gin.Context) {
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

	token, err := c.JWTAuth.GenerateToken(registerUser.ID)
	helper.PanicIfError(err)

	apiResponse := helper.APIResponse(
		"register user is successfully",
		200,
		"success",
		formatter.UserResponseFormatter(registerUser, token),
	)

	ctx.JSON(200, &apiResponse)
}

func NewUserController(userService service.UserService, auth auth.JWTAuth) UserController {
	return &UserControllerImpl{UserService: userService, JWTAuth: auth}
}
