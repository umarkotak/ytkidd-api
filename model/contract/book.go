package contract

import (
	"github.com/lib/pq"
	"github.com/umarkotak/ytkidd-api/model"
)

type (
	InsertFromPdf struct {
		Slug           string
		Title          string
		Description    string
		PdfBytes       []byte
		ImgFormat      string
		BookType       string
		OriginalPdfUrl string
		Storage        string // Enum: local, r2
		StorePdf       bool
	}

	InsertFromPdfUrl struct {
		PdfUrl      string `json:"pdf_url"`
		Slug        string `json:"slug"`
		Title       string `json:"title"`
		Description string `json:"description"`
		ImgFormat   string `json:"img_format"`
		BookType    string `json:"book_type"`
		Storage     string `json:"storage"`
		StorePdf    bool   `json:"store_pdf"`
	}

	GetBooks struct {
		UserGuid string         `db:"-"`
		BookID   int64          `db:"-"`
		Title    string         `db:"title"`
		Tags     pq.StringArray `db:"tags"`
		Types    pq.StringArray `db:"types"`
		Sort     string         `db:"sort"`
		model.Pagination
	}

	DeleteBook struct {
		BookID int64 `db:"id"`
	}
)
