package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	BENEFIT_TYPE_SUBSCRIPTION = "subscription"
)

type (
	Product struct {
		ID          int64           `db:"id"`
		CreatedAt   time.Time       `db:"created_at"`
		UpdatedAt   time.Time       `db:"updated_at"`
		DeletedAt   sql.NullTime    `db:"deleted_at"`
		Code        string          `db:"code"`
		BenefitType string          `db:"benefit_type"`
		Name        string          `db:"name"`
		ImageUrl    string          `db:"image_url"`
		BasePrice   int64           `db:"base_price"`
		Price       int64           `db:"price"`
		Metadata    ProductMetadata `db:"metadata"`
	}

	ProductMetadata struct {
		DurationDays int64 `json:"duration_days,omitempty"`
	}
)

func (m ProductMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *ProductMetadata) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
