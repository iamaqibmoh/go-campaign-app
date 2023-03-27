package app

import (
	"bwa-campaign-app/auth"
	"bwa-campaign-app/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	//Cors
	router.Use(cors.Default())

	//Cookie
	newStore := cookie.NewStore(auth.SECRET_KEY)
	router.Use(sessions.Sessions("test", newStore))

	//HTML Multi Templating
	router.HTMLRender = helper.LoadTemplates("./web/templates")

	//Static Files
	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	return router
}
