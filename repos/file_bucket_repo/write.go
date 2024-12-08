package file_bucket_repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/utils/random"
)

func Insert(ctx context.Context, tx *sqlx.Tx, fileBucket model.FileBucket) (int64, string, error) {
	var err error
	newID := int64(0)

	fileBucket.Guid = random.MustGenUUIDTimes(2)

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
