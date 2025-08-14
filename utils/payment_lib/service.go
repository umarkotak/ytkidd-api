package payment_lib

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
)

type (
	CreatePaymentParams struct {
		PaymentPlatform   string
		PaymentType       string
		OrderNumber       string
		OrderItems        []OrderItem
		CustomerFirstName string
		CustomerLastName  string
		CustomerEmail     string
		CustomerPhone     string
		Amount            int64
	}

	OrderItem struct {
		ID       string
		Price    int64
		Quantity int64
		Name     string
	}

	CreatePaymentData struct {
		Number            string
		MidtransSnapToken string
		MidtransSnapUrl   string
	}
)

func CreatePayment(ctx context.Context, tx *sqlx.Tx, params CreatePaymentParams) (CreatePaymentData, error) {
	var err error

	if params.PaymentPlatform != PAYMENT_PLATFORM_MIDTRANS {
		return CreatePaymentData{}, fmt.Errorf("payment platform not supported")
	}

	items := make([]midtrans.ItemDetails, 1, len(params.OrderItems))
	for _, orderItem := range params.OrderItems {
		items = append(items, midtrans.ItemDetails{
			ID:    orderItem.ID,
			Price: orderItem.Price,
			Qty:   int32(orderItem.Quantity),
			Name:  orderItem.Name,
		})
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  params.OrderNumber,
			GrossAmt: params.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: params.CustomerFirstName,
			LName: params.CustomerLastName,
			Email: params.CustomerEmail,
			Phone: params.CustomerPhone,
		},
		Items: &items,
	}

	snapResp, midtransErr := midtransSnapClient.CreateTransaction(req)
	if midtransErr != nil {
		logrus.WithContext(ctx).Error(midtransErr)
		return CreatePaymentData{}, midtransErr
	}

	if len(snapResp.ErrorMessages) > 0 {
		return CreatePaymentData{}, fmt.Errorf("midtrans error message: %v", snapResp.ErrorMessages)
	}

	payment := Payment{
		OrderNumber:     params.OrderNumber,
		PaymentPlatform: params.PaymentPlatform,
		PaymentType:     params.PaymentType,
		Status:          STATUS_INITIALIZED,
		ReferenceStatus: sql.NullString{},
		ReferenceNumber: sql.NullString{},
		FraudStatus:     sql.NullString{},
		MaskedCard:      sql.NullString{},
		Amount:          params.Amount,
		Metadata:        PaymentMetadata{},
	}

	payment.ID, payment.Number, err = Insert(ctx, tx, payment)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return CreatePaymentData{}, err
	}

	return CreatePaymentData{
		Number:            payment.Number,
		MidtransSnapToken: snapResp.Token,
		MidtransSnapUrl:   snapResp.RedirectURL,
	}, nil
}
