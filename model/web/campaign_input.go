package web

type CampaignIDFromURI struct {
	ID int `uri:"id" binding:"required"`
}
