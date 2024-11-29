package middlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/utils/common_ctx"
	"github.com/umarkotak/ytkidd-api/utils/ratelimit_lib"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

// it mean N request / X duration. Eg: 10 request per 2 minute
func ReqRateLimit(maxReqCount int64, duration time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			commonCtx := common_ctx.GetFromCtx(ctx)

			ok, err := ratelimit_lib.Check(
				ctx, commonCtx.DeviceID, maxReqCount, duration,
			)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
			}
			if !ok {
				render.Error(w, r, model.ErrTooManyRequests, "")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
