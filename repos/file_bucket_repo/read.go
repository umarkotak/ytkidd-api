package file_bucket_repo

import (
	"context"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByID(ctx context.Context, id int64) (model.FileBucket, error) {
	obj := model.FileBucket{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByGuid(ctx context.Context, guid string) (model.FileBucket, error) {
	obj := model.FileBucket{}
	err := stmtGetByGuid.GetContext(ctx, &obj, map[string]any{
		"guid": guid,
	})
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"guid": guid,
		}).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByGuids(ctx context.Context, guids pq.StringArray) ([]model.FileBucket, error) {
	objs := []model.FileBucket{}
	err := stmtGetByGuids.SelectContext(ctx, &objs, map[string]any{
		"guids": guids,
	})
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"guids": guids,
		}).Error(err)
		return objs, err
	}
	return objs, nil
}
