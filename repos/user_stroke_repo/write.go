package user_stroke_repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func Insert(ctx context.Context, tx *sqlx.Tx, obj model.UserStroke) (int64, error) {
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

func Upsert(ctx context.Context, tx *sqlx.Tx, obj model.UserStroke) (int64, error) {
	var err error

	namedStmt := stmtUpsert
	if tx != nil {
		namedStmt, err = tx.PrepareNamedContext(ctx, queryUpsert)
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

func DeleteByBookID(ctx context.Context, tx *sqlx.Tx, bookID int64) error {
	var err error

	stmt := stmtDeleteByBookID
	if tx != nil {
		stmt, err = tx.PrepareNamedContext(ctx, queryDeleteByBookID)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}
	}

	_, err = stmt.ExecContext(ctx, map[string]any{
		"book_id": bookID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
