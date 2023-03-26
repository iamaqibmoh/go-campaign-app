package domain

import (
	"github.com/leekchan/accounting"
	"time"
)

type Campaign struct {
	ID             int
	UserID         int
	Name           string
	Summary        string
	Description    string
	Perks          string
	BackerCount    int
	GoalAmount     int
	CurrentAmount  int
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           User
	CampaignImages []CampaignImage
}

func (c Campaign) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
		Thousand:  ".",
		Decimal:   ",",
	}
	return ac.FormatMoney(c.CurrentAmount)
}

func (c Campaign) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
		Thousand:  ".",
		Decimal:   ",",
	}
	return ac.FormatMoney(c.GoalAmount)
}
