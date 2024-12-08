package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type (
	Book struct {
		ID            int64          `db:"id"`
		CreatedAt     time.Time      `db:"created_at"`
		UpdatedAt     time.Time      `db:"updated_at"`
		DeletedAt     sql.NullTime   `db:"deleted_at"`
		Title         string         `db:"title"`
		Description   string         `db:"description"`
		CoverImageUrl string         `db:"cover_image_url"`
		Tags          pq.StringArray `db:"tags"`
		Type          string         `db:"type"`
		PdfFileUrl    string         `db:"pdf_file_url"`
		Active        bool           `db:"active"`
	}
)
