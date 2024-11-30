package resp_contract

type (
	YoutubeVideosHome struct {
		Channels   []YoutubeVideosHomeChannel `json:"channels"`
		Videos     []YoutubeVideosHomeVideo   `json:"videos"`
		ExcludeIDs []int64                    `json:"exclude_ids"`
	}

	YoutubeVideosHomeVideo struct {
		ID       int64                    `json:"id"`
		ImageUrl string                   `json:"image_url"`
		Title    string                   `json:"title"`
		Channel  YoutubeVideosHomeChannel `json:"channel"`
		Tags     []string                 `json:"tags"`
	}

	YoutubeVideosHomeChannel struct {
		ID       int64    `json:"id"`
		ImageUrl string   `json:"image_url"`
		Name     string   `json:"name"`
		Tags     []string `json:"tags"`
	}
)
