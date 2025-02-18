package youtube_video_repo

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
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
	objs := []model.YoutubeVideoDetailed{}

	params.SetDefault()
	if params.Tags == nil {
		params.Tags = []string{}
	}
	if params.ChannelIDs == nil {
		params.ChannelIDs = []int64{}
	}
	if params.ExcludeIDs == nil {
		params.ExcludeIDs = []int64{}
	}
	if params.ExcludeChannelIDs == nil {
		params.ExcludeChannelIDs = []int64{}
	}

	if params.Sort == "" {
		params.Sort = "RANDOM()"
	} else if params.Sort == "id_desc" {
		params.Sort = "ytvid.id DESC"
	} else {
		params.Sort = "RANDOM()"
	}

	query := fmt.Sprintf(queryGetByParams, params.Sort)
	stmtGetByParams, err := datastore.Get().Db.PrepareNamed(query)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}

	err = stmtGetByParams.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}
