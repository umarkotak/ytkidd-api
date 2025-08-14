package resp_contract

import (
	"github.com/umarkotak/ytkidd-api/model"
)

type (
	PublicProduct struct {
		Code        string                `json:"code"`
		BenefitType string                `json:"benefit_type"`
		Name        string                `json:"name"`
		ImageUrl    string                `json:"image_url"`
		BasePrice   int64                 `json:"base_price"`
		Price       int64                 `json:"price"`
		Metadata    model.ProductMetadata `json:"metadata"`
	}
)
