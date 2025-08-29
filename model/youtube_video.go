package model

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type (
	YoutubeVideo struct {
		ID               int64          `db:"id"`
		CreatedAt        time.Time      `db:"created_at"`
		UpdatedAt        time.Time      `db:"updated_at"`
		DeletedAt        sql.NullTime   `db:"deleted_at"`
		YoutubeChannelID int64          `db:"youtube_channel_id"`
		ExternalId       string         `db:"external_id"`
		Title            string         `db:"title"`
		ImageUrl         string         `db:"image_url"`
		Tags             pq.StringArray `db:"tags"`
		Active           bool           `db:"active"`
		PublishedAt      time.Time      `db:"published_at"`
	}

	YoutubeVideoDetailed struct {
		ID                       int64          `db:"id"`
		ExternalId               string         `db:"external_id"`
		Title                    string         `db:"title"`
		ImageUrl                 string         `db:"image_url"`
		Tags                     pq.StringArray `db:"tags"`
		Active                   bool           `db:"active"`
		YoutubeChannelID         int64          `db:"youtube_channel_id"`
		YoutubeChannelExternalID string         `db:"youtube_channel_external_id"`
		YoutubeChannelName       string         `db:"youtube_channel_name"`
		YoutubeChannelUsername   string         `db:"youtube_channel_username"`
		YoutubeChannelImageUrl   string         `db:"youtube_channel_image_url"`
		YoutubeChannelTags       pq.StringArray `db:"youtube_channel_tags"`
		YoutubeChannelActive     bool           `db:"youtube_channel_active"`
	}
)
