package contract

type (
	InsertFromPdf struct {
		Title       string
		Description string
		PdfBytes    []byte
	}

	GetBooks struct {
	}
)
