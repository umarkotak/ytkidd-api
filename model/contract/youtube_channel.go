package contract

import (
	"github.com/lib/pq"
)

type (
	YoutubeChannelSearch struct {
		Name string         `db:"name"`
		Tags pq.StringArray `db:"tags"`
	}
)
