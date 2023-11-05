package orders

import (
	"github.com/emptyhopes/orders-test/internal/converter"
	model "github.com/emptyhopes/orders-test/internal/model/orders"
	repositoryModel "github.com/emptyhopes/orders-test/internal/repository/model/orders"
)

type Converter struct{}

var _ converter.OrdersConverterInterface = &Converter{}

func (c *Converter) OrderModelToOrderModelRepository(model *model.OrderModel) *repositoryModel.OrderModelRepository {
	return &repositoryModel.OrderModelRepository{
		OrderUid:          model.OrderUid,
		TrackNumber:       model.TrackNumber,
		Entry:             model.Entry,
		Delivery:          c.OrderDeliveryModelToOrderDeliveryModelRepository(model.Delivery),
		Payment:           c.OrderPaymentModelToOrderPaymentModelRepository(model.Payment),
		Items:             c.OrderItemsModelToOrderItemsModelRepository(model.Items),
		Locale:            model.Locale,
		InternalSignature: model.InternalSignature,
		CustomerId:        model.CustomerId,
		DeliveryService:   model.DeliveryService,
		Shardkey:          model.Shardkey,
		SmId:              model.SmId,
		DateCreated:       model.DateCreated,
		OofShard:          model.OofShard,
	}
}

func (c *Converter) OrderPaymentModelToOrderPaymentModelRepository(model *model.OrderPaymentModel) *repositoryModel.OrderPaymentModelRepository {
	return &repositoryModel.OrderPaymentModelRepository{
		Transaction:  model.Transaction,
		RequestId:    model.RequestId,
		Currency:     model.Currency,
		Provider:     model.Provider,
		Amount:       model.Amount,
		PaymentDt:    model.PaymentDt,
		Bank:         model.Bank,
		DeliveryCost: model.DeliveryCost,
		GoodsTotal:   model.GoodsTotal,
		CustomFee:    model.CustomFee,
	}
}

func (c *Converter) OrderDeliveryModelToOrderDeliveryModelRepository(model *model.OrderDeliveryModel) *repositoryModel.OrderDeliveryModelRepository {
	return &repositoryModel.OrderDeliveryModelRepository{
		Name:    model.Name,
		Phone:   model.Phone,
		Zip:     model.Zip,
		City:    model.City,
		Address: model.Address,
		Region:  model.Region,
		Email:   model.Email,
	}
}

func (c *Converter) OrderItemModelToOrderItemModelRepository(model *model.OrderItemModel) *repositoryModel.OrderItemModelRepository {
	return &repositoryModel.OrderItemModelRepository{
		ChrtId:      model.ChrtId,
		TrackNumber: model.TrackNumber,
		Price:       model.Price,
		Rid:         model.Rid,
		Name:        model.Name,
		Sale:        model.Sale,
		Size:        model.Size,
		TotalPrice:  model.TotalPrice,
		NmId:        model.NmId,
		Brand:       model.Brand,
		Status:      model.Status,
	}
}

func (c *Converter) OrderItemsModelToOrderItemsModelRepository(models *[]model.OrderItemModel) *[]repositoryModel.OrderItemModelRepository {
	repositoryModels := make([]repositoryModel.OrderItemModelRepository, len(*models))

	for index, value := range *models {
		repositoryModels[index] = *c.OrderItemModelToOrderItemModelRepository(&value)
	}

	return &repositoryModels
}

func (c *Converter) OrderModelRepositoryToOrderModel(repositoryModel *repositoryModel.OrderModelRepository) *model.OrderModel {
	return &model.OrderModel{
		OrderUid:          repositoryModel.OrderUid,
		TrackNumber:       repositoryModel.TrackNumber,
		Entry:             repositoryModel.Entry,
		Delivery:          c.OrderDeliveryModelRepositoryToOrderDeliveryModel(repositoryModel.Delivery),
		Payment:           c.OrderPaymentModelRepositoryToOrderPaymentModel(repositoryModel.Payment),
		Items:             c.OrderItemsModelRepositoryToOrderItemsModel(repositoryModel.Items),
		Locale:            repositoryModel.Locale,
		InternalSignature: repositoryModel.InternalSignature,
		CustomerId:        repositoryModel.CustomerId,
		DeliveryService:   repositoryModel.DeliveryService,
		Shardkey:          repositoryModel.Shardkey,
		SmId:              repositoryModel.SmId,
		DateCreated:       repositoryModel.DateCreated,
		OofShard:          repositoryModel.OofShard,
	}
}

func (c *Converter) OrderPaymentModelRepositoryToOrderPaymentModel(repositoryModel *repositoryModel.OrderPaymentModelRepository) *model.OrderPaymentModel {
	return &model.OrderPaymentModel{
		Transaction:  repositoryModel.Transaction,
		RequestId:    repositoryModel.RequestId,
		Currency:     repositoryModel.Currency,
		Provider:     repositoryModel.Provider,
		Amount:       repositoryModel.Amount,
		PaymentDt:    repositoryModel.PaymentDt,
		Bank:         repositoryModel.Bank,
		DeliveryCost: repositoryModel.DeliveryCost,
		GoodsTotal:   repositoryModel.GoodsTotal,
		CustomFee:    repositoryModel.CustomFee,
	}
}

func (c *Converter) OrderDeliveryModelRepositoryToOrderDeliveryModel(repositoryModel *repositoryModel.OrderDeliveryModelRepository) *model.OrderDeliveryModel {
	return &model.OrderDeliveryModel{
		Name:    repositoryModel.Name,
		Phone:   repositoryModel.Phone,
		Zip:     repositoryModel.Zip,
		City:    repositoryModel.City,
		Address: repositoryModel.Address,
		Region:  repositoryModel.Region,
		Email:   repositoryModel.Email,
	}
}

func (c *Converter) OrderItemModelRepositoryToOrderItemModel(repositoryModel *repositoryModel.OrderItemModelRepository) *model.OrderItemModel {
	return &model.OrderItemModel{
		ChrtId:      repositoryModel.ChrtId,
		TrackNumber: repositoryModel.TrackNumber,
		Price:       repositoryModel.Price,
		Rid:         repositoryModel.Rid,
		Name:        repositoryModel.Name,
		Sale:        repositoryModel.Sale,
		Size:        repositoryModel.Size,
		TotalPrice:  repositoryModel.TotalPrice,
		NmId:        repositoryModel.NmId,
		Brand:       repositoryModel.Brand,
		Status:      repositoryModel.Status,
	}
}

func (c *Converter) OrderItemsModelRepositoryToOrderItemsModel(repositoryModels *[]repositoryModel.OrderItemModelRepository) *[]model.OrderItemModel {
	models := make([]model.OrderItemModel, len(*repositoryModels))

	for index, value := range *repositoryModels {
		models[index] = *c.OrderItemModelRepositoryToOrderItemModel(&value)
	}

	return &models
}
