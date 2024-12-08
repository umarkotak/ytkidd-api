package model

import (
	"database/sql"
	"time"
)

type (
	WorksheetContent struct {
		ID            int64        `db:"id"`
		CreatedAt     time.Time    `db:"created_at"`
		UpdatedAt     time.Time    `db:"updated_at"`
		DeletedAt     sql.NullTime `db:"deleted_at"`
		WorksheetID   int64        `db:"worksheet_id"`
		Title         string       `db:"title"`
		Description   string       `db:"description"`
		ImageFileGuid string       `db:"image_file_guid"`
	}
)
