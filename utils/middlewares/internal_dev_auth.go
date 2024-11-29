package middlewares

import (
	"net/http"
	"strings"

	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func InternalDevAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		tokens := strings.Split(authorization, " ")

		if len(tokens) != 2 || tokens[0] != "Basic" {
			render.Error(w, r, model.ErrUnauthorized, "")
			return
		}

		username, password, err := decodeBasicAuth(tokens[1])
		if err != nil {
			render.Error(w, r, model.ErrUnauthorized, "")
			return
		}

		if username != config.Get().DevInternalClientID || password != config.Get().DevInternalSecretKey {
			render.Error(w, r, model.ErrUnauthorized, "")
			return
		}

		next.ServeHTTP(w, r)
	})
}
