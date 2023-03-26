package controller

import (
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CampaignCMSController struct {
	service.CampaignService
	service.CampaignImageService
	service.UserService
}

func (c *CampaignCMSController) ShowDetail(ctx *gin.Context) {
	campID, _ := strconv.Atoi(ctx.Param("id"))
	campaign, err := c.CampaignService.FindDetailCampaignByID(web.CampaignIDFromURI{ID: campID})
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	ctx.HTML(200, "campaign_show.gohtml", campaign)
}

func (c *CampaignCMSController) PostUpdate(ctx *gin.Context) {
	campID, _ := strconv.Atoi(ctx.Param("id"))
	input := web.UpdateCampaignCMSInput{}
	err := ctx.ShouldBind(&input)
	if err != nil {
		input.ID = campID
		input.Error = err.Error()
		ctx.HTML(http.StatusInternalServerError, "campaign_update.gohtml", input)
		return
	}

	campaign, err := c.CampaignService.FindDetailCampaignByID(web.CampaignIDFromURI{ID: campID})
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}
	//campaign.Name = input.Name
	//campaign.Summary = input.ShortDescription
	//campaign.Description = input.Description
	//campaign.GoalAmount = input.GoalAmount
	//campaign.Perks = input.Perks
	user, err := c.UserService.FindUserByID(campaign.UserID)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	_, err = c.CampaignService.UpdateCampaign(campaign.ID, web.CreateCampaignInput{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		User:             user,
	})
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/campaigns")

}

func (c *CampaignCMSController) Update(ctx *gin.Context) {
	campID, _ := strconv.Atoi(ctx.Param("id"))
	campaign, err := c.CampaignService.FindDetailCampaignByID(web.CampaignIDFromURI{ID: campID})
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	input := web.UpdateCampaignCMSInput{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.Summary,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		Perks:            campaign.Perks,
	}

	ctx.HTML(200, "campaign_update.gohtml", input)
}

func (c *CampaignCMSController) PostUploadImage(ctx *gin.Context) {
	campID, _ := strconv.Atoi(ctx.Param("id"))

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	campaign, err := c.CampaignService.FindDetailCampaignByID(web.CampaignIDFromURI{ID: campID})
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	path := fmt.Sprintf("images/campaign-images/%d-%s", campaign.UserID, fileHeader.Filename)

	err = ctx.SaveUploadedFile(fileHeader, path)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	_, err = c.CampaignImageService.CreateCampaignImage(web.CreateCampaignImageInput{
		CampaignID: campID,
		IsPrimary:  true,
		User:       campaign.User,
	}, path)

	ctx.Redirect(http.StatusFound, "/campaigns")

}

func (c *CampaignCMSController) UploadImage(ctx *gin.Context) {
	campID, _ := strconv.Atoi(ctx.Param("id"))
	ctx.HTML(200, "campaign_image.gohtml", gin.H{"ID": campID})
}

func (c *CampaignCMSController) PostCreate(ctx *gin.Context) {
	input := web.CreateCampaignCMSInput{}
	err := ctx.ShouldBind(&input)
	if err != nil {
		users, e := c.UserService.FindAllUsers()
		if e != nil {
			ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
			return
		}
		input.Users = users
		input.Error = err.Error()

		ctx.HTML(http.StatusInternalServerError, "campaign_create.gohtml", input)
		return
	}

	user, err := c.UserService.FindUserByID(input.UserID)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	_, err = c.CampaignService.CreateCampaign(web.CreateCampaignInput{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		User:             user,
	})

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/campaigns")

}

func (c *CampaignCMSController) Create(ctx *gin.Context) {
	users, err := c.UserService.FindAllUsers()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}
	input := web.CreateCampaignCMSInput{Users: users}
	ctx.HTML(200, "campaign_create.gohtml", input)
}

func (c *CampaignCMSController) Index(ctx *gin.Context) {
	campaigns, err := c.CampaignService.FindCampaigns(0)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	ctx.HTML(200, "campaign_index.gohtml", gin.H{"campaigns": campaigns})
}

func NewCampaignCMSController(campaignService service.CampaignService, campaignImageService service.CampaignImageService, userService service.UserService) *CampaignCMSController {
	return &CampaignCMSController{CampaignService: campaignService, CampaignImageService: campaignImageService, UserService: userService}
}
