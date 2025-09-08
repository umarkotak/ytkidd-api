package user_stroke_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"us.id",
		"us.created_at",
		"us.updated_at",
		"us.deleted_at",
		"us.user_id",
		"us.app_session",
		"us.book_id",
		"us.book_content_id",
		"us.image_url",
		"us.strokes",
	}, ", ")

	queryGetByUserAndContent = fmt.Sprintf(`
		SELECT
			%s
		FROM user_strokes us
		WHERE
			us.user_id = :user_id,
			us.app_session = :app_session,
			us.book_content_id = :book_content_id,
			AND us.deleted_at IS NULL
	`, allColumns)

	queryInsert = `
		INSERT INTO user_strokes (
			user_id,
			app_session,
			book_id,
			book_content_id,
			image_url,
			strokes
		) VALUES (
			:user_id,
			:app_session,
			:book_id,
			:book_content_id,
			:image_url,
			:strokes
		) RETURNING id
	`

	queryUpsert = `
		INSERT INTO user_strokes (
			user_id,
			app_session,
			book_id,
			book_content_id,
			image_url,
			strokes
		) VALUES (
			:user_id,
			:app_session,
			:book_id,
			:book_content_id,
			:image_url,
			:strokes
		) RETURNING id
		ON CONFLICT (user_id, app_session, book_content_id)
		DO UPDATE SET
			updated_at = NOW(),
			strokes = :strokes
	`
)

var (
	stmtGetByUserAndContent *sqlx.NamedStmt
	stmtInsert              *sqlx.NamedStmt
	stmtUpsert              *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByUserAndContent, err = datastore.Get().Db.PrepareNamed(queryGetByUserAndContent)
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
