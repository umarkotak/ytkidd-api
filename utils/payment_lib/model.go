package payment_lib

import (
	"crypto/sha512"
	"database/sql"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"
)

type (
	Payment struct {
		ID              int64           `db:"id"`
		CreatedAt       time.Time       `db:"created_at"`
		UpdatedAt       time.Time       `db:"updated_at"`
		DeletedAt       sql.NullTime    `db:"deleted_at"`
		OrderNumber     string          `db:"order_number"`
		Number          string          `db:"number"`
		PaymentPlatform string          `db:"payment_platform"`
		PaymentType     string          `db:"payment_type"`
		Status          string          `db:"status"`
		ReferenceStatus sql.NullString  `db:"reference_status"`
		ReferenceNumber sql.NullString  `db:"reference_number"`
		FraudStatus     sql.NullString  `db:"fraud_status"`
		MaskedCard      sql.NullString  `db:"masked_card"`
		Amount          int64           `db:"amount"`
		Metadata        PaymentMetadata `db:"metadata"`
	}

	MidtransNotification struct {
		TransactionType   string         `json:"transaction_type"`
		TransactionTime   string         `json:"transaction_time"`
		TransactionStatus string         `json:"transaction_status"`
		TransactionID     string         `json:"transaction_id"`
		StatusMessage     string         `json:"status_message"`
		StatusCode        string         `json:"status_code"`
		SignatureKey      string         `json:"signature_key"`
		ReferenceID       string         `json:"reference_id"`
		PaymentType       string         `json:"payment_type"`
		OrderID           string         `json:"order_id"`
		Metadata          map[string]any `json:"metadata"`
		MerchantID        string         `json:"merchant_id"`
		GrossAmount       string         `json:"gross_amount"`
		FraudStatus       string         `json:"fraud_status"`
		ExpiryTime        string         `json:"expiry_time"`
		Currency          string         `json:"currency"`
		Acquirer          string         `json:"acquirer"`
	}

	PaymentMetadata struct{}
)

func (m *Payment) SyncStatus() {
	if m.PaymentPlatform == PAYMENT_PLATFORM_MIDTRANS {
		if slices.Contains(MIDTRANS_STATUS_TO_SUCCESS, m.ReferenceStatus.String) {
			m.Status = STATUS_SUCCESS
		}
		if slices.Contains(MIDTRANS_STATUS_TO_PENDING, m.ReferenceStatus.String) {
			m.Status = STATUS_PENDING
		}
		if slices.Contains(MIDTRANS_STATUS_TO_FAILED, m.ReferenceStatus.String) {
			m.Status = STATUS_FAILED
		}
		if slices.Contains(MIDTRANS_STATUS_TO_CANCELED, m.ReferenceStatus.String) {
			m.Status = STATUS_CANCELED
		}
		if slices.Contains(MIDTRANS_STATUS_TO_EXPIRED, m.ReferenceStatus.String) {
			m.Status = STATUS_EXPIRED
		}
		if slices.Contains(MIDTRANS_STATUS_TO_REFUNDED, m.ReferenceStatus.String) {
			m.Status = STATUS_REFUNDED
		}
		if m.FraudStatus.String == FRAUD_STATUS_DENY {
			m.Status = STATUS_FAILED
		}
	}
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

const SIGNATURE_FORMAT = "%s%s%s%s"

func (mc *MidtransNotification) ValidateSignature(serverKey string) error {
	signatureSeed := fmt.Sprintf(
		SIGNATURE_FORMAT,
		mc.OrderID, mc.StatusCode, mc.GrossAmount, serverKey,
	)

	hash := sha512.Sum512([]byte(signatureSeed))

	generatedSignature := hex.EncodeToString(hash[:])

	if generatedSignature != mc.SignatureKey {
		return fmt.Errorf("signature key missmatch")
	}

	return nil
}
