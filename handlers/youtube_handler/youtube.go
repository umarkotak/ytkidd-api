package youtube_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/services/youtube_service"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func ScrapVideos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.ScrapVideos{}
	err := utils.BindJson(r, &params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	for i := 1; i <= model.YOUTUBE_MAX_PAGE; i++ {
		nextPageToken, err := youtube_service.ScrapVideos(ctx, params)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			render.Error(w, r, err, "")
			return
		}

		if !params.All {
			break
		}

		params.PageToken = nextPageToken
		if params.PageToken == "" {
			break
		}
	}

	render.Response(w, r, 200, nil)
}
