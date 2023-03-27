package controller

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type SessionCMSController struct {
	service.UserService
}

func (c *SessionCMSController) Destroy(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Delete("test")
	err := session.Save()
	helper.PanicIfError(err)

	ctx.Redirect(http.StatusFound, "/login")
}

func NewSessionCMSController(userService service.UserService) *SessionCMSController {
	return &SessionCMSController{UserService: userService}
}

func (c *SessionCMSController) NewSession(ctx *gin.Context) {
	ctx.HTML(200, "session_index.gohtml", nil)
}

func (c SessionCMSController) Create(ctx *gin.Context) {
	input := web.LoginUserInput{}
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.HTML(200, "session_index.gohtml", nil)
		return
	}

	user, err := c.UserService.LoginUser(input)
	if err != nil || user.Role != "admin" {
		ctx.HTML(200, "session_index.gohtml", nil)
		return
	}

	session := sessions.Default(ctx)
	userID, _ := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(user.ID)), bcrypt.DefaultCost)
	session.Set("test", userID)
	session.Save()

	ctx.Redirect(http.StatusFound, "/users")
}
