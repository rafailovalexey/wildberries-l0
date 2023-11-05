package converter

import (
	model "github.com/emptyhopes/orders-test/internal/model/orders"
	repositoryModel "github.com/emptyhopes/orders-test/internal/repository/model/orders"
)

type OrdersConverterInterface interface {
	OrderModelToOrderModelRepository(model model.OrderModel) repositoryModel.OrderModelRepository
	OrderPaymentModelToOrderPaymentModelRepository(model model.OrderPaymentModel) repositoryModel.OrderPaymentModelRepository
	OrderDeliveryModelToOrderDeliveryModelRepository(model model.OrderDeliveryModel) repositoryModel.OrderDeliveryModelRepository
	OrderItemsModelToOrderItemsModelRepository(models []model.OrderItemModel) []repositoryModel.OrderItemModelRepository

	OrderModelRepositoryToOrderModel(models repositoryModel.OrderModelRepository) model.OrderModel
	OrderPaymentModelRepositoryToOrderPaymentModel(models repositoryModel.OrderPaymentModelRepository) model.OrderPaymentModel
	OrderDeliveryModelRepositoryToOrderDeliveryModel(models repositoryModel.OrderDeliveryModelRepository) model.OrderDeliveryModel
	OrderItemsModelRepositoryToOrderItemsModel(models []repositoryModel.OrderItemModelRepository) []model.OrderItemModel
}
