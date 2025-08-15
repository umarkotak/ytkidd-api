package order_service

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/order_repo"
	"github.com/umarkotak/ytkidd-api/repos/product_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/utils/payment_lib"
)

func CreateOrder(ctx context.Context, params contract.CreateOrder) (resp_contract.CreateOrder, error) {
	user, err := user_repo.GetByGuid(ctx, params.UserGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.CreateOrder{}, err
	}

	product, err := product_repo.GetByCode(ctx, params.ProductCode)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.CreateOrder{}, err
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
		return resp_contract.CreateOrder{}, err
	}

	order.PaymentNumber = sql.NullString{createPaymentData.Number, true}

	order.ID, err = order_repo.Insert(ctx, nil, order)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.CreateOrder{}, err
	}

	return resp_contract.CreateOrder{
		OrderNumber:                order.Number,
		MidtransSnapToken:          createPaymentData.MidtransSnapToken,
		MidtransSnapUrl:            createPaymentData.MidtransSnapUrl,
		MidtransSandboxPaymentPage: "https://simulator.sandbox.midtrans.com/v2/qris/index",
		MidtransNotificationPage:   "https://dashboard.sandbox.midtrans.com/settings/vtweb_configuration/history",
	}, nil
}

func GetOrderDetail(ctx context.Context, userGuid, orderNumber string) (resp_contract.OrderDetail, error) {
	user, err := user_repo.GetByGuid(ctx, userGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.OrderDetail{}, err
	}

	order, err := order_repo.GetByNumber(ctx, orderNumber)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.OrderDetail{}, err
	}

	if order.UserID != user.ID {
		return resp_contract.OrderDetail{}, model.ErrForbidden
	}

	return resp_contract.OrderDetail{
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
	}, nil
}

func CheckOrderPayment(ctx context.Context, userGuid, orderNumber string) (resp_contract.CheckOrderPayment, error) {
	user, err := user_repo.GetByGuid(ctx, userGuid)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.CheckOrderPayment{}, err
	}

	order, err := order_repo.GetByNumber(ctx, orderNumber)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.CheckOrderPayment{}, err
	}

	if order.UserID != user.ID {
		return resp_contract.CheckOrderPayment{}, model.ErrForbidden
	}

	return resp_contract.CheckOrderPayment{
		OrderNumber:   order.Number,
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
	}, nil
}

func GetOrderList(ctx context.Context, params contract.GetOrderByParams) (resp_contract.OrderList, error) {
	orders, err := order_repo.GetByParams(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.OrderList{}, err
	}

	ordersListData := make([]resp_contract.OrderListData, 0, len(orders))
	for _, order := range orders {
		ordersListData = append(ordersListData, resp_contract.OrderListData{
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
		})
	}

	return resp_contract.OrderList{
		Orders: ordersListData,
	}, nil
}
