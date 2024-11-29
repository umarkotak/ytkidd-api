package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	PaymentCallback struct {
		ID              int64                   `db:"id"`               //
		CreatedAt       time.Time               `db:"created_at"`       //
		UpdatedAt       time.Time               `db:"updated_at"`       //
		DeletedAt       sql.NullTime            `db:"deleted_at"`       //
		PaymentPlatform string                  `db:"payment_platform"` //
		OrderNumber     string                  `db:"order_number"`     //
		Metadata        PaymentCallbackMetadata `db:"metadata"`         //
	}

	PaymentCallbackMetadata struct {
	}
)

func (m PaymentCallbackMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *PaymentCallbackMetadata) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
