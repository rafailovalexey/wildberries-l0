package orders

import (
	"context"
	"fmt"
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

	fmt.Println("ya tyt 1")
	orderModel, err := r.getOrder(pool, orderUid)
	if err != nil {
		return nil, err
	}

	fmt.Println("ya tyt 2")
	orderDeliveryModel, err := r.getOrderDelivery(pool, orderModel.DeliveryUid)
	if err != nil {
		return nil, err
	}

	fmt.Println("ya tyt 3")
	orderPaymentModel, err := r.getOrderPayment(pool, orderModel.PaymentUid)
	if err != nil {
		return nil, err
	}

	fmt.Println("ya tyt 4")
	orderItemsModel, err := r.getOrderItems(pool, orderModel.OrderUid)
	if err != nil {
		return nil, err
	}

	fmt.Println("ya tyt 5")
	orderDto := converterOrders.MapOrderModelToOrderDto(orderModel, orderDeliveryModel, orderPaymentModel, orderItemsModel)

	return orderDto, nil
}

func (r *Repository) getOrder(pool *pgxpool.Pool, orderUid string) (*model.OrderModel, error) {
	query := `
        SELECT
           	order_uid,         
			track_number,      
			entry,            
			delivery_uid,      
			payment_uid,       
			locale,           
			internal_signature,
			customer_id,      
			delivery_service, 
			shardkey,       
			sm_id,          
			date_created,    
			oof_shard         
        FROM orders WHERE order_uid = $1
    `

	order := model.OrderModel{}

	err := pool.QueryRow(
		context.Background(),
		query,
		orderUid,
	).Scan(
		&order.OrderUid,
		&order.TrackNumber,
		&order.Entry,
		&order.DeliveryUid,
		&order.PaymentUid,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *Repository) getOrderPayment(pool *pgxpool.Pool, paymentUid string) (*model.OrderPaymentModel, error) {
	var payment model.OrderPaymentModel

	query := `
        SELECT
			payment_uid,   
			transaction,  
			request_id,    
			currency,     
			provider,     
			amount,       
			payment_dt,    
			bank,         
			delivery_cost, 
			goods_total,   
			custom_fee    
        FROM orders_payment WHERE payment_uid = $1
    `

	err := pool.QueryRow(
		context.Background(),
		query,
		paymentUid,
	).Scan(
		&payment.PaymentUid,
		&payment.Transaction,
		&payment.RequestId,
		&payment.Currency,
		&payment.Provider,
		&payment.Amount,
		&payment.PaymentDt,
		&payment.Bank,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
	)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *Repository) getOrderDelivery(pool *pgxpool.Pool, deliveryUid string) (*model.OrderDeliveryModel, error) {
	var delivery model.OrderDeliveryModel

	query := `
        SELECT 
            delivery_uid, 
			name,       
			phone,      
			zip,       
			city,       
			address,    
			region,     
			email      
        FROM orders_delivery WHERE delivery_uid = $1
    `

	err := pool.QueryRow(
		context.Background(),
		query,
		deliveryUid,
	).Scan(
		&delivery.DeliveryUid,
		&delivery.Name,
		&delivery.Phone,
		&delivery.Zip,
		&delivery.City,
		&delivery.Address,
		&delivery.Region,
		&delivery.Email,
	)

	if err != nil {
		return nil, err
	}

	return &delivery, nil
}

func (r *Repository) getOrderItems(pool *pgxpool.Pool, orderUid string) (*[]model.OrderItemModel, error) {
	query := `
        SELECT
            chrt_id,     
			track_number,
			price,      
			rid,        
			name,       
			sale,       
			size,       
			total_price, 
			nm_id,       
			brand,      
			status,     
			order_uid   
        FROM orders_items WHERE order_uid = $1
    `

	rows, err := pool.Query(
		context.Background(),
		query,
		orderUid,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]model.OrderItemModel, 0, 10)

	for rows.Next() {
		item := model.OrderItemModel{}

		err := rows.Scan(
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
			&item.OrderUid,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return &items, nil
}
