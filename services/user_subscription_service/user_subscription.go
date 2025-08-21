package user_subscription_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_subscription_repo"
	"github.com/umarkotak/ytkidd-api/utils"
)

func GetUserSubscriptionInfo(ctx context.Context, userGuid string) (contract_resp.UserSubscriptionInfo, error) {
	user, err := user_repo.GetByGuid(ctx, userGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.UserSubscriptionInfo{}, err
	}

	activeSubscriptions, err := user_subscription_repo.GetAllActiveByUserID(ctx, user.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.UserSubscriptionInfo{}, err
	}

	activeSubscriptionsData := make([]contract_resp.UserSubscription, 0, len(activeSubscriptions))
	for _, activeSubscription := range activeSubscriptions {
		activeSubscriptionsData = append(activeSubscriptionsData, contract_resp.UserSubscription{
			StartedAt: activeSubscription.StartedAt,
			EndedAt:   activeSubscription.EndedAt,
		})
	}

	if len(activeSubscriptionsData) <= 0 {
		return contract_resp.UserSubscriptionInfo{}, nil
	}

	endedAt := activeSubscriptionsData[len(activeSubscriptionsData)-1].EndedAt
	return contract_resp.UserSubscriptionInfo{
		Active:              true,
		EndedAt:             endedAt,
		RemainingDay:        utils.RemainingDays(endedAt),
		ActiveSubscriptions: activeSubscriptionsData,
	}, nil
}
