package file_bucket_repo

import (
	"context"

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
