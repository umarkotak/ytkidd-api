package product_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByID(ctx context.Context, id int64) (model.Product, error) {
	obj := model.Product{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByCode(ctx context.Context, code string) (model.Product, error) {
	obj := model.Product{}
	err := stmtGetByCode.GetContext(ctx, &obj, map[string]any{
		"code": code,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}
