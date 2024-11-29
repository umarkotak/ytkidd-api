package youtube_video_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"ytvid.id",
		"ytvid.created_at",
		"ytvid.updated_at",
		"ytvid.deleted_at",
		"ytvid.youtube_channel_id",
		"ytvid.external_id",
		"ytvid.title",
		"ytvid.image_url",
		"ytvid.tags",
		"ytvid.active",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM youtube_videos ytvid
		WHERE
			ytvid.id = :id
			AND ytvid.deleted_at IS NULL
	`, allColumns)

	queryGetByExternalID = fmt.Sprintf(`
		SELECT
			%s
		FROM youtube_videos ytvid
		WHERE
			ytvid.external_id = :external_id
			AND ytvid.deleted_at IS NULL
	`, allColumns)

	queryGetForSearch = fmt.Sprintf(`
		SELECT
			%s
		FROM youtube_videos ytvid
		WHERE
			1 = 1
			AND (:title = '' OR ytvid.title = :title)
			AND (:tags = '{}' OR ytvid.tags @> :tags)
			AND (ytvid.active = :active)
			AND ytvid.deleted_at IS NULL
	`, allColumns)

	queryInsert = `
		INSERT INTO youtube_videos (
			youtube_channel_id,
			external_id,
			title,
			image_url,
			tags,
			active
		) VALUES (
			:youtube_channel_id,
			:external_id,
			:title,
			:image_url,
			:tags,
			:active
		) RETURNING id
	`

	queryUpdate = `
		UPDATE youtube_videos
		SET
			title = :title,
			image_url = :image_url,
			tags = :tags,
			active = :active
		WHERE
			id = :id
	`

	querySoftDelete = `
		UPDATE youtube_videos
		SET
			deleted_at = NOW()
		WHERE
			id = :id
	`
)

var (
	stmtGetByID         *sqlx.NamedStmt
	stmtGetByExternalID *sqlx.NamedStmt
	stmtGetForSearch    *sqlx.NamedStmt
	stmtInsert          *sqlx.NamedStmt
	stmtUpdate          *sqlx.NamedStmt
	stmtSoftDelete      *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByID, err = datastore.Get().Db.PrepareNamed(queryGetByID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByExternalID, err = datastore.Get().Db.PrepareNamed(queryGetByExternalID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetForSearch, err = datastore.Get().Db.PrepareNamed(queryGetForSearch)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtInsert, err = datastore.Get().Db.PrepareNamed(queryInsert)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtUpdate, err = datastore.Get().Db.PrepareNamed(queryUpdate)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtSoftDelete, err = datastore.Get().Db.PrepareNamed(querySoftDelete)
	if err != nil {
		logrus.Fatal(err)
	}
}
