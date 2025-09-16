package user_activity_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"ua.id",
		"ua.created_at",
		"ua.updated_at",
		"ua.deleted_at",
		"ua.user_id",
		"ua.app_session",
		"ua.youtube_video_id",
		"ua.book_id",
		"ua.book_content_id",
		"ua.metadata",
	}, ", ")

	queryGetByParams = fmt.Sprintf(`
		SELECT
			%s
		FROM user_activities ua
		WHERE
			1 = 1
			AND ua.user_id = :user_id
			AND ua.app_session = :app_session
			AND ua.deleted_at IS NULL
		ORDER BY ua.updated_at DESC
	`, allColumns)

	queryGetFullByParams = `
		SELECT
			ua.id,
			ua.created_at,
			ua.updated_at,
			ua.deleted_at,
			ua.user_id,
			ua.app_session,
			ua.youtube_video_id,
			ua.book_id,
			ua.book_content_id,
			ua.metadata
		FROM user_activities ua
		LEFT JOIN books b ON b.id = ua.book_id
		LEFT JOIN book_contents bc on bc.id = ua.book_content_id
		LEFT JOIN youtube_videos yv ON yv.id = ua.youtube_video_id
		WHERE
			1 = 1
			AND ua.user_id = :user_id
			AND ua.app_session = :app_session
			AND ua.deleted_at IS NULL
		ORDER BY ua.updated_at DESC
	`

	queryGetByUserActivity = fmt.Sprintf(`
		SELECT
			%s
		FROM user_activities ua
		WHERE
			1 = 1
			AND ua.user_id = :user_id
			AND ua.app_session = :app_session
			AND ua.youtube_video_id = :youtube_video_id
			AND ua.book_id = :book_id
			AND ua.book_content_id = :book_content_id
			AND ua.deleted_at IS NULL
	`, allColumns)

	queryGetFullByUserActivity = `
		SELECT
			ua.id,
			ua.created_at,
			ua.updated_at,
			ua.deleted_at,
			ua.user_id,
			ua.app_session,
			ua.youtube_video_id,
			ua.book_id,
			ua.book_content_id,
			ua.metadata
		FROM user_activities ua
		LEFT JOIN books b ON b.id = ua.book_id
		LEFT JOIN book_contents bc on bc.id = ua.book_content_id
		LEFT JOIN youtube_videos yv ON yv.id = ua.youtube_video_id
		WHERE
			1 = 1
			AND ua.user_id = :user_id
			AND ua.app_session = :app_session
			AND ua.youtube_video_id = :youtube_video_id
			AND ua.book_id = :book_id
			AND ua.book_content_id = :book_content_id
			AND ua.deleted_at IS NULL
	`

	queryInsert = `
		INSERT INTO user_activities (
			user_id,
			app_session,
			youtube_video_id,
			book_id,
			book_content_id,
			metadata
		) VALUES (
			:user_id,
			:app_session,
			:youtube_video_id,
			:book_id,
			:book_content_id,
			:metadata
		) RETURNING id
	`

	queryUpsert = `
		INSERT INTO user_activities (
			user_id,
			app_session,
			youtube_video_id,
			book_id,
			book_content_id,
			metadata
		) VALUES (
			:user_id,
			:app_session,
			:youtube_video_id,
			:book_id,
			:book_content_id,
			:metadata
		)
		ON CONFLICT (user_id, app_session, youtube_video_id, book_id, book_content_id)
		DO UPDATE SET
			updated_at = NOW(),
			metadata = :metadata
		RETURNING id
	`
)

var (
	stmtGetByParams           *sqlx.NamedStmt
	stmtGetFullByParams       *sqlx.NamedStmt
	stmtGetByUserActivity     *sqlx.NamedStmt
	stmtGetFullByUserActivity *sqlx.NamedStmt
	stmtInsert                *sqlx.NamedStmt
	stmtUpsert                *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByParams, err = datastore.Get().Db.PrepareNamed(queryGetByParams)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetFullByParams, err = datastore.Get().Db.PrepareNamed(queryGetFullByParams)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByUserActivity, err = datastore.Get().Db.PrepareNamed(queryGetByUserActivity)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetFullByUserActivity, err = datastore.Get().Db.PrepareNamed(queryGetFullByUserActivity)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtInsert, err = datastore.Get().Db.PrepareNamed(queryInsert)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtUpsert, err = datastore.Get().Db.PrepareNamed(queryUpsert)
	if err != nil {
		logrus.Fatal(err)
	}
}
