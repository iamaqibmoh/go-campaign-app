package web

type CampaignDetailResponse struct {
	ID               int                              `json:"id"`
	Name             string                           `json:"name"`
	UserID           int                              `json:"user_id"`
	ShortDescription string                           `json:"short_description"`
	Description      string                           `json:"description"`
	ImageURL         string                           `json:"image_url"`
	BackerCount      int                              `json:"backer_count"`
	GoalAmount       int                              `json:"goal_amount"`
	CurrentAmount    int                              `json:"current_amount"`
	Slug             string                           `json:"slug"`
	Perks            []string                         `json:"perks"`
	User             UserOfCampaignDetail             `json:"user"`
	Images           []CampaignImagesOfCampaignDetail `json:"images"`
}

type UserOfCampaignDetail struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImagesOfCampaignDetail struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type CampaignResponse struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}
