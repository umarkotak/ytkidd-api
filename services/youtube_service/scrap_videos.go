package youtube_service

import (
	"context"
	"database/sql"
	"html"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/repos/youtube_channel_repo"
	"github.com/umarkotak/ytkidd-api/repos/youtube_video_repo"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func ScrapVideos(ctx context.Context, params contract.ScrapVideos) (string, error) {
	nextPageToken := ""

	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(config.Get().YoutubeApiKey))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nextPageToken, err
	}

	call := youtubeService.Search.List([]string{"id", "snippet"})
	call = call.ChannelId(params.ChannelID)
	call = call.Q(params.Query)
	call = call.Type("video")
	call = call.PageToken(params.PageToken)
	call = call.MaxResults(50) // Get up to 50 results

	response, err := call.Do()
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return nextPageToken, err
	}

	for _, item := range response.Items {
		if item.Id.Kind != "youtube#video" {
			continue
		}

		youtubeChannel, err := youtube_channel_repo.GetByExternalID(ctx, params.ChannelID)
		if err != nil && err != sql.ErrNoRows {
			logrus.WithContext(ctx).Error(err)
			return nextPageToken, err
		}

		if youtubeChannel.ID == 0 {
			youtubeChannel = model.YoutubeChannel{
				ExternalID: params.ChannelID,
				Name:       html.UnescapeString(item.Snippet.ChannelTitle),
				Username:   html.UnescapeString(item.Snippet.ChannelTitle),
				ImageUrl:   "",
				Tags:       []string{},
			}
			youtubeChannel.ID, err = youtube_channel_repo.Insert(ctx, nil, youtubeChannel)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return nextPageToken, err
			}
		}

		youtubeVideo, err := youtube_video_repo.GetByExternalID(ctx, item.Id.VideoId)
		if err != nil && err != sql.ErrNoRows {
			logrus.WithContext(ctx).Error(err)
			return nextPageToken, err
		}

		if youtubeVideo.ID != 0 {
			continue
		}

		youtubeVideo = model.YoutubeVideo{
			YoutubeChannelID: youtubeChannel.ID,
			ExternalId:       item.Id.VideoId,
			Title:            html.UnescapeString(item.Snippet.Title),
			ImageUrl:         item.Snippet.Thumbnails.Medium.Url,
			Tags:             []string{},
			Active:           true,
		}
		youtubeVideo.ID, err = youtube_video_repo.Insert(ctx, nil, youtubeVideo)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return nextPageToken, err
		}
	}

	nextPageToken = response.NextPageToken

	return nextPageToken, nil
}
