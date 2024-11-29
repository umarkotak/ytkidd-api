package resp_contract

type (
	GiftPlanData struct {
		Sku             string `json:"sku"`              //
		TitleId         string `json:"title_id"`         //
		TitleEn         string `json:"title_en"`         //
		ImageUrl        string `json:"image_url"`        //
		Price           int64  `json:"price"`            //
		PaymentPlatform string `json:"payment_platform"` // Enum: playstore, appstore
		Codename        string `json:"codename"`         //
		Active          bool   `json:"active"`           //
	}
)
