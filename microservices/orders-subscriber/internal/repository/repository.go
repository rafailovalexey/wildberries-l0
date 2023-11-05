package repository

import (
	"github.com/emptyhopes/orders-subscriber/internal/helpers"
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
	"github.com/emptyhopes/orders-subscriber/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Cache = storage.ConstructorCache()

func init() {
	database := &storage.Database{}

	database.Initialize()

	pool := database.GetPool()
	defer pool.Close()

	database.CreateTables(pool)
}

type OrdersRepositoryInterface interface {
	Cache(*model.OrderModel) (bool, error)
	CreateOrder(pool *pgxpool.Pool, order *model.OrderModel, delivery *model.OrderDeliveryModel, payment *model.OrderPaymentModel, items *[]model.OrderItemModel) error
	insertOrderPayment(transactions *helpers.Transactions, payment *model.OrderPaymentModel) error
	insertOrderDelivery(transactions *helpers.Transactions, delivery *model.OrderDeliveryModel) error
	insertOrder(transactions *helpers.Transactions, order *model.OrderModel) error
	insertOrderItems(transactions *helpers.Transactions, items *[]model.OrderItemModel) error
}
