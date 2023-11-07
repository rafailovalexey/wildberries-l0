package service

import (
	"errors"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	mockService "github.com/emptyhopes/orders/internal/service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestServiceGetOrderByIdWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockService.NewMockOrderServiceInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"

	service.EXPECT().GetOrderById(id).Return(nil, errors.New("error"))
	order, err := service.GetOrderById(id)

	require.Nil(t, order)
	require.Error(t, err)
}

func TestServiceGetOrderByIdWithoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockService.NewMockOrderServiceInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"
	orderDto := getSubOrderDto()

	service.EXPECT().GetOrderById(id).Return(orderDto, nil)
	order, err := service.GetOrderById(id)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("order has the wrong type")
	}

	require.Nil(t, err)
}

func getSubOrderDeliveryDto() *dto.OrderDeliveryDto {
	delivery := dto.NewOrderDeliveryDto(
		"Test Testov",
		"+9720000000",
		"2639809",
		"Kiryat Mozkin",
		"Ploshad Mira 15",
		"Kraiot",
		"test@gmail.com",
	)

	return delivery
}

func getSubOrderPaymentDto() *dto.OrderPaymentDto {
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

	return payment
}

func getSubOrderItemDto() *dto.OrderItemDto {
	item := dto.NewOrderItemDto(
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

	return item
}

func getSubOrderItemsDto() *dto.OrderItemsDto {
	item1 := getSubOrderItemDto()
	item2 := getSubOrderItemDto()
	item3 := getSubOrderItemDto()
	item4 := getSubOrderItemDto()

	items := dto.NewOrderItemsDto(
		item1,
		item2,
		item3,
		item4,
	)

	return items
}

func getSubOrderDto() *dto.OrderDto {
	delivery := getSubOrderDeliveryDto()
	payment := getSubOrderPaymentDto()
	items := getSubOrderItemsDto()

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
