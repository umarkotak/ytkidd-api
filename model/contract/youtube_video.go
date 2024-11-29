package contract

import (
	"github.com/lib/pq"
)

type (
	YoutubeVideoSearch struct {
		Name   string         `db:"name"`
		Tags   pq.StringArray `db:"tags"`
		Active bool           `db:"active"`
	}
)
