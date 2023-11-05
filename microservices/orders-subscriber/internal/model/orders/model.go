package orders

type OrderModel struct {
	TrackNumber       string              `json:"track_number"`
	Entry             string              `json:"entry"`
	Delivery          *OrderDeliveryModel `json:"delivery"`
	Payment           *OrderPaymentModel  `json:"payment"`
	Items             *[]OrderItemModel   `json:"items"`
	Locale            string              `json:"locale"`
	InternalSignature string              `json:"internal_signature"`
	CustomerId        string              `json:"customer_id"`
	DeliveryService   string              `json:"delivery_service"`
	Shardkey          string              `json:"shardkey"`
	SmId              int                 `json:"sm_id"`
	DateCreated       string              `json:"date_created"`
	OofShard          string              `json:"oof_shard"`
}

type OrderDeliveryModel struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type OrderPaymentModel struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type OrderItemModel struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func ConstructorOrderModel(
	OrderUid string,
	TrackNumber string,
	Entry string,
	Delivery *OrderDeliveryModel,
	Payment *OrderPaymentModel,
	Items *[]OrderItemModel,
	Locale string,
	InternalSignature string,
	CustomerId string,
	DeliveryService string,
	Shardkey string,
	SmId int,
	DateCreated string,
	OofShard string,
) *OrderModel {
	result := &OrderModel{
		TrackNumber:       TrackNumber,
		Entry:             Entry,
		Delivery:          Delivery,
		Payment:           Payment,
		Items:             Items,
		Locale:            Locale,
		InternalSignature: InternalSignature,
		CustomerId:        CustomerId,
		DeliveryService:   DeliveryService,
		Shardkey:          Shardkey,
		SmId:              SmId,
		DateCreated:       DateCreated,
		OofShard:          OofShard,
	}

	return result
}

func ConstructorOrderDeliveryModel(
	Name string,
	Phone string,
	Zip string,
	City string,
	Address string,
	Region string,
	Email string,
) *OrderDeliveryModel {
	result := &OrderDeliveryModel{
		Name:    Name,
		Phone:   Phone,
		Zip:     Zip,
		City:    City,
		Address: Address,
		Region:  Region,
		Email:   Email,
	}

	return result
}

func ConstructorOrderPaymentModel(
	Transaction string,
	RequestId string,
	Currency string,
	Provider string,
	Amount int,
	PaymentDt int,
	Bank string,
	DeliveryCost int,
	GoodsTotal int,
	CustomFee int,
) *OrderPaymentModel {
	result := &OrderPaymentModel{
		Transaction:  Transaction,
		RequestId:    RequestId,
		Currency:     Currency,
		Provider:     Provider,
		Amount:       Amount,
		PaymentDt:    PaymentDt,
		Bank:         Bank,
		DeliveryCost: DeliveryCost,
		GoodsTotal:   GoodsTotal,
		CustomFee:    CustomFee,
	}

	return result
}

func ConstructorOrderItemModel(
	ChrtId int,
	TrackNumber string,
	Price int,
	Rid string,
	Name string,
	Sale int,
	Size string,
	TotalPrice int,
	NmId int,
	Brand string,
	Status int,
) *OrderItemModel {
	result := &OrderItemModel{
		ChrtId:      ChrtId,
		TrackNumber: TrackNumber,
		Price:       Price,
		Rid:         Rid,
		Name:        Name,
		Sale:        Sale,
		Size:        Size,
		TotalPrice:  TotalPrice,
		NmId:        NmId,
		Brand:       Brand,
		Status:      Status,
	}

	return result
}
