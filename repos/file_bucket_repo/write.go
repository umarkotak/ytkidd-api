package file_bucket_repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/utils/random"
)

func Insert(ctx context.Context, tx *sqlx.Tx, fileBucket model.FileBucket) (int64, string, error) {
	var err error
	newID := int64(0)

	if fileBucket.Guid == "" {
		fileBucket.Guid = random.MustGenUUIDTimes(2)
	}

	stmt := stmtInsert
	if tx != nil {
		stmt, err = tx.PrepareNamedContext(ctx, queryInsert)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return newID, "", err
		}
	}

	err = stmt.GetContext(ctx, &newID, fileBucket)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return newID, "", err
	}

	return newID, fileBucket.Guid, nil
}

func DeleteByGuids(ctx context.Context, tx *sqlx.Tx, guids pq.StringArray) error {
	var err error

	stmt := stmtDeleteByGuids
	if tx != nil {
		stmt, err = tx.PrepareNamedContext(ctx, queryDeleteByGuids)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}
	}

	_, err = stmt.ExecContext(ctx, map[string]any{
		"guids": guids,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
