package order_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/services/order_service"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/common_ctx"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func PostCreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commonContext := common_ctx.GetFromCtx(ctx)

	params := contract.CreateOrder{
		UserGuid: commonContext.UserAuth.GUID,
	}
	err := utils.BindJson(r, &params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	data, err := order_service.CreateOrder(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, http.StatusOK, data)
}
