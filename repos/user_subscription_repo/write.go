package user_subscription_repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func Insert(ctx context.Context, tx *sqlx.Tx, obj model.UserSubscription) (int64, error) {
	var err error

	namedStmt := stmtInsert
	if tx != nil {
		namedStmt, err = tx.PrepareNamedContext(ctx, queryInsert)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return 0, err
		}
	}

	err = namedStmt.GetContext(ctx, &obj.ID, obj)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return 0, err
	}

	return obj.ID, nil
}
