package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	UserActivity struct {
		ID        int64        `db:"id" json:"id"`
		CreatedAt time.Time    `db:"created_at" json:"created_at"`
		UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`

		UserID               sql.NullString       `db:"user_id" json:"user_id"`
		AppSession           sql.NullString       `db:"app_session" json:"app_session"`
		YoutubeVideoID       sql.NullString       `db:"youtube_video_id" json:"youtube_video_id"`
		BookID               sql.NullString       `db:"book_id" json:"book_id"`
		BookContentID        sql.NullString       `db:"book_content_id" json:"book_content_id"`
		UserActivityMetadata UserActivityMetadata `db:"metadata" json:"metadata"`
	}

	UserActivityMetadata struct {
	}
)

func (m UserActivityMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *UserActivityMetadata) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
