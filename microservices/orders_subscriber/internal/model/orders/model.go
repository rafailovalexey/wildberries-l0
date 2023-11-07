package orders

import (
	"time"
)

type OrderModel struct {
	OrderUid          string
	TrackNumber       string
	Entry             string
	DeliveryUid       string
	PaymentUid        string
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	Shardkey          string
	SmId              int
	DateCreated       time.Time
	OofShard          string
}

type OrderDeliveryModel struct {
	DeliveryUid string
	Name        string
	Phone       string
	Zip         string
	City        string
	Address     string
	Region      string
	Email       string
}

type OrderPaymentModel struct {
	PaymentUid   string
	Transaction  string
	RequestId    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDt    time.Time
	Bank         string
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}

type OrderItemModel struct {
	ChrtId      int
	TrackNumber string
	Price       int
	Rid         string
	Name        string
	Sale        int
	Size        string
	TotalPrice  int
	NmId        int
	Brand       string
	Status      int
	OrderUid    string
}

type OrderItemsModel = []OrderItemModel

func NewOrderModel(
	orderUid string,
	trackNumber string,
	entry string,
	deliveryUid string,
	paymentUid string,
	locale string,
	internalSignature string,
	customerId string,
	deliveryService string,
	shardkey string,
	smId int,
	dateCreated time.Time,
	oofShard string,
) *OrderModel {
	return &OrderModel{
		OrderUid:          orderUid,
		TrackNumber:       trackNumber,
		Entry:             entry,
		DeliveryUid:       deliveryUid,
		PaymentUid:        paymentUid,
		Locale:            locale,
		InternalSignature: internalSignature,
		CustomerId:        customerId,
		DeliveryService:   deliveryService,
		Shardkey:          shardkey,
		SmId:              smId,
		DateCreated:       dateCreated,
		OofShard:          oofShard,
	}
}

func NewOrderDeliveryModel(
	deliveryUid string,
	name string,
	phone string,
	zip string,
	city string,
	address string,
	region string,
	email string,
) *OrderDeliveryModel {
	return &OrderDeliveryModel{
		DeliveryUid: deliveryUid,
		Name:        name,
		Phone:       phone,
		Zip:         zip,
		City:        city,
		Address:     address,
		Region:      region,
		Email:       email,
	}
}

func NewOrderPaymentModel(
	paymentUid string,
	transaction string,
	requestId string,
	currency string,
	provider string,
	amount int,
	paymentDt time.Time,
	bank string,
	deliveryCost int,
	goodsTotal int,
	customFee int,
) *OrderPaymentModel {
	return &OrderPaymentModel{
		PaymentUid:   paymentUid,
		Transaction:  transaction,
		RequestId:    requestId,
		Currency:     currency,
		Provider:     provider,
		Amount:       amount,
		PaymentDt:    paymentDt,
		Bank:         bank,
		DeliveryCost: deliveryCost,
		GoodsTotal:   goodsTotal,
		CustomFee:    customFee,
	}
}

func NewOrderItemModel(
	chrtId int,
	trackNumber string,
	price int,
	rid string,
	name string,
	sale int,
	size string,
	totalPrice int,
	nmId int,
	brand string,
	status int,
	orderUid string,
) *OrderItemModel {
	return &OrderItemModel{
		ChrtId:      chrtId,
		TrackNumber: trackNumber,
		Price:       price,
		Rid:         rid,
		Name:        name,
		Sale:        sale,
		Size:        size,
		TotalPrice:  totalPrice,
		NmId:        nmId,
		Brand:       brand,
		Status:      status,
		OrderUid:    orderUid,
	}
}

func NewOrderItemsModel(
	models ...*OrderItemModel,
) *OrderItemsModel {
	items := make(OrderItemsModel, 0, 10)

	for _, model := range models {
		items = append(items, *model)
	}

	return &items
}
