package youtube_video_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/services/youtube_video_service"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func GetYoutubeVideosHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.GetYoutubeVideosHome{
		Tags:       utils.StringMustSliceString(r.URL.Query().Get("tags"), ","),
		ExcludeIDs: utils.StringMustSliceInt64(r.URL.Query().Get("exclude_ids"), ","),
	}

	youtubeVideosHome, err := youtube_video_service.GetVideosHome(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, youtubeVideosHome)
}
