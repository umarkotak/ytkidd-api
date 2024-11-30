package youtube_video_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/youtube_video_repo"
)

func GetVideosHome(ctx context.Context, params contract.GetYoutubeVideosHome) (resp_contract.YoutubeVideosHome, error) {
	youtubeVideosHome := resp_contract.YoutubeVideosHome{}

	youtubeVideosDetailed, err := youtube_video_repo.GetForHome(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return youtubeVideosHome, err
	}

	excludeIDs := []int64{}
	youtubeVideosHomeVideo := []resp_contract.YoutubeVideosHomeVideo{}
	for _, videoDetailed := range youtubeVideosDetailed {
		youtubeVideosHomeVideo = append(youtubeVideosHomeVideo, resp_contract.YoutubeVideosHomeVideo{
			ID:       videoDetailed.ID,
			ImageUrl: videoDetailed.ImageUrl,
			Title:    videoDetailed.Title,
			Tags:     videoDetailed.Tags,
			Channel: resp_contract.YoutubeVideosHomeChannel{
				ID:       videoDetailed.YoutubeChannelID,
				ImageUrl: videoDetailed.YoutubeChannelImageUrl,
				Name:     videoDetailed.YoutubeChannelName,
				Tags:     videoDetailed.YoutubeChannelTags,
			},
		})

		excludeIDs = append(excludeIDs, videoDetailed.ID)
	}

	youtubeVideosHome = resp_contract.YoutubeVideosHome{
		Channels:   []resp_contract.YoutubeVideosHomeChannel{},
		Videos:     youtubeVideosHomeVideo,
		ExcludeIDs: excludeIDs,
	}

	return youtubeVideosHome, nil
}
