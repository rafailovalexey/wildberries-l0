package dto

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
)

type OrderDtoInterface interface {
	NewOrderDto(
		OrderUid string,
		TrackNumber string,
		Entry string,
		Delivery *dto.OrderDeliveryDto,
		Payment *dto.OrderPaymentDto,
		Items *dto.OrderItemsDto,
		Locale string,
		InternalSignature string,
		CustomerId string,
		DeliveryService string,
		Shardkey string,
		SmId int,
		DateCreated int64,
		OofShard string,
	) *dto.OrderDto

	NewOrderDeliveryDto(
		Name string,
		Phone string,
		Zip string,
		City string,
		Address string,
		Region string,
		Email string,
	) *dto.OrderDeliveryDto

	NewOrderPaymentDto(
		Transaction string,
		RequestId string,
		Currency string,
		Provider string,
		Amount int,
		PaymentDt int64,
		Bank string,
		DeliveryCost int,
		GoodsTotal int,
		CustomFee int,
	) *dto.OrderPaymentDto

	NewOrderItemDto(
		TrackNumber string,
		Price int,
		Rid string,
		Name string,
		Sale int,
		Size string,
		TotalPrice int,
		NmId int,
		Brand string,
		Status int,
	) *dto.OrderItemDto

	NewOrderItemsDto(
		dtos ...*dto.OrderItemDto,
	) *dto.OrderItemsDto
}
