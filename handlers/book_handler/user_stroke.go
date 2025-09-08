package book_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/services/book_service"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/common_ctx"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func GetUserStroke(w http.ResponseWriter, r *http.Request) {
	var err error

	ctx := r.Context()

	commonCtx := common_ctx.GetFromCtx(ctx)

	params := contract.GetUserStroke{
		UserGuid:      commonCtx.UserAuth.GUID,
		AppSession:    commonCtx.AppSession,
		BookContentID: utils.StringMustInt64(r.URL.Query().Get("book_content_id")),
	}

	data, err := book_service.GetUserStroke(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, http.StatusOK, data)
}

func StoreUserStroke(w http.ResponseWriter, r *http.Request) {
	var err error

	ctx := r.Context()

	commonCtx := common_ctx.GetFromCtx(ctx)

	params := contract.StoreUserStroke{
		UserGuid:   commonCtx.UserAuth.GUID,
		AppSession: commonCtx.AppSession,
	}
	err = utils.BindJson(r, &params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	err = book_service.StoreUserStroke(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, http.StatusOK, nil)
}
