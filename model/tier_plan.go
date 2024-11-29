package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	TIER_FREE   = "free"
	TIER_SILVER = "silver"
	TIER_GOLD   = "gold"
)

var (
	NEXT_TIER_MAP = map[string]string{
		TIER_FREE:   TIER_SILVER,
		TIER_SILVER: TIER_GOLD,
		TIER_GOLD:   TIER_GOLD,
	}
)

type (
	TierPlan struct {
		ID              int64        `db:"id"`               //
		CreatedAt       time.Time    `db:"created_at"`       //
		UpdatedAt       time.Time    `db:"updated_at"`       //
		DeletedAt       sql.NullTime `db:"deleted_at"`       //
		Sku             string       `db:"sku"`              //
		TitleId         string       `db:"title_id"`         //
		TitleEn         string       `db:"title_en"`         //
		ImageUrl        string       `db:"image_url"`        //
		Duration        int64        `db:"duration"`         //
		Price           int64        `db:"price"`            //
		PaymentPlatform string       `db:"payment_platform"` // Enum: free, playstore, appstore
		Active          bool         `db:"active"`           //
		Priority        int64        `db:"priority"`         //
		Codename        string       `db:"codename"`         //
		Benefit         TierBenefit  `db:"benefit"`          //
	}

	TierBenefit struct {
		BrowsingLimit  int64 `json:"browsing_limit"`   //
		AllowSearch    bool  `json:"allow_search"`     //
		AllowVoiceCall bool  `json:"allow_voice_call"` //
		AllowVideoCall bool  `json:"allow_video_call"` //
		AllowChat      bool  `json:"allow_chat"`       //
	}
)

func (m TierBenefit) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *TierBenefit) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}
