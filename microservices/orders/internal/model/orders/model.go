package orders

import "time"

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
