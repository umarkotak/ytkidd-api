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
		ID        int64        `db:"id"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt time.Time    `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`

		UserID               int64                `db:"user_id"`
		AppSession           string               `db:"app_session"`
		YoutubeVideoID       int64                `db:"youtube_video_id"`
		BookID               int64                `db:"book_id"`
		BookContentID        int64                `db:"book_content_id"`
		UserActivityMetadata UserActivityMetadata `db:"metadata"`
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
