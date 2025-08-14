package product_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"p.id",
		"p.created_at",
		"p.updated_at",
		"p.deleted_at",
		"p.code",
		"p.benefit_type",
		"p.name",
		"p.image_url",
		"p.base_price",
		"p.price",
		"p.metadata",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM products p
		WHERE
			p.id = :id
			AND p.deleted_at IS NULL
	`, allColumns)

	queryGetByCode = fmt.Sprintf(`
		SELECT
			%s
		FROM products p
		WHERE
			p.code = :code
			AND p.deleted_at IS NULL
	`, allColumns)
)

var (
	stmtGetByID   *sqlx.NamedStmt
	stmtGetByCode *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByID, err = datastore.Get().Db.PrepareNamed(queryGetByID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByCode, err = datastore.Get().Db.PrepareNamed(queryGetByCode)
	if err != nil {
		logrus.Fatal(err)
	}
}
