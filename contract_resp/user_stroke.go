package contract_resp

import (
	"database/sql"

	"github.com/umarkotak/ytkidd-api/model"
)

type (
	GetUserStroke struct {
		ID            int64         `db:"id"`
		BookID        sql.NullInt64 `db:"book_id"`
		BookContentID sql.NullInt64 `db:"book_content_id"`
		ImageUrl      string        `db:"image_url"`
		Strokes       model.Strokes `db:"strokes"`
	}
)
