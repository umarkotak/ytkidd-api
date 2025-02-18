package contract

type (
	ScrapVideos struct {
		ChannelID     string `json:"channel_id"`
		PageToken     string `json:"page_token"`
		Query         string `json:"query"`
		MaxPage       int64  `json:"max_page"`
		BreakOnExists bool   `json:"break_on_exists"`
	}
)
