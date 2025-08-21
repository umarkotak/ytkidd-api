package youtube_channel_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/repos/youtube_channel_repo"
)

func GetChannels(ctx context.Context, params contract.GetYoutubeChannels) ([]contract_resp.YoutubeChannel, error) {
	respYoutubeChannels := []contract_resp.YoutubeChannel{}

	youtubeChannels, err := youtube_channel_repo.GetForSearch(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return respYoutubeChannels, err
	}

	for _, youtubeChannel := range youtubeChannels {
		respYoutubeChannels = append(respYoutubeChannels, contract_resp.YoutubeChannel{
			ID:         youtubeChannel.ID,
			ImageUrl:   youtubeChannel.ImageUrl,
			Name:       youtubeChannel.Name,
			Tags:       youtubeChannel.Tags,
			ExternalID: youtubeChannel.ExternalID,
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
