package middleware

import (
	"bwa-campaign-app/auth"
	"bwa-campaign-app/helper"
	"bwa-campaign-app/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(auth auth.JWTAuth, service service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Getting the header value
		getHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(getHeader, "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil))
			return
		}

		//Get token string from header value
		getHeaderSplit := strings.Split(getHeader, " ")
		token := getHeaderSplit[1]

		//Validate token string
		validateToken, err := auth.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				err.Error()))
			return
		}

		//Get payload/claim from token
		claims := validateToken.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		//Get user data based on userID from jwt payload
		findUserByID, err := service.FindUserByID(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.APIResponse(
				"InternalServerError",
				http.StatusInternalServerError,
				"error",
				err.Error()))
			return
		}

		//Set context which contain user data
		ctx.Set("currentUser", findUserByID)
	}
}
