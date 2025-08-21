package order_handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
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

func GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commonContext := common_ctx.GetFromCtx(ctx)

	data, err := order_service.GetOrderDetail(ctx, commonContext.UserAuth.GUID, chi.URLParam(r, "order_number"))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, http.StatusOK, data)
}

func PostCheckOrderPayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commonContext := common_ctx.GetFromCtx(ctx)

	data, err := order_service.CheckOrderPayment(ctx, commonContext.UserAuth.GUID, chi.URLParam(r, "order_number"))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, http.StatusOK, data)
}

func GetOrderList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commonContext := common_ctx.GetFromCtx(ctx)

	user, err := user_repo.GetByGuid(ctx, commonContext.UserAuth.GUID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	params := contract.GetOrderByParams{
		UserID: user.ID,
		Pagination: model.Pagination{
			Limit: utils.StringMustInt64(r.URL.Query().Get("limit")),
			Page:  utils.StringMustInt64(r.URL.Query().Get("page")),
		},
	}
	params.Pagination.SetDefault()

	data, err := order_service.GetOrderList(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, http.StatusOK, data)
}
