package controller

import (
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserViewsController struct {
	service.UserService
}

func NewUserViewsController(userService service.UserService) *UserViewsController {
	return &UserViewsController{userService}
}

func (c *UserViewsController) PostCreate(ctx *gin.Context) {
	input := web.FormCreateUserInput{}
	err := ctx.ShouldBind(&input)
	if err != nil {
		input.Error = err
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
