package test

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/repository"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newCampaignRepository() repository.CampaignRepository {
	db := app.DBConnection()
	return repository.NewCampaignRepository(db)
}

func TestSave(t *testing.T) {
	campaignRepository := newCampaignRepository()
	save, err := campaignRepository.Save(domain.Campaign{
		UserID:        2,
		Name:          "Campaign 3",
		Summary:       "Short Description",
		Description:   "Long Description",
		Perks:         "perks 1, perks 2, perks 3",
		BackerCount:   0,
		GoalAmount:    100000000,
		CurrentAmount: 0,
		Slug:          "campaign-tiga",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	})
	helper.PanicIfError(err)
	assert.Equal(t, "campaign-tiga", save.Slug)
}

func TestFindAll(t *testing.T) {
	campaignRepository := newCampaignRepository()
	campaigns, err := campaignRepository.FindAll()
	helper.PanicIfError(err)
	assert.Equal(t, 3, len(campaigns))
}

func TestPreloadCampImages(t *testing.T) {
	campaignRepository := newCampaignRepository()
	campaigns, err := campaignRepository.FindByUserID(1)
	helper.PanicIfError(err)
	bytes, _ := json.Marshal(campaigns)
	fmt.Println(string(bytes))
	//fmt.Println(campaigns[0].CampaignImages[0].FileName)
}

func TestFindByUserID(t *testing.T) {
	campaignRepository := newCampaignRepository()
	campaigns1, err := campaignRepository.FindByUserID(1)
	helper.PanicIfError(err)
	assert.Equal(t, 2, len(campaigns1))
	assert.Equal(t, "dua.jpg", campaigns1[0].CampaignImages[0].FileName)

	campaigns2, err := campaignRepository.FindByUserID(2)
	helper.PanicIfError(err)
	assert.Equal(t, 2, len(campaigns2))
	for _, campaign := range campaigns2 {
		assert.Equal(t, 0, len(campaign.CampaignImages))
	}
}
