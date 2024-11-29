package contract

type (
	ScrapVideos struct {
		ChannelID string `json:"channel_id"`
		PageToken string `json:"page_token"`
		Query     string `json:"query"`

		All bool `json:"all"`
	}
)
