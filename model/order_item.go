package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	ORDER_ITEM_TYPE_GIFT = "gift"
	ORDER_ITEM_TYPE_TIER = "tier"
)

type (
	OrderItem struct {
		ID        int64             `db:"id"`         //
		CreatedAt time.Time         `db:"created_at"` //
		UpdatedAt time.Time         `db:"updated_at"` //
		DeletedAt sql.NullTime      `db:"deleted_at"` //
		OrderId   int64             `db:"order_id"`   //
		ItemType  string            `db:"item_type"`  //
		ItemId    int64             `db:"item_id"`    //
		Price     int64             `db:"price"`      //
		Metadata  OrderItemMetadata `db:"metadata"`   //
	}

	OrderItemMetadata struct {
		ToUserID int64 `json:"to_user_id"`
	}
)

func (m OrderItemMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *OrderItemMetadata) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
