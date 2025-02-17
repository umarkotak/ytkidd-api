package resp_contract

type (
	GetBooks struct {
		Books []Book `json:"books"`
	}

	Book struct {
		ID           int64    `json:"id"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		CoverFileUrl string   `json:"cover_file_url"`
		Tags         []string `json:"tags"`
		Type         string   `json:"type"`
	}

	BookDetail struct {
		ID           int64         `json:"id"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		CoverFileUrl string        `json:"cover_file_url"`
		Tags         []string      `json:"tags"`
		Type         string        `json:"type"`
		Contents     []BookContent `json:"contents"`
		PdfUrl       string        `json:"pdf_url"`
	}

	BookContent struct {
		ID           int64  `json:"id"`
		BookID       int64  `json:"book_id"`
		PageNumber   int64  `json:"page_number"`
		ImageFileUrl string `json:"image_file_url"`
		Description  string `json:"description"`
	}
)
