package converter

import (
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
)

type OrdersConverterInterface interface {
	OrderDtoToOrderModel(*dto.OrderDto) *model.OrderModel
	OrderPaymentDtoToOrderPaymentModel(*dto.OrderPaymentDto) *model.OrderPaymentModel
	OrderDeliveryDtoToOrderDeliveryModel(*dto.OrderDeliveryDto) *model.OrderDeliveryModel
	OrderItemDtoToOrderItemModel(*dto.OrderItemDto) *model.OrderItemModel
	OrderItemsDtoToOrderItemsModel(*[]dto.OrderItemDto) *[]model.OrderItemModel

	OrderModelToOrderDto(*model.OrderModel) *dto.OrderDto
	OrderPaymentModelToOrderPaymentDto(*model.OrderPaymentModel) *dto.OrderPaymentDto
	OrderDeliveryModelToOrderDeliveryDto(*model.OrderDeliveryModel) *dto.OrderDeliveryDto
	OrderItemModelToOrderItemDto(*model.OrderItemModel) *dto.OrderItemDto
	OrderItemsModelToOrderItemsDto(*[]model.OrderItemModel) *[]dto.OrderItemDto
}
