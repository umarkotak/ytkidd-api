package user_subscription_repo

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
		"us.order_id",
		"us.product_code",
		"us.started_at",
		"us.ended_at",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM user_subscriptions us
		WHERE
			us.id = :id
			AND us.deleted_at IS NULL
	`, allColumns)

	queryGetByUserID = fmt.Sprintf(`
		SELECT
			%s
		FROM user_subscriptions us
		WHERE
			us.user_id = :user_id
			AND us.deleted_at IS NULL
		ORDER BY us.id DESC
		LIMIT :limit OFFSET :offset
	`, allColumns)

	queryGetActiveByUserID = fmt.Sprintf(`
		SELECT
			%s
		FROM user_subscriptions us
		WHERE
			us.user_id = :user_id
			AND NOW() BETWEEN us.started_at AND us.ended_at
			AND us.deleted_at IS NULL
		ORDER BY us.id DESC
		LIMIT :limit OFFSET :offset
	`, allColumns)

	queryInsert = `
		INSERT INTO user_subscriptions (
			user_id,
			order_id,
			product_code,
			started_at,
			ended_at
		) VALUES (
			:user_id,
			:order_id,
			:product_code,
			:started_at,
			:ended_at
		) RETURNING id
	`
)

var (
	stmtGetByID           *sqlx.NamedStmt
	stmtGetByUserID       *sqlx.NamedStmt
	stmtGetActiveByUserID *sqlx.NamedStmt
	stmtInsert            *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByID, err = datastore.Get().Db.PrepareNamed(queryGetByID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByUserID, err = datastore.Get().Db.PrepareNamed(queryGetByUserID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetActiveByUserID, err = datastore.Get().Db.PrepareNamed(queryGetActiveByUserID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtInsert, err = datastore.Get().Db.PrepareNamed(queryInsert)
	if err != nil {
		logrus.Fatal(err)
	}
}
