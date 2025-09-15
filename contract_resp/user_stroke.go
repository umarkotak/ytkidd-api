package contract_resp

import (
	"github.com/umarkotak/ytkidd-api/model"
)

type (
	GetUserStroke struct {
		ID            int64         `json:"id"`
		BookID        int64         `json:"book_id"`
		BookContentID int64         `json:"book_content_id"`
		ImageUrl      string        `json:"image_url"`
		Strokes       model.Strokes `json:"strokes"`
	}
)
