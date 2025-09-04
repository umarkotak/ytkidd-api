package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	STORAGE_LOCAL = "local"
	STORAGE_R2    = "r2"

	PURPOSE_BOOK_CONTENT = "book_content"
	PURPOSE_BOOK_COVER   = "book_cover"
	PURPOSE_BOOK_PDF     = "book_pdf"
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
		Storage         string             `db:"storage"`
		SizeKb          int64              `db:"size_kb"`
	}

	FileBucketMetadata struct {
		Purpose string `json:"purpose"` // Enum: book_content, book_cover, book_pdf
	}
)

func (m FileBucketMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *FileBucketMetadata) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
