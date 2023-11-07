package converter

import (
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	model "github.com/emptyhopes/orders_subscriber/internal/model/orders"
)

type OrderConverterInterface interface {
	MapOrderDtoToOrderModel(*dto.OrderDto, string, string) *model.OrderModel
	MapOrderPaymentDtoToOrderPaymentModel(*dto.OrderPaymentDto) *model.OrderPaymentModel
	MapOrderDeliveryDtoToOrderDeliveryModel(*dto.OrderDeliveryDto) *model.OrderDeliveryModel
	MapOrderItemDtoToOrderItemModel(*dto.OrderItemDto, string) *model.OrderItemModel
	MapOrderItemsDtoToOrderItemsModel(*[]dto.OrderItemDto, string) *model.OrderItemsModel

	MapOrderModelToOrderDto(*model.OrderModel, *model.OrderDeliveryModel, *model.OrderPaymentModel, *model.OrderItemsModel) *dto.OrderDto
	MapOrderPaymentModelToOrderPaymentDto(*model.OrderPaymentModel) *dto.OrderPaymentDto
	MapOrderDeliveryModelToOrderDeliveryDto(*model.OrderDeliveryModel) *dto.OrderDeliveryDto
	MapOrderItemModelToOrderItemDto(*model.OrderItemModel) *dto.OrderItemDto
	MapOrderItemsModelToOrderItemsDto(*model.OrderItemsModel) *[]dto.OrderItemDto
}
