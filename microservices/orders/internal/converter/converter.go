package converter

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	model "github.com/emptyhopes/orders/internal/model/orders"
)

type OrderConverterInterface interface {
	MapOrderDtoToOrderModel(*dto.OrderDto, string, string) *model.OrderModel
	MapOrderPaymentDtoToOrderPaymentModel(*dto.OrderPaymentDto) *model.OrderPaymentModel
	MapOrderDeliveryDtoToOrderDeliveryModel(*dto.OrderDeliveryDto) *model.OrderDeliveryModel
	MapOrderItemDtoToOrderItemModel(*dto.OrderItemDto, string) *model.OrderItemModel
	MapOrderItemsDtoToOrderItemsModel(*[]dto.OrderItemDto, string) *[]model.OrderItemModel

	MapOrderModelToOrderDto(*model.OrderModel, *model.OrderDeliveryModel, *model.OrderPaymentModel, *[]model.OrderItemModel) *dto.OrderDto
	MapOrderPaymentModelToOrderPaymentDto(*model.OrderPaymentModel) *dto.OrderPaymentDto
	MapOrderDeliveryModelToOrderDeliveryDto(*model.OrderDeliveryModel) *dto.OrderDeliveryDto
	MapOrderItemModelToOrderItemDto(*model.OrderItemModel) *dto.OrderItemDto
	MapOrderItemsModelToOrderItemsDto(*[]model.OrderItemModel) *[]dto.OrderItemDto
}
