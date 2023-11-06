package converter

import (
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
)

type OrdersConverterInterface interface {
	MapOrderDtoToOrderModel(*dto.OrderDto, string, string) *model.OrderModel
	MapOrderPaymentDtoToOrderPaymentModel(*dto.OrderPaymentDto) *model.OrderPaymentModel
	MapOrderDeliveryDtoToOrderDeliveryModel(*dto.OrderDeliveryDto) *model.OrderDeliveryModel
	MapOrderItemDtoToOrderItemModel(*dto.OrderItemDto, string) *model.OrderItemModel
	MapOrderItemsDtoToOrderItemsModel(*[]dto.OrderItemDto, string) *[]model.OrderItemModel
}