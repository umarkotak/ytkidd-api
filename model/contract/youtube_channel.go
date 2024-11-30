package contract

import (
	"github.com/lib/pq"
)

type (
	GetYoutubeChannels struct {
		Name string         `db:"name"`
		Tags pq.StringArray `db:"tags"`
	}
)
