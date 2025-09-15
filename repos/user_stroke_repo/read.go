package user_stroke_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

func GetByUserAndContent(ctx context.Context, userID int64, appSession string, bookContentID int64) (model.UserStroke, error) {
	// logrus.Infof("DAT: %+v", userID)
	// logrus.Infof("DAT: %+v", appSession)
	// logrus.Infof("DAT: %+v", bookContentID)

	obj := model.UserStroke{}
	err := stmtGetByUserAndContent.GetContext(ctx, &obj, map[string]any{
		"user_id":         userID,
		"app_session":     appSession,
		"book_content_id": bookContentID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return obj, err
	}
	return obj, nil
}
