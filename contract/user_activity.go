package contract

import "github.com/umarkotak/ytkidd-api/model"

type (
	GetUserActivity struct {
		UserGuid       string `db:"-"`
		UserID         int64  `db:"user_id"`
		AppSession     string `db:"app_session"`
		YoutubeVideoId int64  `db:"youtube_video_id"`
		BookId         int64  `db:"book_id"`
		BookContentId  int64  `db:"book_content_id"`
		model.Pagination
	}

	RecordUserActivity struct {
		UserGuid             string                     `json:"-"`
		UserID               int64                      `json:"-"`
		AppSession           string                     `json:"-"`
		YoutubeVideoID       int64                      `json:"youtube_video_id"`
		BookID               int64                      `json:"book_id"`
		BookContentID        int64                      `json:"book_content_id"`
		UserActivityMetadata model.UserActivityMetadata `json:"metadata"`
	}
)
