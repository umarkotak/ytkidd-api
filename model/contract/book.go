package contract

import "github.com/lib/pq"

type (
	InsertFromPdf struct {
		Slug            string
		Title           string
		Description     string
		PdfBytes        []byte
		ImgFormat       string
		BookType        string
		CustomImageSlug string
		OriginalPdfUrl  string
	}

	InsertFromPdfUrl struct {
		PdfUrl          string `json:"pdf_url"`
		Slug            string `json:"slug"`
		Title           string `json:"title"`
		Description     string `json:"description"`
		ImgFormat       string `json:"img_format"`
		BookType        string `json:"book_type"`
		CustomImageSlug string `json:"custom_image_slug"`
	}

	GetBooks struct {
		BookID int64          `db:"-"`
		Title  string         `db:"title"`
		Tags   pq.StringArray `db:"tags"`
		Types  pq.StringArray `db:"types"`
	}

	DeleteBook struct {
		BookID int64 `db:"id"`
	}
)
