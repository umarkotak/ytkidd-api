package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type (
	YoutubeChannel struct {
		ID          int64          `db:"id"`
		CreatedAt   time.Time      `db:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at"`
		DeletedAt   sql.NullTime   `db:"deleted_at"`
		ExternalID  string         `db:"external_id"`
		Name        string         `db:"name"`
		Username    string         `db:"username"`
		ImageUrl    string         `db:"image_url"`
		Tags        pq.StringArray `db:"tags"`
		Active      bool           `db:"active"`
		ChannelLink string         `db:"channel_link"`
	}
)
