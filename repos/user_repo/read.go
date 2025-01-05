package user_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByID(ctx context.Context, userID int64) (model.User, error) {
	user := model.User{}

	err := stmtGetByID.GetContext(ctx, &user, map[string]any{
		"id": userID,
	})
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"user_id": userID,
		}).Error(err)
		return user, err
	}

	return user, nil
}

func GetByGuid(ctx context.Context, userGuid string) (model.User, error) {
	user := model.User{}

	err := stmtGetByGuid.GetContext(ctx, &user, map[string]any{
		"guid": userGuid,
	})
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"user_guid": userGuid,
		}).Error(err)
		return user, err
	}

	return user, nil
}

func GetByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{}

	err := stmtGetByEmail.GetContext(ctx, &user, map[string]any{
		"email": email,
	})
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"email": email,
		}).Error(err)
		return user, err
	}

	return user, nil
}
