package book_handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/services/book_service"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.GetBooks{}

	getBooks, err := book_service.GetBooks(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, getBooks)
}

func GetBookDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.GetBooks{
		BookID: utils.StringMustInt64(chi.URLParam(r, "id")),
	}

	bookDetail, err := book_service.GetBookDetail(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, bookDetail)
}
