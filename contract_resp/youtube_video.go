package contract_resp

type (
	YoutubeVideosHome struct {
		Videos     []YoutubeVideo `json:"videos"`
		ExcludeIDs []int64        `json:"exclude_ids"`
	}

	YoutubeVideo struct {
		ID         int64          `json:"id"`
		ImageUrl   string         `json:"image_url"`
		Title      string         `json:"title"`
		Channel    YoutubeChannel `json:"channel"`
		Tags       []string       `json:"tags"`
		ExternalID string         `json:"external_id,omitempty"`
		CanAction  bool           `json:"can_action"`
	}
)
