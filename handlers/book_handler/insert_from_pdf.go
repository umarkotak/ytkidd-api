package book_handler

import (
	"fmt"
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
	if err != nil && err != http.ErrMissingFile {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	pdfBytes := []byte{}
	if pdfFile != nil {
		defer pdfFile.Close()
		pdfBytes, err = io.ReadAll(pdfFile)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			render.Error(w, r, err, "")
			return
		}
	}

	if r.FormValue("pdf_url") != "" {
		httpClient := http.Client{}
		resp, err := httpClient.Get(r.FormValue("pdf_url"))
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			render.Error(w, r, err, "")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logrus.WithContext(ctx).Error(err)
			render.Error(w, r, err, "")
			return
		}

		pdfBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			render.Error(w, r, err, "")
			return
		}
	}

	if len(pdfBytes) == 0 {
		err = fmt.Errorf("pdf bytes empty")
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
