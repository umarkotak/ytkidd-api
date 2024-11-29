package contract

import "time"

type (
	CreateDatingPlan struct {
		FromUserGuid string    `json:"-"`
		ToUserGuid   string    `json:"to_user_guid"`
		PlannedAt    time.Time `json:"planned_at"`
	}

	DatingFeedback struct {
		UserGuid   string `json:"-"`
		UserPlanID int64  `json:"-"`
		Message    string `json:"message"`
	}
)
