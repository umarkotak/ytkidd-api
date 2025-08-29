package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	UserStroke struct {
		ID           int64          `db:"id" json:"id"`
		CreatedAt    time.Time      `db:"created_at" json:"created_at"`
		UpdatedAt    time.Time      `db:"updated_at" json:"updated_at"`
		DeletedAt    sql.NullTime   `db:"deleted_at" json:"deleted_at"`
		BookID       sql.NullString `db:"book_id" json:"book_id"`
		UserID       sql.NullString `db:"user_id" json:"user_id"`
		AppSession   sql.NullString `db:"app_session" json:"app_session"`
		ImageUrl     string         `db:"image_url" json:"image_url"`
		Tool         string         `db:"tool" json:"tool"`
		Color        string         `db:"color" json:"color"`
		RelativeSize float64        `db:"relative_size" json:"relative_size"`
		Opacity      float64        `db:"opacity" json:"opacity"`
		Points       StrokePoints   `db:"points" json:"points"`
	}

	StrokePoints []struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}
)

func (m StrokePoints) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *StrokePoints) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
