package model

type (
	Pagination struct {
		Limit    int64 `json:"limit"`
		Page     int64 `json:"page"`
		Total    int64 `json:"total"`
		NextPage bool  `json:"next_page"`
	}
)

func (m *Pagination) SetIsNextPage() {
	offset := (m.Page - 1) * m.Limit
	if offset+m.Limit < m.Total {
		m.NextPage = true
	} else {
		m.NextPage = false
	}
}
