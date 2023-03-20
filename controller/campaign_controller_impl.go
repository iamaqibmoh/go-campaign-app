package controller

import (
	"bwa-campaign-app/formatter"
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CampaignControllerImpl struct {
	service.CampaignService
}

func (c *CampaignControllerImpl) UpdateCampaign(ctx *gin.Context) {
	inputURI := web.CampaignIDFromURI{}
	err := ctx.ShouldBindUri(&inputURI)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to update campaign",
			http.StatusBadRequest,
			"BAD REQUEST",
			helper.ValidationErrorsFormatter(err)))
		return
	}

	inputForm := web.CreateCampaignInput{}
	err = ctx.ShouldBindJSON(&inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to update campaign",
			http.StatusBadRequest,
			"BAD REQUEST",
			helper.ValidationErrorsFormatter(err)))
		return
	}

	inputForm.User = ctx.MustGet("currentUser").(domain.User)

	campaign, err := c.CampaignService.UpdateCampaign(inputURI.ID, inputForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to update campaign",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			err.Error()))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"Campaign is successfully updated",
		200, "success",
		formatter.CampaignFormatter(campaign)))
}

func (c *CampaignControllerImpl) CreateCampaign(ctx *gin.Context) {
	input := web.CreateCampaignInput{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to create campaign",
			http.StatusBadRequest,
			"BAD REQUEST",
			err.Error()))
		return
	}

	user := ctx.MustGet("currentUser").(domain.User)
	input.User = user

	campaign, err := c.CampaignService.CreateCampaign(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to create campaign",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			err.Error()))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"Campaign is successfully created",
		200, "success",
		formatter.CampaignFormatter(campaign)))
}

func (c *CampaignControllerImpl) FindCampaignByID(ctx *gin.Context) {
	input := web.CampaignIDFromURI{}
	err := ctx.ShouldBindUri(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to get detail of campaign",
			http.StatusBadRequest,
			"BAD REQUEST",
			err.Error()))
		return
	}

	campaign, err := c.CampaignService.FindDetailCampaignByID(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to get detail of campaign",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			err.Error()))
		return
	}

	if campaign.ID == 0 {
		ctx.JSON(404, helper.APIResponse(
			"Campaign not found",
			404,
			"NOT FOUND",
			nil))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"Detail of campaign",
		200,
		"success",
		formatter.CampaignDetailFormatter(campaign)))
}

func (c *CampaignControllerImpl) FindCampaigns(ctx *gin.Context) {
	value := ctx.Query("user_id")
	userID, _ := strconv.Atoi(value)

	campaigns, err := c.CampaignService.FindCampaigns(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to get campaigns",
			500,
			"INTERNAL SERVER ERROR",
			nil))
		return
	}

	if len(campaigns) == 0 {
		ctx.JSON(http.StatusNotFound, helper.APIResponse(
			"Failed to get campaigns",
			404,
			"NOT FOUND",
			campaigns))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"List of campaigns",
		200,
		"success",
		formatter.CampaignsFormatter(campaigns)))
}

func NewCampaignController(campaignService service.CampaignService) CampaignController {
	return &CampaignControllerImpl{CampaignService: campaignService}
}
