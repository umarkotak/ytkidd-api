package user_service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/utils/random"
	"github.com/umarkotak/ytkidd-api/utils/user_auth"
)

func SignUp(ctx context.Context, params contract.UserSignUp) (resp_contract.UserSignUp, error) {
	user, err := user_repo.GetByEmail(ctx, params.Email)
	if err != nil && err != sql.ErrNoRows {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSignUp{}, err
	}

	if user.ID != 0 {
		err = fmt.Errorf("user already exists")
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSignUp{}, err
	}

	user = model.User{
		Email:    params.Email,
		About:    params.About,
		Password: params.Password,
		Name:     params.Name,
		Username: params.Username,
	}

	user.ID, err = user_repo.Insert(ctx, nil, user)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSignUp{}, err
	}

	return resp_contract.UserSignUp{}, nil
}

func SignIn(ctx context.Context, params contract.UserSignIn) (resp_contract.UserSignIn, error) {
	user, err := user_repo.GetByEmail(ctx, params.Email)
	if err != nil && err != sql.ErrNoRows {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSignIn{}, err
	}

	if user.ID == 0 {
		err = fmt.Errorf("invalid email or password")
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSignIn{}, err
	}

	if user.Password != params.Password {
		err = fmt.Errorf("invalid email or password")
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSignIn{}, err
	}

	authPayload := user_auth.Payload{
		JWTID:          random.MustGenUUIDTimes(2),
		SID:            random.MustGenUUIDTimes(2),
		Issuer:         "cookiekid-be",
		IssuedAt:       time.Now().Unix(),
		ExpirationTime: time.Now().Add(model.USER_AUTH_EXPIRY).Unix(),
		GUID:           user.Guid,
		Name:           user.Name,
		Username:       user.Username,
		Email:          user.Email,
	}
	accessToken, err := user_auth.GenToken(ctx, authPayload)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSignIn{}, err
	}

	return resp_contract.UserSignIn{
		AccessToken: accessToken,
	}, nil
}
