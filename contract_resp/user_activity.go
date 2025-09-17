package contract_resp

import "github.com/umarkotak/ytkidd-api/model"

type (
	GetUserActivity struct {
		Activities []UserActivitySimple `json:"activities"`
	}

	UserActivitySimple struct {
		ActivityType         string                     `json:"activity_type"`    // Enum: book, video
		YoutubeVideoID       int64                      `json:"youtube_video_id"` //
		BookID               int64                      `json:"book_id"`          //
		BookContentID        int64                      `json:"book_content_id"`  //
		UserActivityMetadata model.UserActivityMetadata `json:"metadata"`         //
		Book                 UserActivityBook           `json:"book"`             //
		Video                UserActivityVideo          `json:"video"`            //
	}

	UserActivityBook struct {
		Title        string `json:"title"`
		ImageUrl     string `json:"image_url"`
		RedirectPath string `json:"redirect_path"`
	}

	UserActivityVideo struct {
		Title           string `json:"title"`
		ImageUrl        string `json:"image_url"`
		RedirectPath    string `json:"redirect_path"`
		ChannelName     string `json:"channel_name"`
		ChannelImageUrl string `json:"channel_image_url"`
	}
)
