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

func UpdateChannel(ctx context.Context, params contract.UpdateYoutubeChannel) error {
	youtubeChannel, err := youtube_channel_repo.GetByID(ctx, params.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	youtubeChannel.ExternalID = params.ExternalID
	youtubeChannel.Name = params.Name
	youtubeChannel.Username = params.Username
	youtubeChannel.ImageUrl = params.ImageUrl
	youtubeChannel.Active = params.Active
	youtubeChannel.ChannelLink = params.ChannelLink

	err = youtube_channel_repo.Update(ctx, nil, youtubeChannel)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
