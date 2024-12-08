package book_content_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByID(ctx context.Context, id int64) (model.BookContent, error) {
	obj := model.BookContent{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByBookID(ctx context.Context, bookID int64) ([]model.BookContent, error) {
	objs := []model.BookContent{}
	err := stmtGetByBookID.SelectContext(ctx, &objs, map[string]any{
		"book_id": bookID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}
