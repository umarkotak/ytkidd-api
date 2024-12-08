package model

import (
	"database/sql"
	"time"
)

type (
	Worksheet struct {
		ID            int64        `db:"id"`
		CreatedAt     time.Time    `db:"created_at"`
		UpdatedAt     time.Time    `db:"updated_at"`
		DeletedAt     sql.NullTime `db:"deleted_at"`
		Title         string       `db:"title"`
		CoverFileGuid string       `db:"cover_file_guid"`
	}
)
