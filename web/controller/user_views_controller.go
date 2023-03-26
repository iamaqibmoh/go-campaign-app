package controller

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type UserViewsController struct {
	service.UserService
}

func NewUserViewsController(userService service.UserService) *UserViewsController {
	return &UserViewsController{userService}
}

func (c *UserViewsController) PostUpdateAvatar(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	fileHeader, _ := ctx.FormFile("avatar")
	path := fmt.Sprintf("images/%d-%s", userID, fileHeader.Filename)
	_ = ctx.SaveUploadedFile(fileHeader, path)

	user, _ := c.UserService.FindUserByID(userID)
	if user.Avatar != "" {
		err := os.Remove(user.Avatar)
		helper.PanicIfError(err)
	}

	_, err := c.UserService.UploadAvatar(userID, path)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}
	ctx.Redirect(http.StatusFound, "/users")
}

func (c UserViewsController) UpdateAvatar(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	ctx.HTML(200, "user_avatar.gohtml", gin.H{"ID": userID})
}

func (c *UserViewsController) PostUpdate(ctx *gin.Context) {
	input := web.FormUpdateUserInput{}
	err := ctx.ShouldBind(&input)
	if err != nil {
		input.Error = err.Error()
		ctx.HTML(http.StatusBadRequest, "user_update.gohtml", input)
		return
	}
	userID, _ := strconv.Atoi(ctx.Param("id"))
	input.ID = userID

	_, err = c.UserService.UpdateUser(input)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/users")
}

func (c UserViewsController) Update(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.UserService.FindUserByID(userID)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	input := web.FormUpdateUserInput{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
	}

	ctx.HTML(200, "user_update.gohtml", input)
}

func (c *UserViewsController) PostCreate(ctx *gin.Context) {
	input := web.FormCreateUserInput{}
	err := ctx.ShouldBind(&input)
	if err != nil {
		input.Error = err.Error()
		ctx.HTML(200, "user_create.gohtml", input)
		return
	}

	register := web.RegisterUserInput{
		Name:       input.Name,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   input.Password,
	}
	_, err = c.UserService.RegisterUser(register)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/users")
}

func (c *UserViewsController) Create(ctx *gin.Context) {
	ctx.HTML(200, "user_create.gohtml", nil)
}

func (c *UserViewsController) Index(ctx *gin.Context) {
	users, err := c.UserService.FindAllUsers()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}
	ctx.HTML(200, "user_index.gohtml", gin.H{"users": users})
}
