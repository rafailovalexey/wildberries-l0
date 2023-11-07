package orders

import (
	dto "github.com/emptyhopes/orders_publisher/internal/dto/orders"
	definition "github.com/emptyhopes/orders_publisher/internal/repository"
	"github.com/google/uuid"
	"time"
)

type repository struct{}

var _ definition.OrderRepositoryInterface = (*repository)(nil)

func NewOrderRepository() *repository {
	return &repository{}
}

func (r *repository) GetOrder() *dto.OrderDto {
	delivery := dto.NewOrderDeliveryDto(
		"Test Testov",
		"+9720000000",
		"2639809",
		"Kiryat Mozkin",
		"Ploshad Mira 15",
		"Kraiot",
		"test@gmail.com",
	)
	payment := dto.NewOrderPaymentDto(
		uuid.New().String(),
		"1",
		"USD",
		"wbpay",
		1817,
		time.Now().Unix(),
		"alpha",
		1500,
		317,
		0,
	)
	item1 := dto.NewOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	item2 := dto.NewOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	item3 := dto.NewOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	item4 := dto.NewOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	items := dto.NewOrderItemsDto(
		item1,
		item2,
		item3,
		item4,
	)

	order := dto.NewOrderDto(
		uuid.New().String(),
		"WBILMTESTTRACK",
		"WBIL",
		delivery,
		payment,
		items,
		"en",
		"1",
		"test",
		"meest",
		"9",
		99,
		time.Now().Unix(),
		"1",
	)

	return order
}
