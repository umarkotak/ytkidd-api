package resp_contract

import (
	"time"
)

type (
	UserSubscriptionInfo struct {
		Active              bool               `json:"active"`
		EndedAt             time.Time          `json:"ended_at"`
		RemainingDay        int64              `json:"remaining_day"`
		ActiveSubscriptions []UserSubscription `json:"active_subscriptions"`
	}

	UserSubscription struct {
		StartedAt time.Time `json:"started_at"`
		EndedAt   time.Time `json:"ended_at"`
	}
)
