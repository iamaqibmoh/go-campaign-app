package test

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/service"
	"fmt"
	"testing"
)

func newCampaignService() service.CampaignService {
	campaignRepository := newCampaignRepository()
	return service.NewCampaignService(campaignRepository)
}

func TestFindCampaigns(t *testing.T) {
	campaignService := newCampaignService()
	campaigns, err := campaignService.FindCampaigns(3)
	helper.PanicIfError(err)

	fmt.Println(len(campaigns))
}
