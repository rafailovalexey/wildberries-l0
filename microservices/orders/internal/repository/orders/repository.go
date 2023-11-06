package orders

import (
	"context"
	converter "github.com/emptyhopes/orders/internal/converter/orders"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	model "github.com/emptyhopes/orders/internal/model/orders"
	"github.com/emptyhopes/orders/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type Repository struct {
	rwmutex sync.RWMutex
}

var _ repository.OrdersRepositoryInterface = &Repository{}

//func (r *Repository) GetCachedOrderById(orderUid string, orderDto *dto.OrderDto) (*dto.OrderDto, error) {
//	orderCached, isExist := repository.Cache.Get(orderUid)
//
//	if !isExist {
//		repository.Cache.Set(orderUid, orderDto, 5*time.Minute)
//	}
//
//	return orderCached, nil
//}

func (r *Repository) GetOrderById(orderUid string) (*dto.OrderDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := repository.Database.GetPool()
	defer pool.Close()

	converterOrders := &converter.Converter{}

	orderModel, err := r.getOrder(pool, orderUid)
	if err != nil {
		return nil, err
	}

	orderDeliveryModel, err := r.getOrderDelivery(pool, orderModel.DeliveryUid)
	if err != nil {
		return nil, err
	}

	orderPaymentModel, err := r.getOrderPayment(pool, orderModel.PaymentUid)
	if err != nil {
		return nil, err
	}

	orderItemsModel, err := r.getOrderItems(pool, orderModel.OrderUid)
	if err != nil {
		return nil, err
	}

	orderDto := converterOrders.MapOrderModelToOrderDto(orderModel, orderDeliveryModel, orderPaymentModel, orderItemsModel)

	return orderDto, nil
}

func (r *Repository) getOrder(pool *pgxpool.Pool, orderUid string) (*model.OrderModel, error) {
	var order model.OrderModel

	query := `
        SELECT * FROM orders WHERE order_uid = $1
    `

	err := pool.QueryRow(
		context.Background(),
		query,
		orderUid,
	).Scan(&order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *Repository) getOrderPayment(pool *pgxpool.Pool, paymentUid string) (*model.OrderPaymentModel, error) {
	var payment model.OrderPaymentModel

	query := `
        SELECT * FROM orders_payment WHERE payment_uid = $1
    `

	err := pool.QueryRow(
		context.Background(),
		query,
		paymentUid,
	).Scan(&payment)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *Repository) getOrderDelivery(pool *pgxpool.Pool, deliveryUid string) (*model.OrderDeliveryModel, error) {
	var delivery model.OrderDeliveryModel

	query := `
        SELECT * FROM orders_delivery WHERE delivery_uid = $1
    `

	err := pool.QueryRow(
		context.Background(),
		query,
		deliveryUid,
	).Scan(&delivery)

	if err != nil {
		return nil, err
	}

	return &delivery, nil
}

func (r *Repository) getOrderItems(pool *pgxpool.Pool, orderUid string) (*[]model.OrderItemModel, error) {
	items := make([]model.OrderItemModel, 0, 10)

	query := `
        SELECT * FROM orders_items WHERE order_uid = $1
    `

	err := pool.QueryRow(
		context.Background(),
		query,
		orderUid,
	).Scan(&items)

	if err != nil {
		return nil, err
	}

	return &items, nil
}
