package order_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByID(ctx context.Context, id int64) (model.Order, error) {
	obj := model.Order{}
	err := stmtGetByID.GetContext(ctx, &obj, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByNumber(ctx context.Context, number string) (model.Order, error) {
	obj := model.Order{}
	err := stmtGetByNumber.GetContext(ctx, &obj, map[string]any{
		"number": number,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}

func GetByUserID(ctx context.Context, params contract.GetOrderByUserID) ([]model.Order, error) {
	objs := []model.Order{}
	err := stmtGetByUserID.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}

func GetByParams(ctx context.Context, params contract.GetOrderByParams) ([]model.Order, error) {
	objs := []model.Order{}
	err := stmtGetByParams.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}

func GetByParamsWithPayment(ctx context.Context, params contract.GetOrderByParams) ([]model.OrderWithPayment, error) {
	objs := []model.OrderWithPayment{}
	err := stmtGetByParamsWithPayment.SelectContext(ctx, &objs, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return objs, err
	}
	return objs, nil
}
