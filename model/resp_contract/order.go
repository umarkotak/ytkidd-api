package resp_contract

type (
	CreateOrder struct {
		OrderNumber       string `json:"order_number"`
		MidtransSnapToken string `json:"midtrans_snap_token"`
		MidtransSnapUrl   string `json:"midtrans_snap_url"`
	}
)
