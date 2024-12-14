package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type (
	Book struct {
		ID             int64          `db:"id"`
		CreatedAt      time.Time      `db:"created_at"`
		UpdatedAt      time.Time      `db:"updated_at"`
		DeletedAt      sql.NullTime   `db:"deleted_at"`
		Title          string         `db:"title"`
		Description    string         `db:"description"`
		CoverFileGuid  string         `db:"cover_file_guid"`
		Tags           pq.StringArray `db:"tags"`
		Type           string         `db:"type"`
		PdfFileGuid    string         `db:"pdf_file_guid"`
		Active         bool           `db:"active"`
		OriginalPdfUrl string         `db:"original_pdf_url"`

		CoverFilePath string `db:"cover_file_path"` // join from file bucket
	}
)
