package youtube_video_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/youtube_channel_repo"
	"github.com/umarkotak/ytkidd-api/repos/youtube_video_repo"
)

func GetVideoDetail(ctx context.Context, youtubeVideoID int64) (resp_contract.YoutubeVideo, error) {
	youtubeVideoDetail := resp_contract.YoutubeVideo{}

	youtubeVideo, err := youtube_video_repo.GetByID(ctx, youtubeVideoID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return youtubeVideoDetail, err
	}

	youtubeChannel, err := youtube_channel_repo.GetByID(ctx, youtubeVideo.YoutubeChannelID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return youtubeVideoDetail, err
	}

	youtubeVideoDetail = resp_contract.YoutubeVideo{
		ID:         youtubeVideo.ID,
		ImageUrl:   youtubeVideo.ImageUrl,
		Title:      youtubeVideo.Title,
		Tags:       youtubeVideo.Tags,
		ExternalID: youtubeVideo.ExternalId,
		Channel: resp_contract.YoutubeChannel{
			ID:       youtubeChannel.ID,
			ImageUrl: youtubeChannel.ImageUrl,
			Name:     youtubeChannel.Name,
			Tags:     youtubeChannel.Tags,
		},
	}

	return youtubeVideoDetail, nil
}
