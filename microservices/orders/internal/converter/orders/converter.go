package orders

import (
	definition "github.com/emptyhopes/orders/internal/converter"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	model "github.com/emptyhopes/orders/internal/model/orders"
	"time"
)

type converter struct{}

var _ definition.OrderConverterInterface = (*converter)(nil)

func NewOrderConverter() *converter {
	return &converter{}
}

func (c *converter) MapOrderDtoToOrderModel(dto *dto.OrderDto, deliveryUid string, paymentUid string) *model.OrderModel {
	return &model.OrderModel{
		OrderUid:          dto.OrderUid,
		TrackNumber:       dto.TrackNumber,
		Entry:             dto.Entry,
		DeliveryUid:       deliveryUid,
		PaymentUid:        paymentUid,
		Locale:            dto.Locale,
		InternalSignature: dto.InternalSignature,
		CustomerId:        dto.CustomerId,
		DeliveryService:   dto.DeliveryService,
		Shardkey:          dto.Shardkey,
		SmId:              dto.SmId,
		DateCreated:       time.Unix(dto.DateCreated, 0),
		OofShard:          dto.OofShard,
	}
}

func (c *converter) MapOrderPaymentDtoToOrderPaymentModel(dto *dto.OrderPaymentDto) *model.OrderPaymentModel {
	return &model.OrderPaymentModel{
		Transaction:  dto.Transaction,
		RequestId:    dto.RequestId,
		Currency:     dto.Currency,
		Provider:     dto.Provider,
		Amount:       dto.Amount,
		PaymentDt:    time.Unix(dto.PaymentDt, 0),
		Bank:         dto.Bank,
		DeliveryCost: dto.DeliveryCost,
		GoodsTotal:   dto.GoodsTotal,
		CustomFee:    dto.CustomFee,
	}
}

func (c *converter) MapOrderDeliveryDtoToOrderDeliveryModel(dto *dto.OrderDeliveryDto) *model.OrderDeliveryModel {
	return &model.OrderDeliveryModel{
		Name:    dto.Name,
		Phone:   dto.Phone,
		Zip:     dto.Zip,
		City:    dto.City,
		Address: dto.Address,
		Region:  dto.Region,
		Email:   dto.Email,
	}
}

func (c *converter) MapOrderItemDtoToOrderItemModel(dto *dto.OrderItemDto, orderUid string) *model.OrderItemModel {
	return &model.OrderItemModel{
		TrackNumber: dto.TrackNumber,
		Price:       dto.Price,
		Rid:         dto.Rid,
		Name:        dto.Name,
		Sale:        dto.Sale,
		Size:        dto.Size,
		TotalPrice:  dto.TotalPrice,
		NmId:        dto.NmId,
		Brand:       dto.Brand,
		Status:      dto.Status,
		OrderUid:    orderUid,
	}
}

func (c *converter) MapOrderItemsDtoToOrderItemsModel(dtos *[]dto.OrderItemDto, orderUid string) *model.OrderItemsModel {
	models := make(model.OrderItemsModel, len(*dtos))

	for index, value := range *dtos {
		models[index] = *c.MapOrderItemDtoToOrderItemModel(&value, orderUid)
	}

	return &models
}

func (c *converter) MapOrderModelToOrderDto(order *model.OrderModel, delivery *model.OrderDeliveryModel, payment *model.OrderPaymentModel, items *model.OrderItemsModel) *dto.OrderDto {
	return &dto.OrderDto{
		OrderUid:          order.OrderUid,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Delivery:          c.MapOrderDeliveryModelToOrderDeliveryDto(delivery),
		Payment:           c.MapOrderPaymentModelToOrderPaymentDto(payment),
		Items:             c.MapOrderItemsModelToOrderItemsDto(items),
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerId:        order.CustomerId,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmId:              order.SmId,
		DateCreated:       order.DateCreated.Unix(),
		OofShard:          order.OofShard,
	}
}

func (c *converter) MapOrderPaymentModelToOrderPaymentDto(model *model.OrderPaymentModel) *dto.OrderPaymentDto {
	return &dto.OrderPaymentDto{
		Transaction:  model.Transaction,
		RequestId:    model.RequestId,
		Currency:     model.Currency,
		Provider:     model.Provider,
		Amount:       model.Amount,
		PaymentDt:    model.PaymentDt.Unix(),
		Bank:         model.Bank,
		DeliveryCost: model.DeliveryCost,
		GoodsTotal:   model.GoodsTotal,
		CustomFee:    model.CustomFee,
	}
}

func (c *converter) MapOrderDeliveryModelToOrderDeliveryDto(model *model.OrderDeliveryModel) *dto.OrderDeliveryDto {
	return &dto.OrderDeliveryDto{
		Name:    model.Name,
		Phone:   model.Phone,
		Zip:     model.Zip,
		City:    model.City,
		Address: model.Address,
		Region:  model.Region,
		Email:   model.Email,
	}
}

func (c *converter) MapOrderItemModelToOrderItemDto(model *model.OrderItemModel) *dto.OrderItemDto {
	return &dto.OrderItemDto{
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

func (c *converter) MapOrderItemsModelToOrderItemsDto(models *model.OrderItemsModel) *[]dto.OrderItemDto {
	dtos := make([]dto.OrderItemDto, len(*models))

	for index, value := range *models {
		dtos[index] = *c.MapOrderItemModelToOrderItemDto(&value)
	}

	return &dtos
}
