package orders

import (
	"context"
	"fmt"
	"github.com/emptyhopes/orders-subscriber/internal/helpers"
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
	"github.com/emptyhopes/orders-subscriber/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Repository struct{}

var _ repository.OrdersRepositoryInterface = &Repository{}

func (r *Repository) Cache(model *model.OrderModel) (bool, error) {
	value, isExist := repository.Cache.Get(model.OrderUid)

	fmt.Println(value, isExist)

	if !isExist {
		fmt.Println("set value")

		repository.Cache.Set(model.OrderUid, model, 5*time.Minute)
	}

	return true, nil
}

func (r *Repository) CreateOrder(pool *pgxpool.Pool, order *model.OrderModel, delivery *model.OrderDeliveryModel, payment *model.OrderPaymentModel, items *[]model.OrderItemModel) error {
	transactions, err := helpers.ConstructorTransactions(context.Background(), pool)

	if err != nil {
		return err
	}

	defer transactions.Rollback(context.Background())

	if err := r.insertOrderPayment(transactions, payment); err != nil {
		return err
	}

	if err := r.insertOrderDelivery(transactions, delivery); err != nil {
		return err
	}

	if err := r.insertOrder(transactions, order); err != nil {
		return err
	}

	if err := r.insertOrderItems(transactions, items); err != nil {
		return err
	}

	if err := transactions.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func (r *Repository) insertOrderPayment(transactions *helpers.Transactions, payment *model.OrderPaymentModel) error {
	query := `
        INSERT INTO payment (id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `

	_ = transactions.QueryRow(
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
	)

	return nil
}

func (r *Repository) insertOrderDelivery(transactions *helpers.Transactions, delivery *model.OrderDeliveryModel) error {
	query := `
        INSERT INTO delivery (id, name, phone, zip, city, address, region, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_ = transactions.QueryRow(
		context.Background(),
		query,
		//delivery.Id,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	)

	return nil
}

func (r *Repository) insertOrder(transactions *helpers.Transactions, order *model.OrderModel) error {
	query := `
        INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    `

	_ = transactions.QueryRow(
		context.Background(),
		query,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		//order.Delivery.Id,
		//order.Payment.Id,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	)

	return nil
}

func (r *Repository) insertOrderItems(transactions *helpers.Transactions, items *[]model.OrderItemModel) error {
	query := `
        INSERT INTO order_items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `

	for _, item := range *items {
		_ = transactions.QueryRow(
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
			//item.OrderId,
		)
	}

	return nil
}
