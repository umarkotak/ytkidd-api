package book_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"b.id",
		"b.created_at",
		"b.updated_at",
		"b.deleted_at",
		"b.title",
		"b.description",
		"b.cover_image_url",
		"b.tags",
		"b.type",
		"b.pdf_file_url",
		"b.active",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM books b
		WHERE
			b.id = :id
			AND b.deleted_at IS NULL
	`, allColumns)

	queryGetForSearch = fmt.Sprintf(`
		SELECT
			%s
		FROM books b
		WHERE
			1 = 1
			AND (:title = '' OR b.title = :title)
			AND (:tags = '{}' OR b.tags @> :tags)
			AND (b.active = :active)
			AND b.deleted_at IS NULL
	`, allColumns)

	queryInsert = `
		INSERT INTO books (
			title,
			description,
			cover_image_url,
			tags,
			type,
			pdf_file_url
		) VALUES (
			:title,
			:description,
			:cover_image_url,
			:tags,
			:type,
			:pdf_file_url
		) RETURNING id
	`

	queryUpdate = `
		UPDATE books
		SET
			title = :title,
			description = :description,
			cover_image_url = :cover_image_url,
			tags = :tags,
			type = :type,
			pdf_file_url = :pdf_file_url
		WHERE
			id = :id
	`

	querySoftDelete = `
		UPDATE books
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
	stmtGetByParams     *sqlx.NamedStmt
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
