package user_repo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/model"
)

func SetInitPassword(ctx context.Context, initPassData model.InitPasswordData) error {
	key := fmt.Sprintf("%s:PASSWORD_INITIALIZATION:%s", model.REDIS_PREFIX, initPassData.InitPasswordToken)

	initPassJsonByte, err := json.Marshal(initPassData)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	_, err = datastore.Get().Redis.Set(ctx, key, string(initPassJsonByte), model.INIT_PASS_VERIFICATION_EXPIRY).Result()
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}

func GetInitPasswordByToken(ctx context.Context, initPassToken string) (model.InitPasswordData, error) {
	initPassData := model.InitPasswordData{}

	key := fmt.Sprintf("%s:PASSWORD_INITIALIZATION:%s", model.REDIS_PREFIX, initPassToken)

	initPassJsonString, err := datastore.Get().Redis.Get(ctx, key).Result()
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return initPassData, err
	}

	err = json.Unmarshal([]byte(initPassJsonString), &initPassData)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return initPassData, err
	}

	return initPassData, nil
}

func RemoveInitPasswordByToken(ctx context.Context, initPassToken string) error {
	key := fmt.Sprintf("%s:PASSWORD_INITIALIZATION:%s", model.REDIS_PREFIX, initPassToken)

	_, err := datastore.Get().Redis.Del(ctx, key).Result()
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
