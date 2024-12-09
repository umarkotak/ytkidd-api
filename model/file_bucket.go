package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	FileBucket struct {
		ID              int64              `db:"id"`
		CreatedAt       time.Time          `db:"created_at"`
		UpdatedAt       time.Time          `db:"updated_at"`
		DeletedAt       sql.NullTime       `db:"deleted_at"`
		Guid            string             `db:"guid"`
		Name            string             `db:"name"`
		BaseType        string             `db:"base_type"`
		Extension       string             `db:"extension"`
		HttpContentType string             `db:"http_content_type"`
		Metadata        FileBucketMetadata `db:"metadata"`
		Data            []byte             `db:"data"`
		ExactPath       string             `db:"exact_path"`
	}

	FileBucketMetadata struct {
	}
)

func (m FileBucketMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *FileBucketMetadata) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
