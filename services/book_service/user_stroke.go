package book_service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_stroke_repo"
)

func GetUserStroke(ctx context.Context, params contract.GetUserStroke) (contract_resp.GetUserStroke, error) {
	var err error

	if params.AppSession == "" {
		err = fmt.Errorf("missing app session")
		return contract_resp.GetUserStroke{}, err
	}

	var userID sql.NullInt64
	if params.UserGuid != "" {
		user, err := user_repo.GetByGuid(ctx, params.UserGuid)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return contract_resp.GetUserStroke{}, err
		}
		userID = sql.NullInt64{user.ID, true}
	}

	userStroke, err := user_stroke_repo.GetByUserAndContent(ctx, userID, sql.NullString{params.AppSession, true}, sql.NullInt64{params.BookContentID, true})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.GetUserStroke{}, err
	}

	return contract_resp.GetUserStroke{
		ID:            userStroke.ID,
		BookID:        userStroke.BookID,
		BookContentID: userStroke.BookContentID,
		ImageUrl:      userStroke.ImageUrl,
		Strokes:       userStroke.Strokes,
	}, nil
}

func StoreUserStroke(ctx context.Context, params contract.StoreUserStroke) error {
	var err error

	if params.AppSession == "" {
		err = fmt.Errorf("missing app session")
		return err
	}

	var userID sql.NullInt64
	if params.UserGuid != "" {
		user, err := user_repo.GetByGuid(ctx, params.UserGuid)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}
		userID = sql.NullInt64{user.ID, true}
	}

	userStroke := model.UserStroke{
		UserID:        userID,
		AppSession:    sql.NullString{params.AppSession, true},
		BookID:        sql.NullInt64{params.BookID, true},
		BookContentID: sql.NullInt64{params.BookContentID, true},
		ImageUrl:      params.ImageUrl,
		Strokes:       params.Strokes,
	}
	_, err = user_stroke_repo.Upsert(ctx, nil, userStroke)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
