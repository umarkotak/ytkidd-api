package file_bucket_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"fb.id",
		"fb.created_at",
		"fb.updated_at",
		"fb.deleted_at",
		"fb.guid",
		"fb.name",
		"fb.base_type",
		"fb.extension",
		"fb.http_content_type",
		"fb.metadata",
		"fb.exact_path",
	}, ", ")

	allColumnsWithData = strings.Join([]string{
		"fb.id",
		"fb.created_at",
		"fb.updated_at",
		"fb.deleted_at",
		"fb.guid",
		"fb.name",
		"fb.base_type",
		"fb.extension",
		"fb.http_content_type",
		"fb.metadata",
		"fb.exact_path",
		"fb.data",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM file_bucket fb
		WHERE
			fb.id = :id
			AND fb.deleted_at IS NULL
	`, allColumnsWithData)

	queryGetByGuid = fmt.Sprintf(`
		SELECT
			%s
		FROM file_bucket fb
		WHERE
			fb.guid = :guid
			AND fb.deleted_at IS NULL
	`, allColumnsWithData)

	queryInsert = `
		INSERT INTO file_bucket (
			guid,
			name,
			base_type,
			extension,
			http_content_type,
			metadata,
			data,
			exact_path
		) VALUES (
			:guid,
			:name,
			:base_type,
			:extension,
			:http_content_type,
			:metadata,
			:data,
			:exact_path
		) RETURNING id
	`

	querySoftDelete = `
		UPDATE file_bucket
		SET
			deleted_at = NOW()
		WHERE
			id = :id
	`

	queryDelete = `
		DELETE FROM file_bucket
		WHERE
			id = :id
	`
)

var (
	stmtGetByID    *sqlx.NamedStmt
	stmtGetByGuid  *sqlx.NamedStmt
	stmtInsert     *sqlx.NamedStmt
	stmtSoftDelete *sqlx.NamedStmt
	stmtDelete     *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByID, err = datastore.Get().Db.PrepareNamed(queryGetByID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByGuid, err = datastore.Get().Db.PrepareNamed(queryGetByGuid)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtInsert, err = datastore.Get().Db.PrepareNamed(queryInsert)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtSoftDelete, err = datastore.Get().Db.PrepareNamed(querySoftDelete)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtDelete, err = datastore.Get().Db.PrepareNamed(queryDelete)
	if err != nil {
		logrus.Fatal(err)
	}
}
