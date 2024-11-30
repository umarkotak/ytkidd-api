package youtube_video_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
)

func GetByID(ctx context.Context, id int64) (model.YoutubeVideo, error) {
	obj := model.YoutubeVideo{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByExternalID(ctx context.Context, externalID string) (model.YoutubeVideo, error) {
	obj := model.YoutubeVideo{}
	err := stmtGetByExternalID.GetContext(ctx, &obj, map[string]any{
		"external_id": externalID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetForSearch(ctx context.Context, params contract.YoutubeVideoSearch) ([]model.YoutubeVideo, error) {
	objs := []model.YoutubeVideo{}
	err := stmtGetForSearch.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}

func GetByParams(ctx context.Context, params contract.GetYoutubeVideos) ([]model.YoutubeVideoDetailed, error) {
	params.SetDefault()
	objs := []model.YoutubeVideoDetailed{}
	err := stmtGetByParams.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}
