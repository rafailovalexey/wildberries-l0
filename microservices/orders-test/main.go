package main

import (
	"context"
	model "github.com/emptyhopes/orders-test/internal/dto/orders"
	"github.com/emptyhopes/orders-test/storage"
	"github.com/emptyhopes/orders-test/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var database = &storage.Database{}

func init() {
	database.Initialize()
}

func main() {
	pool := database.GetPool()
	defer pool.Close()

	CreateTables(pool)
}

func InsertTest(pool *pgxpool.Pool, order *model.OrderModel, delivery *model.OrderDeliveryModel, payment *model.OrderPaymentModel, items *[]model.OrderItemModel) error {
	transactions, err := utils.ConstructorTransactions(context.Background(), pool)

	if err != nil {
		return err
	}

	defer transactions.Rollback(context.Background())

	if err := InsertOrderPayment(transactions, payment); err != nil {
		return err
	}

	if err := InsertOrderDelivery(transactions, delivery); err != nil {
		return err
	}

	if err := InsertOrder(transactions, order); err != nil {
		return err
	}

	if err := InsertOrderItems(transactions, items); err != nil {
		return err
	}

	if err := transactions.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func InsertOrderPayment(transactions *utils.Transactions, payment *model.OrderPaymentModel) error {
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

func InsertOrderDelivery(transactions *utils.Transactions, delivery *model.OrderDeliveryModel) error {
	query := `
        INSERT INTO delivery (id, name, phone, zip, city, address, region, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_ = transactions.QueryRow(
		context.Background(),
		query,
		delivery.Id,
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

func InsertOrder(transactions *utils.Transactions, order *model.OrderModel) error {
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
		order.Delivery.Id,
		order.Payment.Id,
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

func InsertOrderItems(transactions *utils.Transactions, items *[]model.OrderItemModel) error {
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
			item.OrderId,
		)
	}

	return nil
}

func CreateTables(pool *pgxpool.Pool) {
	CreateOrderPaymentTable(pool)
	CreateOrderDeliveryTable(pool)
	CreateOrderTable(pool)
	CreateOrderItems(pool)
}

func CreateOrderTable(pool *pgxpool.Pool) {
	query := `
    CREATE TABLE IF NOT EXISTS orders (
        order_uid UUID PRIMARY KEY,
        track_number VARCHAR(255),
        entry VARCHAR(255),
        delivery_id UUID REFERENCES delivery(id),
        payment_id UUID REFERENCES payment(id),
        locale VARCHAR(255),
        internal_signature VARCHAR(255),
        customer_id VARCHAR(255),
        delivery_service VARCHAR(255),
        shardkey INT,
        sm_id INT,
        date_created TIMESTAMP,
        oof_shard VARCHAR(255)
    );
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Fatalf("Error creating tables: %v\n", err)
	}
}

func CreateOrderDeliveryTable(pool *pgxpool.Pool) {
	query := `
    CREATE TABLE IF NOT EXISTS delivery (
        id UUID PRIMARY KEY,
        name VARCHAR(255),
        phone VARCHAR(255),
        zip VARCHAR(255),
        city VARCHAR(255),
        address VARCHAR(255),
        region VARCHAR(255),
        email VARCHAR(255)
    );
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Fatalf("Error creating tables: %v\n", err)
	}
}

func CreateOrderPaymentTable(pool *pgxpool.Pool) {
	query := `
    CREATE TABLE IF NOT EXISTS payment (
        id UUID PRIMARY KEY,
        transaction UUID,
        request_id VARCHAR(255),
        currency VARCHAR(255),
        provider VARCHAR(255),
        amount INT,
        payment_dt TIMESTAMP,
        bank VARCHAR(255),
        delivery_cost INT,
        goods_total INT,
        custom_fee INT
    );
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Fatalf("Error creating tables: %v\n", err)
	}
}

func CreateOrderItems(pool *pgxpool.Pool) {
	query := `
	CREATE TABLE IF NOT EXISTS order_items (
		chrt_id INT PRIMARY KEY,
		track_number VARCHAR(255),
		price INT,
		rid UUID,
		name VARCHAR(255),
		sale INT,
		size VARCHAR(255),
		total_price INT,
		nm_id INT,
		brand VARCHAR(255),
		status INT,
		order_id UUID REFERENCES orders(order_uid)
	);
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Fatalf("Error creating tables: %v\n", err)
	}
}
