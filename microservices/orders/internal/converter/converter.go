package converter

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	model "github.com/emptyhopes/orders/internal/model/orders"
)

type OrdersConverterInterface interface {
	OrderDtoToOrderModel(*dto.OrderDto) *model.OrderModel
	OrderPaymentDtoToOrderPaymentModel(*dto.OrderPaymentDto) *model.OrderPaymentModel
	OrderDeliveryDtoToOrderDeliveryModel(*dto.OrderDeliveryDto) *model.OrderDeliveryModel
	OrderItemDtoToOrderItemModel(*dto.OrderItemDto) *model.OrderItemModel
	OrderItemsDtoToOrderItemsModel(*[]dto.OrderItemDto) *[]model.OrderItemModel

	OrderModelToOrderDto(*model.OrderModel, *model.OrderDeliveryModel, *model.OrderPaymentModel, *[]model.OrderItemModel) *dto.OrderDto
	OrderPaymentModelToOrderPaymentDto(*model.OrderPaymentModel) *dto.OrderPaymentDto
	OrderDeliveryModelToOrderDeliveryDto(*model.OrderDeliveryModel) *dto.OrderDeliveryDto
	OrderItemModelToOrderItemDto(*model.OrderItemModel) *dto.OrderItemDto
	OrderItemsModelToOrderItemsDto(*[]model.OrderItemModel) *[]dto.OrderItemDto
}
