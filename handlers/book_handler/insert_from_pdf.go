package book_handler

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/services/book_service"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func InsertFromPdf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(model.PDF_MAX_FILE_SIZE_MB)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	pdfFile, _, err := r.FormFile("pdf_file")
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}
	defer pdfFile.Close()

	pdfBytes, err := io.ReadAll(pdfFile)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	params := contract.InsertFromPdf{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		PdfBytes:    pdfBytes,
	}

	err = book_service.InsertFromPdf(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, nil)
}
