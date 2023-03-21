package web

import "time"

type CampaignTransactionResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionResponse struct {
	ID        int                       `json:"id"`
	Amount    int                       `json:"amount"`
	Status    string                    `json:"status"`
	CreatedAt time.Time                 `json:"created_at"`
	Campaign  CampaignOfUserTransaction `json:"campaign"`
}

type CampaignOfUserTransaction struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}
