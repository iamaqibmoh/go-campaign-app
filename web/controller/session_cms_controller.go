package controller

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type SessionCMSController struct {
	service.UserService
}

func (c *SessionCMSController) Destroy(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Delete("UAccess")

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
	valueSession := uuid.NewString()
	session.Set("UAccess", valueSession)
	session.Save()

	ctx.Redirect(http.StatusFound, "/users")
}
