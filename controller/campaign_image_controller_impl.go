package controller

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CampaignImageControllerImpl struct {
	service.CampaignImageService
}

func (c *CampaignImageControllerImpl) CreateCampaignImage(ctx *gin.Context) {
	input := web.CreateCampaignImageInput{}
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest,
			"BAD REQUEST",
			helper.ValidationErrorsFormatter(err)))
		return
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"is_uploaded": false}))
		return
	}

	currentUser := ctx.MustGet("currentUser").(domain.User)
	path := fmt.Sprintf("images/%d-%s", currentUser.ID, fileHeader.Filename)
	input.User = currentUser

	err = ctx.SaveUploadedFile(fileHeader, path)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"is_uploaded": false}))
		return
	}

	_, err = c.CampaignImageService.CreateCampaignImage(input, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"is_uploaded": false, "errors": err.Error()}))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"Campaign image successfully uploaded",
		200,
		"success",
		gin.H{"is_uploaded": true},
	))
}

func NewCampaignImageController(campaignImageService service.CampaignImageService) CampaignImageController {
	return &CampaignImageControllerImpl{CampaignImageService: campaignImageService}
}
