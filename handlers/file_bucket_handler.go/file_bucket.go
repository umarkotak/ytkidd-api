package file_bucket_handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/repos/file_bucket_repo"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func GetByGuid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fileGuid := chi.URLParam(r, "guid")

	fileBucket, err := file_bucket_repo.GetByGuid(ctx, fileGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	w.Header().Set("Content-Type", fileBucket.HttpContentType)
	_, err = w.Write(fileBucket.Data)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}
}
