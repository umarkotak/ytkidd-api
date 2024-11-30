package youtube_channel_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/services/youtube_channel_service"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func GetYoutubeChannels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := contract.GetYoutubeChannels{
		Name: r.URL.Query().Get("name"),
		Tags: utils.StringMustSliceString(r.URL.Query().Get("tags"), ","),
	}

	youtubeChannels, err := youtube_channel_service.GetChannels(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		render.Error(w, r, err, "")
		return
	}

	render.Response(w, r, 200, youtubeChannels)
}
