package order_repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
)

var (
	allColumns = strings.Join([]string{
		"o.id",
		"o.created_at",
		"o.updated_at",
		"o.deleted_at",
		"o.user_id",
		"o.number",
		"o.order_type",
		"o.description",
		"o.status",
		"o.payment_status",
		"o.base_price",
		"o.price",
		"o.discount_amount",
		"o.final_price",
		"o.payment_number",
		"o.metadata",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM orders o
		WHERE
			o.id = :id
			AND o.deleted_at IS NULL
	`, allColumns)

	queryGetByNumber = fmt.Sprintf(`
		SELECT
			%s
		FROM orders o
		WHERE
			o.number = :number
			AND o.deleted_at IS NULL
	`, allColumns)

	queryGetByUserID = fmt.Sprintf(`
		SELECT
			%s
		FROM orders o
		WHERE
			o.user_id = :user_id
			AND o.deleted_at IS NULL
		ORDER BY o.id DESC
		LIMIT :limit OFFSET :offset
	`, allColumns)

	queryGetByParams = fmt.Sprintf(`
		SELECT
			%s
		FROM orders o
		WHERE
			(:user_id = 0 OR o.user_id = :user_id)
			AND o.deleted_at IS NULL
		ORDER BY o.id DESC
		LIMIT :limit OFFSET :offset
	`, allColumns)

	queryGetByParamsWithPayment = fmt.Sprintf(`
		SELECT
			%s,
			p.expired_at AS payment_expired_at,
			p.success_at AS payment_success_at,
			p.payment_platform AS payment_payment_platform,
			p.payment_type AS payment_payment_type,
			p.reference_status AS payment_reference_status,
			p.reference_number AS payment_reference_number,
			p.fraud_status AS payment_fraud_status,
			p.masked_card AS payment_masked_card,
			p.amount AS payment_amount,
			p.metadata AS payment_metadata
		FROM orders o
		INNER JOIN payments p ON o.payment_number = p.number
		WHERE
			(:user_id = 0 OR o.user_id = :user_id)
			AND o.deleted_at IS NULL
		ORDER BY o.id DESC
		LIMIT :limit OFFSET :offset
	`, allColumns)

	queryInsert = `
		INSERT INTO orders (
			user_id,
			number,
			order_type,
			description,
			status,
			payment_status,
			base_price,
			price,
			discount_amount,
			final_price,
			payment_number,
			metadata
		) VALUES (
			:user_id,
			:number,
			:order_type,
			:description,
			:status,
			:payment_status,
			:base_price,
			:price,
			:discount_amount,
			:final_price,
			:payment_number,
			:metadata
		) RETURNING id
	`

	queryUpdate = `
		UPDATE orders SET
			updated_at = NOW(),
			user_id = :user_id,
			number = :number,
			order_type = :order_type,
			description = :description,
			status = :status,
			payment_status = :payment_status,
			base_price = :base_price,
			price = :price,
			discount_amount = :discount_amount,
			final_price = :final_price,
			payment_number = :payment_number,
			metadata = :metadata
		WHERE
			id = :id
	`
)

var (
	stmtGetByID                *sqlx.NamedStmt
	stmtGetByNumber            *sqlx.NamedStmt
	stmtInsert                 *sqlx.NamedStmt
	stmtUpdate                 *sqlx.NamedStmt
	stmtGetByUserID            *sqlx.NamedStmt
	stmtGetByParams            *sqlx.NamedStmt
	stmtGetByParamsWithPayment *sqlx.NamedStmt
)

func Initialize() {
	var err error

	stmtGetByID, err = datastore.Get().Db.PrepareNamed(queryGetByID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByNumber, err = datastore.Get().Db.PrepareNamed(queryGetByNumber)
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
	stmtGetByUserID, err = datastore.Get().Db.PrepareNamed(queryGetByUserID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByParams, err = datastore.Get().Db.PrepareNamed(queryGetByParams)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByParamsWithPayment, err = datastore.Get().Db.PrepareNamed(queryGetByParamsWithPayment)
	if err != nil {
		logrus.Fatal(err)
	}
}
