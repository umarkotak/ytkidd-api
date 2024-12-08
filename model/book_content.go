package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	BookContent struct {
		ID          int64               `db:"id"`
		CreatedAt   time.Time           `db:"created_at"`
		UpdatedAt   time.Time           `db:"updated_at"`
		DeletedAt   sql.NullTime        `db:"deleted_at"`
		BookID      int64               `db:"book_id"`
		ImageUrl    string              `db:"image_url"`
		Description string              `db:"description"`
		Metadata    BookContentMetadata `db:"metadata"`
	}

	BookContentMetadata struct {
	}
)

func (m BookContentMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *BookContentMetadata) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
