package user_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/services/user_service"
	"github.com/umarkotak/ytkidd-api/services/user_subscription_service"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/common_ctx"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.UserSignIn{}

	err := utils.BindJson(r, &params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	data, err := user_service.SignIn(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, data)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commonCtx := common_ctx.GetFromCtx(ctx)

	userSubscription, err := user_subscription_service.GetUserSubscriptionInfo(ctx, commonCtx.UserAuth.GUID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	data := map[string]any{
		"user":         commonCtx.UserAuth,
		"subscription": userSubscription,
	}

	render.Response(w, r, 200, data)
}

func MyProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commonCtx := common_ctx.GetFromCtx(ctx)

	render.Response(w, r, 200, commonCtx)
}
