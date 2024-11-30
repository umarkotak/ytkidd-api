package model

type (
	Pagination struct {
		Limit    int64 `json:"limit" db:"limit"`
		Page     int64 `json:"page" db:"page"`
		Offset   int64 `json:"-" db:"offset"`
		Total    int64 `json:"total" db:"total"`
		NextPage bool  `json:"next_page" db:"next_page"`
	}
)

func (m *Pagination) SetDefault() {
	if m.Limit <= 0 {
		m.Limit = 50
	}

	if m.Page <= 0 {
		m.Page = 1
	}

	m.Offset = (m.Page - 1) * m.Limit
}

func (m *Pagination) SetIsNextPage() {
	offset := (m.Page - 1) * m.Limit
	if offset+m.Limit < m.Total {
		m.NextPage = true
	} else {
		m.NextPage = false
	}
}
