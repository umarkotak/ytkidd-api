package book_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.GetBooks{}
	err := utils.BindJson(r, &params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, nil)
}

func GetBookDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.GetBooks{}
	err := utils.BindJson(r, &params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, nil)
}
