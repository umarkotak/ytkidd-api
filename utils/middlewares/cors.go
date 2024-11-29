package middlewares

import (
	"net/http"

	"github.com/umarkotak/ytkidd-api/utils/render"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.SetCorsHeaders(w)

		if r.Method == "OPTIONS" {
			render.Response(w, r, 200, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}
