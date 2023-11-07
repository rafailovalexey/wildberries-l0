package orders

import (
	"context"
	converter "github.com/emptyhopes/orders-subscriber/internal/converter/orders"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders-subscriber/internal/helpers"
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
	def "github.com/emptyhopes/orders-subscriber/internal/repository"
	"github.com/emptyhopes/orders-subscriber/storage"
	"sync"
	"time"
)

type repository struct {
	rwmutex sync.RWMutex
}

var _ def.OrdersRepositoryInterface = &repository{}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetOrdersCache() map[string]storage.CacheItem {
	return def.Cache.GetCache()
}

func (r *repository) GetOrderCacheById(id string) (*dto.OrderDto, bool) {
	orderCached, isExist := def.Cache.Get(id)

	if orderDto, ok := orderCached.(*dto.OrderDto); ok {
		return orderDto, isExist
	}

	return nil, false
}

func (r *repository) SetOrderCache(id string, orderDto *dto.OrderDto) {
	def.Cache.Set(id, orderDto, 5*time.Minute)
}

func (r *repository) DeleteOrderCacheById(id string) {
	def.Cache.Delete(id)
}

func (r *repository) CreateOrder(order *dto.OrderDto) error {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := def.Database.GetPool()
	defer pool.Close()

	converterOrders := converter.NewConverter()

	transactions, err := helpers.NewTransactions(context.Background(), pool)
	if err != nil {
		return err
	}

	defer transactions.Rollback(context.Background())

	orderPaymentModel := converterOrders.MapOrderPaymentDtoToOrderPaymentModel(order.Payment)
	paymentUid, err := r.insertOrderPayment(transactions, orderPaymentModel)
	if err != nil {
		return err
	}

	orderDeliveryModel := converterOrders.MapOrderDeliveryDtoToOrderDeliveryModel(order.Delivery)
	deliveryUid, err := r.insertOrderDelivery(transactions, orderDeliveryModel)
	if err != nil {
		return err
	}

	orderModel := converterOrders.MapOrderDtoToOrderModel(order, deliveryUid, paymentUid)
	if err := r.insertOrder(transactions, orderModel); err != nil {
		return err
	}

	orderItemsModel := converterOrders.MapOrderItemsDtoToOrderItemsModel(order.Items, orderModel.OrderUid)
	if err := r.insertOrderItems(transactions, orderItemsModel); err != nil {
		return err
	}

	if err := transactions.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func (r *repository) insertOrderPayment(transactions *helpers.Transactions, payment *model.OrderPaymentModel) (string, error) {
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

func (r *repository) insertOrderDelivery(transactions *helpers.Transactions, delivery *model.OrderDeliveryModel) (string, error) {
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

func (r *repository) insertOrder(transactions *helpers.Transactions, order *model.OrderModel) error {
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

func (r *repository) insertOrderItems(transactions *helpers.Transactions, items *[]model.OrderItemModel) error {
	query := `
        INSERT INTO orders_items (track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `

	for _, item := range *items {
		_, err := transactions.Exec(
			context.Background(),
			query,
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
