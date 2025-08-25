package order_service

import (
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/order_repo"
	"github.com/umarkotak/ytkidd-api/repos/product_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/utils/payment_lib"
)

func CreateOrder(ctx context.Context, params contract.CreateOrder) (contract_resp.CreateOrder, error) {
	user, err := user_repo.GetByGuid(ctx, params.UserGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.CreateOrder{}, err
	}

	product, err := product_repo.GetByCode(ctx, params.ProductCode)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.CreateOrder{}, err
	}

	order := model.Order{
		UserID:         user.ID,
		OrderType:      product.BenefitType,
		Description:    product.Name,
		Status:         model.ORDER_STATUS_INITIALIZED,
		PaymentStatus:  payment_lib.STATUS_INITIALIZED,
		BasePrice:      product.BasePrice,
		Price:          product.Price,
		DiscountAmount: 0,
		FinalPrice:     product.Price,
		PaymentNumber:  sql.NullString{},
		Metadata: model.OrderMetadata{
			ProductCode: product.Code,
		},
	}
	order.GenNumber()

	createPaymentData, err := payment_lib.CreatePayment(ctx, nil, payment_lib.CreatePaymentParams{
		PaymentPlatform: payment_lib.PAYMENT_PLATFORM_MIDTRANS,
		PaymentType:     payment_lib.PAYMENT_TYPE_MIDTRANS,
		OrderNumber:     order.Number,
		OrderItems: []payment_lib.OrderItem{
			{
				ID:       product.Code,
				Price:    product.Price,
				Quantity: 1,
				Name:     product.Name,
			},
		},
		CustomerFirstName: user.Name,
		CustomerEmail:     user.Email,
		Amount:            order.FinalPrice,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.CreateOrder{}, err
	}

	order.PaymentNumber = sql.NullString{createPaymentData.Number, true}

	order.ID, err = order_repo.Insert(ctx, nil, order)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.CreateOrder{}, err
	}

	return contract_resp.CreateOrder{
		OrderNumber:                order.Number,
		MidtransSnapToken:          createPaymentData.MidtransSnapToken,
		MidtransSnapUrl:            createPaymentData.MidtransSnapUrl,
		MidtransSandboxPaymentPage: "https://simulator.sandbox.midtrans.com/v2/qris/index",
		MidtransNotificationPage:   "https://dashboard.sandbox.midtrans.com/settings/vtweb_configuration/history",
	}, nil
}

func GetOrderDetail(ctx context.Context, userGuid, orderNumber string) (contract_resp.OrderDetail, error) {
	user, err := user_repo.GetByGuid(ctx, userGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.OrderDetail{}, err
	}

	order, err := order_repo.GetByNumber(ctx, orderNumber)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.OrderDetail{}, err
	}

	payment, err := payment_lib.GetByNumber(ctx, order.PaymentNumber.String)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.OrderDetail{}, err
	}

	if order.UserID != user.ID {
		return contract_resp.OrderDetail{}, model.ErrForbidden
	}

	if order.Status != model.ORDER_STATUS_INITIALIZED || payment.ExpiredAt.Time.Before(time.Now()) {
		payment.Metadata = payment_lib.PaymentMetadata{}
	}

	return contract_resp.OrderDetail{
		CreatedAt:      order.CreatedAt,
		UpdatedAt:      order.UpdatedAt,
		UserID:         order.UserID,
		Number:         order.Number,
		OrderType:      order.OrderType,
		Description:    order.Description,
		Status:         order.Status,
		PaymentStatus:  order.PaymentStatus,
		BasePrice:      order.BasePrice,
		Price:          order.Price,
		DiscountAmount: order.DiscountAmount,
		FinalPrice:     order.FinalPrice,
		PaymentNumber:  order.PaymentNumber.String,
		Metadata:       order.Metadata,

		PaymentExpiredAt:       payment.ExpiredAt.Time,
		PaymentSuccessAt:       payment.SuccessAt.Time,
		PaymentPaymentPlatform: payment.PaymentPlatform,
		PaymentPaymentType:     payment.PaymentType,
		PaymentReferenceStatus: payment.ReferenceStatus.String,
		PaymentReferenceNumber: payment.ReferenceNumber.String,
		PaymentFraudStatus:     payment.FraudStatus.String,
		PaymentMaskedCard:      payment.MaskedCard.String,
		PaymentAmount:          payment.Amount,
		PaymentMetadata: model.PaymentMetadata{
			SnapToken: payment.Metadata.SnapToken,
			SnapUrl:   payment.Metadata.SnapUrl,
		},
	}, nil
}

func CheckOrderPayment(ctx context.Context, userGuid, orderNumber string) (contract_resp.CheckOrderPayment, error) {
	user, err := user_repo.GetByGuid(ctx, userGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.CheckOrderPayment{}, err
	}

	order, err := order_repo.GetByNumber(ctx, orderNumber)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.CheckOrderPayment{}, err
	}

	if order.UserID != user.ID {
		return contract_resp.CheckOrderPayment{}, model.ErrForbidden
	}

	return contract_resp.CheckOrderPayment{
		OrderNumber:   order.Number,
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
	}, nil
}

func GetOrderList(ctx context.Context, params contract.GetOrderByParams) (contract_resp.OrderList, error) {
	orders, err := order_repo.GetByParamsWithPayment(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.OrderList{}, err
	}

	ordersListData := make([]contract_resp.OrderListData, 0, len(orders))
	for _, order := range orders {
		if order.Status != model.ORDER_STATUS_INITIALIZED || order.PaymentExpiredAt.Time.Before(time.Now()) {
			order.PaymentMetadata = model.PaymentMetadata{}
		}

		ordersListData = append(ordersListData, contract_resp.OrderListData{
			CreatedAt:      order.CreatedAt,
			UpdatedAt:      order.UpdatedAt,
			UserID:         order.UserID,
			Number:         order.Number,
			OrderType:      order.OrderType,
			Description:    order.Description,
			Status:         order.Status,
			HumanStatus:    order.HumanStatus(),
			PaymentStatus:  order.PaymentStatus,
			BasePrice:      order.BasePrice,
			Price:          order.Price,
			DiscountAmount: order.DiscountAmount,
			FinalPrice:     order.FinalPrice,
			PaymentNumber:  order.PaymentNumber.String,
			Metadata:       order.Metadata,

			PaymentExpiredAt:       order.PaymentExpiredAt.Time,
			PaymentSuccessAt:       order.PaymentSuccessAt.Time,
			PaymentPaymentPlatform: order.PaymentPaymentPlatform,
			PaymentPaymentType:     order.PaymentPaymentType,
			PaymentReferenceStatus: order.PaymentReferenceStatus.String,
			PaymentReferenceNumber: order.PaymentReferenceNumber.String,
			PaymentFraudStatus:     order.PaymentFraudStatus.String,
			PaymentMaskedCard:      order.PaymentMaskedCard.String,
			PaymentAmount:          order.PaymentAmount,
			PaymentMetadata:        order.PaymentMetadata,
		})
	}

	return contract_resp.OrderList{
		Orders: ordersListData,
	}, nil
}
