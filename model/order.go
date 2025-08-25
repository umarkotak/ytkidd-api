package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	ORDER_STATUS_INITIALIZED = "initialized" // v
	ORDER_STATUS_PENDING     = "pending"     // v
	ORDER_STATUS_PAID        = "paid"        // v
	ORDER_STATUS_COMPLETE    = "complete"    //
	ORDER_STATUS_CANCELED    = "canceled"    // v
	ORDER_STATUS_FAILED      = "failed"      // v
	ORDER_STATUS_EXPIRED     = "expired"     // v

	ORDER_NUMBER_FORMAT = "CB-%v"
)

var (
	ORDER_FINAL_STATES = []string{
		ORDER_STATUS_COMPLETE, ORDER_STATUS_CANCELED, ORDER_STATUS_FAILED, ORDER_STATUS_EXPIRED,
	}
)

type (
	Order struct {
		ID             int64          `db:"id"`
		CreatedAt      time.Time      `db:"created_at"`
		UpdatedAt      time.Time      `db:"updated_at"`
		DeletedAt      sql.NullTime   `db:"deleted_at"`
		UserID         int64          `db:"user_id"`
		Number         string         `db:"number"`
		OrderType      string         `db:"order_type"`
		Description    string         `db:"description"`
		Status         string         `db:"status"`
		PaymentStatus  string         `db:"payment_status"`
		BasePrice      int64          `db:"base_price"`
		Price          int64          `db:"price"`
		DiscountAmount int64          `db:"discount_amount"`
		FinalPrice     int64          `db:"final_price"`
		PaymentNumber  sql.NullString `db:"payment_number"`
		Metadata       OrderMetadata  `db:"metadata"`
	}

	OrderMetadata struct {
		ProductCode string `json:"product_code,omitempty"`
	}

	OrderWithPayment struct {
		ID             int64          `db:"id"`
		CreatedAt      time.Time      `db:"created_at"`
		UpdatedAt      time.Time      `db:"updated_at"`
		DeletedAt      sql.NullTime   `db:"deleted_at"`
		UserID         int64          `db:"user_id"`
		Number         string         `db:"number"`
		OrderType      string         `db:"order_type"`
		Description    string         `db:"description"`
		Status         string         `db:"status"`
		PaymentStatus  string         `db:"payment_status"`
		BasePrice      int64          `db:"base_price"`
		Price          int64          `db:"price"`
		DiscountAmount int64          `db:"discount_amount"`
		FinalPrice     int64          `db:"final_price"`
		PaymentNumber  sql.NullString `db:"payment_number"`
		Metadata       OrderMetadata  `db:"metadata"`

		PaymentExpiredAt       sql.NullTime    `db:"payment_expired_at"`
		PaymentSuccessAt       sql.NullTime    `db:"payment_success_at"`
		PaymentPaymentPlatform string          `db:"payment_payment_platform"`
		PaymentPaymentType     string          `db:"payment_payment_type"`
		PaymentReferenceStatus sql.NullString  `db:"payment_reference_status"`
		PaymentReferenceNumber sql.NullString  `db:"payment_reference_number"`
		PaymentFraudStatus     sql.NullString  `db:"payment_fraud_status"`
		PaymentMaskedCard      sql.NullString  `db:"payment_masked_card"`
		PaymentAmount          int64           `db:"payment_amount"`
		PaymentMetadata        PaymentMetadata `db:"payment_metadata"`
	}

	PaymentMetadata struct {
		SnapToken string `json:"snap_token"`
		SnapUrl   string `json:"snap_url"`
	}
)

func (m *Order) GenNumber() {
	m.Number = fmt.Sprintf(ORDER_NUMBER_FORMAT, time.Now().UnixNano())
}

func GenHumanStatus(status string) string {
	switch status {
	case ORDER_STATUS_INITIALIZED:
		return "Menunggu Pembayaran"
	case ORDER_STATUS_PENDING:
		return "Menunggu Pembayaran"
	case ORDER_STATUS_PAID:
		return "Menunggu Diproses"
	case ORDER_STATUS_COMPLETE:
		return "Selesai"
	case ORDER_STATUS_CANCELED:
		return "Dibatalkan"
	case ORDER_STATUS_FAILED:
		return "Gagal"
	case ORDER_STATUS_EXPIRED:
		return "Kadaluarsa"
	default:
		return "Unknown"
	}
}

func (m *Order) HumanStatus() string {
	return GenHumanStatus(m.Status)
}

func (m *OrderWithPayment) HumanStatus() string {
	return GenHumanStatus(m.Status)
}

func (m OrderMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *OrderMetadata) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

func (m PaymentMetadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *PaymentMetadata) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
