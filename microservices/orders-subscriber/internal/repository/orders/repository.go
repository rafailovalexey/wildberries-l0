package orders

import (
	"context"
	"fmt"
	converter "github.com/emptyhopes/orders-subscriber/internal/converter/orders"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders-subscriber/internal/helpers"
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
	"github.com/emptyhopes/orders-subscriber/internal/repository"
	"sync"
)

type Repository struct {
	rwmutex sync.RWMutex
}

var _ repository.OrdersRepositoryInterface = &Repository{}

//func (r *Repository) Cache(order *dto.OrderDto) *dto.OrderDto {
//	value, isExist := repository.Cache.Get(order.OrderUid)
//
//	fmt.Println(value, isExist)
//
//	if !isExist {
//		fmt.Println("set value")
//
//		repository.Cache.Set(order.OrderUid, order, 5*time.Minute)
//	}
//
//	return value
//}

func (r *Repository) CreateOrder(order *dto.OrderDto) error {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := repository.Database.GetPool()
	defer pool.Close()

	converterOrders := &converter.Converter{}

	transactions, err := helpers.ConstructorTransactions(context.Background(), pool)
	if err != nil {
		return err
	}

	defer transactions.Rollback(context.Background())

	orderPaymentModel := converterOrders.MapOrderPaymentDtoToOrderPaymentModel(order.Payment)
	fmt.Println("ya tyt 1")
	paymentUid, err := r.insertOrderPayment(transactions, orderPaymentModel)
	if err != nil {
		return err
	}

	orderDeliveryModel := converterOrders.MapOrderDeliveryDtoToOrderDeliveryModel(order.Delivery)
	fmt.Println("ya tyt 2")
	deliveryUid, err := r.insertOrderDelivery(transactions, orderDeliveryModel)
	if err != nil {
		return err
	}

	orderModel := converterOrders.MapOrderDtoToOrderModel(order, deliveryUid, paymentUid)
	fmt.Println("ya tyt 3")
	if err := r.insertOrder(transactions, orderModel); err != nil {
		return err
	}

	orderItemsModel := converterOrders.MapOrderItemsDtoToOrderItemsModel(order.Items, orderModel.OrderUid)
	fmt.Println("ya tyt 4")
	if err := r.insertOrderItems(transactions, orderItemsModel); err != nil {
		return err
	}

	fmt.Println("ya tyt 5")

	if err := transactions.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func (r *Repository) insertOrderPayment(transactions *helpers.Transactions, payment *model.OrderPaymentModel) (string, error) {
	var paymentUid string

	query := `
        INSERT INTO orders_payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING payment_uid;
    `

	err := transactions.QueryRow(
		context.Background(),
		query,
		payment.Transaction,
		payment.RequestId,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	).Scan(&paymentUid)

	if err != nil {
		return "", err
	}

	return paymentUid, nil
}

func (r *Repository) insertOrderDelivery(transactions *helpers.Transactions, delivery *model.OrderDeliveryModel) (string, error) {
	var deliveryUid string

	query := `
        INSERT INTO orders_delivery (name, phone, zip, city, address, region, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING delivery_uid;
    `

	err := transactions.QueryRow(
		context.Background(),
		query,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	).Scan(&deliveryUid)

	if err != nil {
		return "", err
	}

	return deliveryUid, nil
}

func (r *Repository) insertOrder(transactions *helpers.Transactions, order *model.OrderModel) error {
	query := `
        INSERT INTO orders (order_uid, track_number, entry, delivery_uid, payment_uid, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    `

	_, err := transactions.Exec(
		context.Background(),
		query,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.DeliveryUid,
		order.PaymentUid,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) insertOrderItems(transactions *helpers.Transactions, items *[]model.OrderItemModel) error {
	query := `
        INSERT INTO orders_items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    `

	for _, item := range *items {
		_, err := transactions.Exec(
			context.Background(),
			query,
			item.ChrtId,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmId,
			item.Brand,
			item.Status,
			item.OrderUid,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
