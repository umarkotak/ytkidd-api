package model

import (
	"database/sql"
	"time"
)

type (
	UserSubscription struct {
		ID          int64        `db:"id"`
		CreatedAt   time.Time    `db:"created_at"`
		UpdatedAt   time.Time    `db:"updated_at"`
		DeletedAt   sql.NullTime `db:"deleted_at"`
		UserID      int64        `db:"user_id"`
		OrderID     int64        `db:"order_id"`
		ProductCode string       `db:"product_code"`
		StartedAt   time.Time    `db:"started_at"`
		EndedAt     time.Time    `db:"ended_at"`
	}
)
