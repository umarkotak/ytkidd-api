package contract

import (
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/utils"
)

type (
	OrderCreate struct {
		UserGuid string `json:"-"`
		ItemType string `json:"item_type"` // Enum: tier, gift
		ItemSku  string `json:"item_sku"`  //

		ToUserGuid string `json:"to_user_guid"` // which user you want to send the gift to
	}

	OrderPay struct {
		UserGuid     string `json:"-"`
		OrderNumber  string `json:"-"`
		PaymentToken string `json:"payment_token"`
	}

	OrderBenefit struct {
		UserGuid    string `json:"-"`
		OrderNumber string `json:"-"`
	}
)

func (m *OrderCreate) Validate() error {
	if !utils.SliceStringContain(
		[]string{model.ORDER_ITEM_TYPE_GIFT, model.ORDER_ITEM_TYPE_TIER}, m.ItemType,
	) {
		return model.ErrBadRequest
	}

	return nil
}

func (m *OrderPay) Validate() error {

	return nil
}
