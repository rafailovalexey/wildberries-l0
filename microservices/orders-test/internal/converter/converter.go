package converter

import (
	model "github.com/emptyhopes/orders-test/internal/model/orders"
	repositoryModel "github.com/emptyhopes/orders-test/internal/repository/model/orders"
)

type OrdersConverterInterface interface {
	OrderModelToOrderModelRepository(*model.OrderModel) *repositoryModel.OrderModelRepository
	OrderPaymentModelToOrderPaymentModelRepository(*model.OrderPaymentModel) *repositoryModel.OrderPaymentModelRepository
	OrderDeliveryModelToOrderDeliveryModelRepository(*model.OrderDeliveryModel) *repositoryModel.OrderDeliveryModelRepository
	OrderItemModelToOrderItemModelRepository(*model.OrderItemModel) *repositoryModel.OrderItemModelRepository
	OrderItemsModelToOrderItemsModelRepository(*[]model.OrderItemModel) *[]repositoryModel.OrderItemModelRepository

	OrderModelRepositoryToOrderModel(*repositoryModel.OrderModelRepository) *model.OrderModel
	OrderPaymentModelRepositoryToOrderPaymentModel(*repositoryModel.OrderPaymentModelRepository) *model.OrderPaymentModel
	OrderDeliveryModelRepositoryToOrderDeliveryModel(*repositoryModel.OrderDeliveryModelRepository) *model.OrderDeliveryModel
	OrderItemModelRepositoryToOrderItemModel(*repositoryModel.OrderItemModelRepository) *model.OrderItemModel
	OrderItemsModelRepositoryToOrderItemsModel(*[]repositoryModel.OrderItemModelRepository) *[]model.OrderItemModel
}
