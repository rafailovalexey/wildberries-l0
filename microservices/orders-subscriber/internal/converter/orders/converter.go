package orders

import (
	"github.com/emptyhopes/orders-subscriber/internal/converter"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
	"time"
)

type Converter struct{}

var _ converter.OrdersConverterInterface = &Converter{}

func (c *Converter) MapOrderDtoToOrderModel(dto *dto.OrderDto, deliveryUid string, paymentUid string) *model.OrderModel {
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

func (c *Converter) MapOrderPaymentDtoToOrderPaymentModel(dto *dto.OrderPaymentDto) *model.OrderPaymentModel {
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

func (c *Converter) MapOrderDeliveryDtoToOrderDeliveryModel(dto *dto.OrderDeliveryDto) *model.OrderDeliveryModel {
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

func (c *Converter) MapOrderItemDtoToOrderItemModel(dto *dto.OrderItemDto, orderUid string) *model.OrderItemModel {
	return &model.OrderItemModel{
		ChrtId:      dto.ChrtId,
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

func (c *Converter) MapOrderItemsDtoToOrderItemsModel(dtos *[]dto.OrderItemDto, orderUid string) *[]model.OrderItemModel {
	models := make([]model.OrderItemModel, len(*dtos))

	for index, value := range *dtos {
		models[index] = *c.MapOrderItemDtoToOrderItemModel(&value, orderUid)
	}

	return &models
}
