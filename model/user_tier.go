package model

import (
	"database/sql"
	"time"
)

const (
	TIER_STATUS_ACTIVE   = "active"
	TIER_STATUS_REFUNDED = "refunded"
)

type (
	UserTier struct {
		ID          int64        `db:"id"`            //
		CreatedAt   time.Time    `db:"created_at"`    //
		UpdatedAt   time.Time    `db:"updated_at"`    //
		DeletedAt   sql.NullTime `db:"deleted_at"`    //
		UserId      int64        `db:"user_id"`       //
		Status      string       `db:"status"`        //
		TierPlanId  int64        `db:"tier_plan_id"`  //
		StartedAt   time.Time    `db:"started_at"`    //
		EndedAt     time.Time    `db:"ended_at"`      //
		OrderId     int64        `db:"order_id"`      //
		OrderItemId int64        `db:"order_item_id"` //
	}

	UserTierActive struct {
		ID               int64        `db:"id"`                 //
		CreatedAt        time.Time    `db:"created_at"`         //
		UpdatedAt        time.Time    `db:"updated_at"`         //
		DeletedAt        sql.NullTime `db:"deleted_at"`         //
		UserId           int64        `db:"user_id"`            //
		Status           string       `db:"status"`             //
		TierPlanId       int64        `db:"tier_plan_id"`       //
		StartedAt        time.Time    `db:"started_at"`         //
		EndedAt          time.Time    `db:"ended_at"`           //
		OrderId          int64        `db:"order_id"`           //
		OrderItemId      int64        `db:"order_item_id"`      //
		TierPlanCodename string       `db:"tier_plan_codename"` //
		TierPlanBenefit  TierBenefit  `db:"tier_plan_benefit"`  //
	}
)
