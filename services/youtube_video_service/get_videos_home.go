package youtube_video_service

import (
	"context"
	"slices"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/youtube_video_repo"
)

func GetVideos(ctx context.Context, params contract.GetYoutubeVideos) (contract_resp.YoutubeVideosHome, error) {
	youtubeVideosHome := contract_resp.YoutubeVideosHome{}

	youtubeVideosDetailed, err := youtube_video_repo.GetByParams(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return youtubeVideosHome, err
	}

	excludeIDs := []int64{}
	youtubeVideosHomeVideo := []contract_resp.YoutubeVideo{}
	for _, videoDetailed := range youtubeVideosDetailed {
		youtubeVideosHomeVideo = append(youtubeVideosHomeVideo, contract_resp.YoutubeVideo{
			ID:        videoDetailed.ID,
			ImageUrl:  videoDetailed.ImageUrl,
			Title:     videoDetailed.Title,
			Tags:      videoDetailed.Tags,
			CanAction: slices.Contains(model.ADMIN_ROLES, params.UserRole),
			Channel: contract_resp.YoutubeChannel{
				ID:       videoDetailed.YoutubeChannelID,
				ImageUrl: videoDetailed.YoutubeChannelImageUrl,
				Name:     videoDetailed.YoutubeChannelName,
				Tags:     videoDetailed.YoutubeChannelTags,
			},
		})

		excludeIDs = append(excludeIDs, videoDetailed.ID)
	}

	youtubeVideosHome = contract_resp.YoutubeVideosHome{
		Videos:     youtubeVideosHomeVideo,
		ExcludeIDs: excludeIDs,
	}

	return youtubeVideosHome, nil
}
