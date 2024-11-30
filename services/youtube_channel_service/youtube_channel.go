package youtube_channel_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/youtube_channel_repo"
)

func GetChannels(ctx context.Context, params contract.GetYoutubeChannels) ([]resp_contract.YoutubeChannel, error) {
	respYoutubeChannels := []resp_contract.YoutubeChannel{}

	youtubeChannels, err := youtube_channel_repo.GetForSearch(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return respYoutubeChannels, err
	}

	for _, youtubeChannel := range youtubeChannels {
		respYoutubeChannels = append(respYoutubeChannels, resp_contract.YoutubeChannel{
			ID:       youtubeChannel.ID,
			ImageUrl: youtubeChannel.ImageUrl,
			Name:     youtubeChannel.Name,
			Tags:     youtubeChannel.Tags,
		})
	}

	return respYoutubeChannels, nil
}
