package contract

import (
	"github.com/lib/pq"
	"github.com/umarkotak/ytkidd-api/model"
)

type (
	YoutubeVideoSearch struct {
		Name   string         `db:"name"`
		Tags   pq.StringArray `db:"tags"`
		Active bool           `db:"active"`
	}

	GetYoutubeVideosHome struct {
		Tags       pq.StringArray `db:"tags"`
		ExcludeIDs pq.Int64Array  `db:"exclude_ids"`
		model.Pagination
	}
)
