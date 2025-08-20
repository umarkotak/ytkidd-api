package payment_lib

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
)

type PaymentConf struct {
	IsLive             bool
	MidtransMerchantID string
	MidtransClientKey  string
	MidtransServerKey  string
}

var (
	allColumns = strings.Join([]string{
		"py.id",
		"py.created_at",
		"py.updated_at",
		"py.deleted_at",
		"py.expired_at",
		"py.success_at",
		"py.order_number",
		"py.number",
		"py.payment_platform",
		"py.payment_type",
		"py.status",
		"py.reference_status",
		"py.reference_number",
		"py.fraud_status",
		"py.masked_card",
		"py.amount",
		"py.metadata",
	}, ", ")

	queryGetByID = fmt.Sprintf(`
		SELECT
			%s
		FROM payments py
		WHERE
			py.id = :id
			AND py.deleted_at IS NULL
	`, allColumns)

	queryGetByNumber = fmt.Sprintf(`
		SELECT
			%s
		FROM payments py
		WHERE
			py.number = :number
			AND py.deleted_at IS NULL
	`, allColumns)

	queryGetByOrderNumber = fmt.Sprintf(`
		SELECT
			%s
		FROM payments py
		WHERE
			py.order_number = :order_number
			AND py.deleted_at IS NULL
	`, allColumns)

	queryInsert = `
		INSERT INTO payments (
			expired_at,
			success_at,
			order_number,
			number,
			payment_platform,
			payment_type,
			status,
			reference_status,
			reference_number,
			fraud_status,
			masked_card,
			amount,
			metadata
		) VALUES (
			:expired_at,
			:success_at,
			:order_number,
			:number,
			:payment_platform,
			:payment_type,
			:status,
			:reference_status,
			:reference_number,
			:fraud_status,
			:masked_card,
			:amount,
			:metadata
		) RETURNING id
	`

	queryUpdate = `
		UPDATE payments
		SET
			updated_at = NOW(),
			expired_at = :expired_at,
			success_at = :success_at,
			order_number = :order_number,
			number = :number,
			payment_platform = :payment_platform,
			payment_type = :payment_type,
			status = :status,
			reference_status = :reference_status,
			reference_number = :reference_number,
			fraud_status = :fraud_status,
			masked_card = :masked_card,
			amount = :amount,
			metadata = :metadata
		WHERE
			id = :id
	`
)

var (
	stmtGetByID          *sqlx.NamedStmt
	stmtGetByNumber      *sqlx.NamedStmt
	stmtGetByOrderNumber *sqlx.NamedStmt
	stmtInsert           *sqlx.NamedStmt
	stmtUpdate           *sqlx.NamedStmt

	midtransSnapClient snap.Client
)

func Initialize(db *sqlx.DB, conf PaymentConf) {
	var err error

	stmtGetByID, err = db.PrepareNamed(queryGetByID)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByNumber, err = db.PrepareNamed(queryGetByNumber)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtGetByOrderNumber, err = db.PrepareNamed(queryGetByOrderNumber)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtInsert, err = db.PrepareNamed(queryInsert)
	if err != nil {
		logrus.Fatal(err)
	}
	stmtUpdate, err = db.PrepareNamed(queryUpdate)
	if err != nil {
		logrus.Fatal(err)
	}

	if conf.MidtransMerchantID != "" {
		if conf.IsLive {
			midtrans.Environment = midtrans.Production
		} else {
			midtrans.Environment = midtrans.Sandbox
		}

		midtrans.ServerKey = conf.MidtransServerKey
		midtransSnapClient.New(conf.MidtransServerKey, midtrans.Environment)
	}
}
