package middlewares

import (
	"context"
	"net/http"

	"github.com/umarkotak/ytkidd-api/utils/common_ctx"
)

func CommonCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commonCtx := common_ctx.FromRequestHeader(r)

		ctx := context.WithValue(r.Context(), common_ctx.CommonCtxKey, commonCtx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
