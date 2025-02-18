package book_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
)

func GetByID(ctx context.Context, id int64) (model.Book, error) {
	obj := model.Book{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByParams(ctx context.Context, params contract.GetBooks) ([]model.Book, error) {
	if params.Tags == nil {
		params.Tags = []string{}
	}
	if params.Types == nil {
		params.Types = []string{}
	}

	objs := []model.Book{}
	err := stmtGetByParams.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}
