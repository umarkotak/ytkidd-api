package resp_contract

import (
	"time"

	"github.com/umarkotak/ytkidd-api/model"
)

type (
	CreateOrder struct {
		OrderNumber                string `json:"order_number"`
		MidtransSnapToken          string `json:"midtrans_snap_token"`
		MidtransSnapUrl            string `json:"midtrans_snap_url"`
		MidtransSandboxPaymentPage string `json:"midtrans_sandbox_payment_page"`
		MidtransNotificationPage   string `json:"midtrans_notification_page"`
	}

	OrderDetail struct {
		CreatedAt      time.Time           `json:"created_at"`
		UpdatedAt      time.Time           `json:"updated_at"`
		UserID         int64               `json:"user_id"`
		Number         string              `json:"number"`
		OrderType      string              `json:"order_type"`
		Description    string              `json:"description"`
		Status         string              `json:"status"`
		PaymentStatus  string              `json:"payment_status"`
		BasePrice      int64               `json:"base_price"`
		Price          int64               `json:"price"`
		DiscountAmount int64               `json:"discount_amount"`
		FinalPrice     int64               `json:"final_price"`
		PaymentNumber  string              `json:"payment_number"`
		Metadata       model.OrderMetadata `json:"metadata"`
	}

	CheckOrderPayment struct {
		OrderNumber   string `json:"order_number"`
		Status        string `json:"status"`
		PaymentStatus string `json:"payment_status"`
	}

	OrderList struct {
		Orders []OrderListData `json:"orders"`
	}

	OrderListData struct {
		CreatedAt      time.Time           `json:"created_at"`
		UpdatedAt      time.Time           `json:"updated_at"`
		UserID         int64               `json:"user_id"`
		Number         string              `json:"number"`
		OrderType      string              `json:"order_type"`
		Description    string              `json:"description"`
		Status         string              `json:"status"`
		PaymentStatus  string              `json:"payment_status"`
		BasePrice      int64               `json:"base_price"`
		Price          int64               `json:"price"`
		DiscountAmount int64               `json:"discount_amount"`
		FinalPrice     int64               `json:"final_price"`
		PaymentNumber  string              `json:"payment_number"`
		Metadata       model.OrderMetadata `json:"metadata"`
	}
)
