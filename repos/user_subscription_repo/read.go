package user_subscription_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByID(ctx context.Context, id int64) (model.UserSubscription, error) {
	obj := model.UserSubscription{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByUserID(ctx context.Context, userID int64) ([]model.UserSubscription, error) {
	objs := []model.UserSubscription{}
	err := stmtGetByUserID.SelectContext(ctx, &objs, map[string]any{
		"user_id": userID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}

func GetActiveByUserID(ctx context.Context, userID int64) ([]model.UserSubscription, error) {
	objs := []model.UserSubscription{}
	err := stmtGetActiveByUserID.SelectContext(ctx, &objs, map[string]any{
		"user_id": userID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}

func GetUserLatestActiveSubscription(ctx context.Context, userID int64) (model.UserSubscription, error) {
	obj := model.UserSubscription{}
	err := stmtGetUserLatestActiveSubscription.GetContext(ctx, &obj, map[string]any{
		"user_id": userID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}
