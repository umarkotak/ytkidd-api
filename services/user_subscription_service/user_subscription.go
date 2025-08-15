package user_subscription_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_subscription_repo"
	"github.com/umarkotak/ytkidd-api/utils"
)

func GetUserSubscriptionInfo(ctx context.Context, userGuid string) (resp_contract.UserSubscriptionInfo, error) {
	user, err := user_repo.GetByGuid(ctx, userGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSubscriptionInfo{}, err
	}

	activeSubscriptions, err := user_subscription_repo.GetAllActiveByUserID(ctx, user.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.UserSubscriptionInfo{}, err
	}

	activeSubscriptionsData := make([]resp_contract.UserSubscription, 0, len(activeSubscriptions))
	for _, activeSubscription := range activeSubscriptions {
		activeSubscriptionsData = append(activeSubscriptionsData, resp_contract.UserSubscription{
			StartedAt: activeSubscription.StartedAt,
			EndedAt:   activeSubscription.EndedAt,
		})
	}

	if len(activeSubscriptionsData) <= 0 {
		return resp_contract.UserSubscriptionInfo{}, nil
	}

	endedAt := activeSubscriptionsData[len(activeSubscriptionsData)-1].EndedAt
	return resp_contract.UserSubscriptionInfo{
		Active:              true,
		EndedAt:             endedAt,
		RemainingDay:        utils.RemainingDays(endedAt),
		ActiveSubscriptions: activeSubscriptionsData,
	}, nil
}
