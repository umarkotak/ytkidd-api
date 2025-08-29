package model

import (
	"database/sql"
	"time"
)

type (
	UserActivity struct {
		ID        int64        `db:"id" json:"id"`
		CreatedAt time.Time    `db:"created_at" json:"created_at"`
		UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`

		UserID         sql.NullString `db:"user_id" json:"user_id"`
		AppSession     sql.NullString `db:"app_session" json:"app_session"`
		YoutubeVideoID sql.NullString `db:"youtube_video_id" json:"youtube_video_id"`
		BookID         sql.NullString `db:"book_id" json:"book_id"`
		BookContentID  sql.NullString `db:"book_content_id" json:"book_content_id"`
	}
)
