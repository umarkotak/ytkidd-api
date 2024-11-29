package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/utils/common_ctx"
	"github.com/umarkotak/ytkidd-api/utils/render"
	"github.com/umarkotak/ytkidd-api/utils/user_auth"
)

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			err := fmt.Errorf("unauthorized")
			render.RawError(w, r, 401, err, "unauthorized", "missing bearer token")
			return
		}

		splitted := strings.Split(bearerToken, " ")

		if len(splitted) != 2 {
			err := fmt.Errorf("unauthorized")
			render.RawError(w, r, 401, err, "unauthorized", "invalid bearer token")
			return
		}

		if splitted[0] != "Bearer" {
			err := fmt.Errorf("unauthorized")
			render.RawError(w, r, 401, err, "unauthorized", "invalid bearer token")
			return
		}

		accessToken := splitted[1]

		userAuth, errMsg, err := user_auth.VerifyAccessToken(ctx, accessToken)
		if err != nil {
			render.RawError(w, r, 401, err, "unauthorized", errMsg)
			return
		}

		commonCtx := common_ctx.Get(r)

		commonCtx.UserAuth = common_ctx.UserAuth{
			GUID: userAuth.GUID,
			Name: userAuth.Name,
		}

		ctx = context.WithValue(r.Context(), common_ctx.CommonCtxKey, commonCtx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func OptionalUserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		commonCtx := common_ctx.GetFromCtx(ctx)

		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			next.ServeHTTP(w, r)
			return
		}

		splitted := strings.Split(bearerToken, " ")

		if len(splitted) != 2 {
			next.ServeHTTP(w, r)
			return
		}

		if splitted[0] != "Bearer" {
			next.ServeHTTP(w, r)
			return
		}

		accessToken := splitted[1]

		userAuth, errMsg, err := user_auth.VerifyAccessToken(ctx, accessToken)
		if err != nil {
			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"err_msg": errMsg,
			}).Error(err)
			next.ServeHTTP(w, r)
			return
		}

		commonCtx.UserAuth = common_ctx.UserAuth{
			GUID: userAuth.GUID,
			Name: userAuth.Name,
		}

		ctx = context.WithValue(r.Context(), common_ctx.CommonCtxKey, commonCtx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
