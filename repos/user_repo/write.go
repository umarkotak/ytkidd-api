package user_repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func Insert(ctx context.Context, tx *sqlx.Tx, user model.User) (int64, string, error) {
	randomUUID, _ := uuid.NewRandom()
	user.Guid = randomUUID.String()

	var row *sqlx.Row
	if tx == nil {
		row = stmtInsert.QueryRowContext(ctx, user)

	} else {
		namedStmt, err := tx.PrepareNamedContext(ctx, queryInsert)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return 0, "", err
		}

		row = namedStmt.QueryRowContext(ctx, user)
	}

	err := row.Scan(&user.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return 0, "", err
	}

	return user.ID, user.Guid, nil
}

func Update(ctx context.Context, tx *sqlx.Tx, user model.User) error {
	var err error
	var namedStmt *sqlx.NamedStmt

	if tx == nil {
		_, err = stmtUpdate.ExecContext(ctx, user)

	} else {
		namedStmt, err = tx.PrepareNamedContext(ctx, queryUpdate)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		_, err = namedStmt.ExecContext(ctx, user)
	}
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}

func SoftDelete(ctx context.Context, tx *sqlx.Tx, userID int64) error {
	var err error
	var namedStmt *sqlx.NamedStmt

	if tx == nil {
		_, err = stmtSoftDelete.ExecContext(ctx, map[string]any{
			"id": userID,
		})

	} else {
		namedStmt, err = tx.PrepareNamedContext(ctx, querySoftDelete)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		_, err = namedStmt.ExecContext(ctx, map[string]any{
			"id": userID,
		})
	}
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
