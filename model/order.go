package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

const (
	ORDER_STATUS_INITIALIZED       = "initialized"       //
	ORDER_STATUS_PAYMENT_PENDING   = "payment_pending"   //
	ORDER_STATUS_PAYMENT_COMPLETED = "payment_completed" //
	ORDER_STATUS_EXPIRED           = "expired"           // final state
	ORDER_STATUS_COMPLETED         = "completed"         // final state

	CLEARANCE_NOT_CHECKED = "not_checked"
	CLEARANCE_VERIFIED    = "verified"
	CLEARANCE_REJECTED    = "rejected"

	ORDER_TYPE_GIFT = "gift"
	ORDER_TYPE_TIER = "tier"
)

var (
	VALID_ORDER_STATUSES = []string{
		ORDER_STATUS_INITIALIZED,
	}
	FINAL_ORDER_STATUSES = []string{
		ORDER_STATUS_EXPIRED,
		ORDER_STATUS_COMPLETED,
	}
)

type (
	Order struct {
		ID              int64           `db:"id"`                //
		CreatedAt       time.Time       `db:"created_at"`        //
		UpdatedAt       time.Time       `db:"updated_at"`        //
		DeletedAt       sql.NullTime    `db:"deleted_at"`        //
		UserID          int64           `db:"user_id"`           //
		Number          string          `db:"number"`            //
		OrderType       string          `db:"order_type"`        //
		GrandTotalPrice decimal.Decimal `db:"grand_total_price"` //
		Status          string          `db:"status"`            //
		PaymentPlatform string          `db:"payment_platform"`  //
		PaymentType     string          `db:"payment_type"`      //
		PaymentStatus   string          `db:"payment_status"`    //
		PaymentNumber   string          `db:"payment_number"`    //
		PaymentNotes    string          `db:"payment_notes"`     //
		Metadata        OrderMetadata   `db:"metadata"`          //
		Clearance       string          `db:"clearance"`         //
	}

	OrderMetadata struct {
	}
)

func (m OrderMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *OrderMetadata) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
