package middlewares

import (
	"bytes"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBuff, _ := io.ReadAll(r.Body)
		reader := io.NopCloser(bytes.NewBuffer(bodyBuff))
		r.Body = reader
		logrus.WithContext(r.Context()).Infof("[REQUEST][PATH : %v %v]", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
