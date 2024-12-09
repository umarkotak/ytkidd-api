package contract

import "github.com/lib/pq"

type (
	InsertFromPdf struct {
		Title       string
		Description string
		PdfBytes    []byte
	}

	GetBooks struct {
		BookID int64          `db:"-"`
		Title  string         `db:"title"`
		Tags   pq.StringArray `db:"tags"`
		Type   string         `db:"type"`
	}
)
