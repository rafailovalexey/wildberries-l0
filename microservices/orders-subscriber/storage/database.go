package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

type DatabaseInterface interface {
	Initialize()
	GetPool() *pgxpool.Pool
	CreateTables(*pgxpool.Pool)
	CreateOrderTable(*pgxpool.Pool)
	CreateOrderDeliveryTable(*pgxpool.Pool)
	CreateOrderPaymentTable(*pgxpool.Pool)
	CreateOrderItems(*pgxpool.Pool)
}

type database struct {
	credentials string
}

var _ DatabaseInterface = &database{}

func NewDatabase() *database {
	d := &database{}

	d.Initialize()

	pool := d.GetPool()
	defer pool.Close()

	d.CreateTables(pool)

	return d
}

func (d *database) Initialize() {
	username := os.Getenv("POSTGRESQL_USERNAME")

	if username == "" {
		log.Fatalf("укажите пользователя базы данных")
	}

	password := os.Getenv("POSTGRESQL_PASSWORD")

	if password == "" {
		log.Fatalf("укажите пользователя базы данных")
	}

	database := os.Getenv("POSTGRESQL_DATABASE")

	if database == "" {
		log.Fatalf("укажите название базы данных")
	}

	hostname := os.Getenv("POSTGRESQL_HOSTNAME")

	if hostname == "" {
		log.Fatalf("укажите имя хоста базы данных")
	}

	port := os.Getenv("POSTGRESQL_PORT")

	if port == "" {
		log.Fatalf("укажите порт базы данных")
	}

	sslmode := os.Getenv("POSTGRESQL_SSLMODE")

	if sslmode == "" {
		log.Fatalf("укажите ssl mode базы данных")
	}

	d.credentials = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", username, password, database, hostname, port, sslmode)
}

func (d *database) GetPool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), d.credentials)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	return pool
}

func (d *database) CreateTables(pool *pgxpool.Pool) {
	d.CreateOrderPaymentTable(pool)
	d.CreateOrderDeliveryTable(pool)
	d.CreateOrderTable(pool)
	d.CreateOrderItems(pool)
}

func (d *database) CreateOrderTable(pool *pgxpool.Pool) {
	query := `
    CREATE TABLE IF NOT EXISTS orders (
        order_uid UUID DEFAULT gen_random_uuid() PRIMARY KEY,
        track_number VARCHAR(255),
        entry VARCHAR(255),
        delivery_uid UUID REFERENCES orders_delivery(delivery_uid),
        payment_uid UUID REFERENCES orders_payment(payment_uid),
        locale VARCHAR(255),
        internal_signature VARCHAR(255),
        customer_id VARCHAR(255),
        delivery_service VARCHAR(255),
        shardkey VARCHAR(255),
        sm_id INT,
        date_created TIMESTAMP,
        oof_shard VARCHAR(255)
    );
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Fatalf("ошибка создания таблицы: %v\n", err)
	}
}

func (d *database) CreateOrderDeliveryTable(pool *pgxpool.Pool) {
	query := `
    CREATE TABLE IF NOT EXISTS orders_delivery (
        delivery_uid UUID DEFAULT gen_random_uuid() PRIMARY KEY,
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
		log.Fatalf("ошибка создания таблицы: %v\n", err)
	}
}

func (d *database) CreateOrderPaymentTable(pool *pgxpool.Pool) {
	query := `
    CREATE TABLE IF NOT EXISTS orders_payment (
        payment_uid UUID DEFAULT gen_random_uuid() PRIMARY KEY,
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
		log.Fatalf("ошибка создания таблицы: %v\n", err)
	}
}

func (d *database) CreateOrderItems(pool *pgxpool.Pool) {
	query := `
	CREATE TABLE IF NOT EXISTS orders_items (
		chrt_id SERIAL PRIMARY KEY,
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
		order_uid UUID REFERENCES orders(order_uid)
	);
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Fatalf("ошибка создания таблицы: %v\n", err)
	}
}
