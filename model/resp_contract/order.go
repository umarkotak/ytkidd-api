package resp_contract

type (
	CreateOrder struct {
		OrderNumber                string `json:"order_number"`
		MidtransSnapToken          string `json:"midtrans_snap_token"`
		MidtransSnapUrl            string `json:"midtrans_snap_url"`
		MidtransSandboxPaymentPage string `json:"midtrans_sandbox_payment_page"`
		MidtransNotificationPage   string `json:"midtrans_notification_page"`
	}
)
