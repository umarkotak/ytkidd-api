package payment_lib

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/order_repo"
	"github.com/umarkotak/ytkidd-api/repos/product_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_subscription_repo"
)

var (
	PROCEED_PROCESS_BENEFIT_ORDER_STATUSES = []string{
		model.ORDER_STATUS_INITIALIZED, model.ORDER_STATUS_PENDING, model.ORDER_STATUS_PAID,
	}
)

func ProcessOrderBenefit(ctx context.Context, orderNumber string) error {
	payment, err := GetByOrderNumber(ctx, orderNumber)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	order, err := order_repo.GetByNumber(ctx, payment.OrderNumber)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	// TODO: add order process benefit lock

	if slices.Contains(model.ORDER_FINAL_STATES, order.Status) {
		return nil
	}

	if !slices.Contains(PROCEED_PROCESS_BENEFIT_ORDER_STATUSES, order.Status) {
		return fmt.Errorf("invalid order status for processing benefit")
	}

	order.PaymentStatus = payment.Status
	order.PaymentNumber = sql.NullString{payment.Number, true}

	if order.PaymentStatus == STATUS_PENDING {
		order.Status = model.ORDER_STATUS_PENDING
	}
	if order.PaymentStatus == STATUS_CANCELED {
		order.Status = model.ORDER_STATUS_CANCELED
	}
	if order.PaymentStatus == STATUS_FAILED {
		order.Status = model.ORDER_STATUS_FAILED
	}
	if order.PaymentStatus == STATUS_EXPIRED {
		order.Status = model.ORDER_STATUS_EXPIRED
	}
	if order.PaymentStatus == STATUS_SUCCESS {
		order.Status = model.ORDER_STATUS_PAID
	}

	err = order_repo.Update(ctx, nil, order)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	if order.Status != model.ORDER_STATUS_PAID {
		return nil
	}

	switch order.OrderType {
	case model.BENEFIT_TYPE_SUBSCRIPTION:
		err = giveBenefitSubscription(ctx, order)
	default:
		err = fmt.Errorf("unsupported order type")
	}
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}

func giveBenefitSubscription(ctx context.Context, order model.Order) error {
	product, err := product_repo.GetByCode(ctx, order.Metadata.ProductCode)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	err = datastore.Transaction(ctx, datastore.Get().Db, func(tx *sqlx.Tx) error {
		now := time.Now()

		user_subscription_repo.Insert(ctx, tx, model.UserSubscription{
			UserID:      order.UserID,
			OrderID:     order.ID,
			ProductCode: product.Code,
			StartedAt:   now,
			EndedAt:     now.Add(time.Duration(product.Metadata.DurationDays+1) * 24 * time.Hour),
		})

		order.Status = model.ORDER_STATUS_COMPLETE
		err = order_repo.Update(ctx, tx, order)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
