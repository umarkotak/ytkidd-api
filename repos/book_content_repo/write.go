package book_content_repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func Insert(ctx context.Context, tx *sqlx.Tx, bookContent model.BookContent) (int64, error) {
	var err error
	newID := int64(0)

	stmt := stmtInsert
	if tx != nil {
		stmt, err = tx.PrepareNamedContext(ctx, queryInsert)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return newID, err
		}
	}

	err = stmt.GetContext(ctx, &newID, bookContent)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return newID, err
	}

	return newID, nil
}

func Update(ctx context.Context, tx *sqlx.Tx, bookContent model.BookContent) error {
	var err error

	stmt := stmtUpdate
	if tx != nil {
		stmt, err = tx.PrepareNamedContext(ctx, queryUpdate)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}
	}

	_, err = stmt.ExecContext(ctx, bookContent)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}

func SoftDelete(ctx context.Context, tx *sqlx.Tx, id int64) error {
	var err error

	stmt := stmtSoftDelete
	if tx != nil {
		stmt, err = tx.PrepareNamedContext(ctx, querySoftDelete)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}
	}

	_, err = stmt.ExecContext(ctx, map[string]any{
		"id": id,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
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
