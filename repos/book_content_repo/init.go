package book_content_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"bc.id",
		"bc.created_at",
		"bc.updated_at",
		"bc.deleted_at",
		"bc.book_id",
		"bc.page_number",
		"bc.image_file_guid",
		"bc.description",
		"bc.metadata",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM book_contents bc
		WHERE
			bc.id = :id
			AND bc.deleted_at IS NULL
	`, allColumns)

	queryGetByBookID = fmt.Sprintf(`
		SELECT
			%s
		FROM book_contents bc
		WHERE
			bc.book_id = :book_id
			AND bc.deleted_at IS NULL
	`, allColumns)

	queryInsert = `
		INSERT INTO book_contents (
			book_id,
			page_number,
			image_file_guid,
			description,
			metadata
		) VALUES (
			:book_id,
			:page_number,
			:image_file_guid,
			:description,
			:metadata
		) RETURNING id
	`

	queryUpdate = `
		UPDATE book_contents
		SET
			book_id = :book_id,
			page_number = :page_number,
			image_file_guid = :image_file_guid,
			description = :description,
			metadata = :metadata
		WHERE
			id = :id
	`

	querySoftDelete = `
		UPDATE book_contents
		SET
			deleted_at = NOW()
		WHERE
			id = :id
	`
)

var (
	stmtGetByID     *sqlx.NamedStmt
	stmtGetByBookID *sqlx.NamedStmt
	stmtInsert      *sqlx.NamedStmt
	stmtUpdate      *sqlx.NamedStmt
	stmtSoftDelete  *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByID, err = datastore.Get().Db.PrepareNamed(queryGetByID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByBookID, err = datastore.Get().Db.PrepareNamed(queryGetByBookID)
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
