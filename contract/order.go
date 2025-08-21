package contract

import "github.com/umarkotak/ytkidd-api/model"

type (
	CreateOrder struct {
		UserGuid    string `json:"-"`
		ProductCode string `json:"product_code"`
	}

	GetOrderByUserID struct {
		UserID int64 `db:"user_id"`
		model.Pagination
	}

	GetOrderByParams struct {
		UserID int64 `db:"user_id"`
		model.Pagination
	}
)
