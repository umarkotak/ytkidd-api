package resp_contract

type (
	TierPlanData struct {
		ID              int64       `db:"id"`               //
		Sku             string      `db:"sku"`              //
		TitleId         string      `db:"title_id"`         //
		TitleEn         string      `db:"title_en"`         //
		ImageUrl        string      `db:"image_url"`        //
		Duration        int64       `db:"duration"`         //
		Price           int64       `db:"price"`            //
		PaymentPlatform string      `db:"payment_platform"` // Enum: free, playstore, appstore
		Active          bool        `db:"active"`           //
		Priority        int64       `db:"priority"`         //
		Codename        string      `db:"codename"`         //
		Benefit         TierBenefit `db:"benefit"`          //
	}

	TierBenefit struct {
		BrowsingLimit  int64 `json:"browsing_limit"`   //
		AllowSearch    bool  `json:"allow_search"`     //
		AllowVoiceCall bool  `json:"allow_voice_call"` //
		AllowVideoCall bool  `json:"allow_video_call"` //
	}
)
