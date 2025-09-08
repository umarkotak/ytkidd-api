package contract

import (
	"github.com/umarkotak/ytkidd-api/model"
)

type (
	GetUserStroke struct {
		UserGuid      string `json:"-"`
		AppSession    string `json:"-"`
		BookContentID int64  `json:"book_content_id"`
	}

	StoreUserStroke struct {
		UserGuid      string        `json:"-"`
		AppSession    string        `json:"-"`
		BookID        int64         `json:"book_id"`
		BookContentID int64         `json:"book_content_id"`
		ImageUrl      string        `json:"image_url"`
		Strokes       model.Strokes `json:"strokes"`
	}
)
