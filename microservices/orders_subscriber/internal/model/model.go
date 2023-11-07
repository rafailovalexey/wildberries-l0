package model

import (
	model "github.com/emptyhopes/orders/internal/model/orders"
	"time"
)

type OrderModelInterface interface {
	NewOrderModel(
		orderUid string,
		trackNumber string,
		entry string,
		deliveryUid string,
		paymentUid string,
		locale string,
		internalSignature string,
		customerId string,
		deliveryService string,
		shardkey string,
		smId int,
		dateCreated time.Time,
		oofShard string,
	) *model.OrderModel

	NewOrderDeliveryModel(
		deliveryUid string,
		name string,
		phone string,
		zip string,
		city string,
		address string,
		region string,
		email string,
	) *model.OrderDeliveryModel

	NewOrderPaymentModel(
		paymentUid string,
		transaction string,
		requestId string,
		currency string,
		provider string,
		amount int,
		paymentDt time.Time,
		bank string,
		deliveryCost int,
		goodsTotal int,
		customFee int,
	) *model.OrderPaymentModel

	NewOrderItemModel(
		chrtId int,
		trackNumber string,
		price int,
		rid string,
		name string,
		sale int,
		size string,
		totalPrice int,
		nmId int,
		brand string,
		status int,
		orderUid string,
	) *model.OrderItemModel

	NewOrderItemsModel(models ...*model.OrderItemModel) *model.OrderItemsModel
}
