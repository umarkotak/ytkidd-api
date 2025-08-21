package youtube_channel_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByID(ctx context.Context, id int64) (model.YoutubeChannel, error) {
	obj := model.YoutubeChannel{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByExternalID(ctx context.Context, externalID string) (model.YoutubeChannel, error) {
	obj := model.YoutubeChannel{}
	err := stmtGetByExternalID.GetContext(ctx, &obj, map[string]any{
		"external_id": externalID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetForSearch(ctx context.Context, params contract.GetYoutubeChannels) ([]model.YoutubeChannel, error) {
	objs := []model.YoutubeChannel{}
	err := stmtGetForSearch.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}
