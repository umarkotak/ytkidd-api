package resp_contract

type (
	YoutubeChannel struct {
		ID       int64    `json:"id"`
		ImageUrl string   `json:"image_url"`
		Name     string   `json:"name"`
		Tags     []string `json:"tags"`
	}
)
