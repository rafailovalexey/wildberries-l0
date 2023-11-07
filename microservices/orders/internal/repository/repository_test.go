package repository

import (
	"errors"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	mockRepository "github.com/emptyhopes/orders/internal/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestRepositorySetOrderCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"
	order := getSubOrderDto()

	repository.EXPECT().SetOrderCache(id, order).Return()
	repository.SetOrderCache(id, order)

	repository.EXPECT().GetOrderCacheById(id).Return(order, true)
	orderCached, isExist := repository.GetOrderCacheById(id)

	if !reflect.TypeOf(orderCached).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("orderCached has the wrong type")
	}

	require.True(t, isExist)
}

func TestRepositoryGetOrderCacheByIdWithUncachedId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"

	repository.EXPECT().GetOrderCacheById(id).Return(nil, false)

	order, isExist := repository.GetOrderCacheById(id)

	require.Nil(t, order)
	require.False(t, isExist)
}

func TestRepositoryGetOrderCacheByIdWithCachedId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"
	order := getSubOrderDto()

	repository.EXPECT().SetOrderCache(id, order).Return()
	repository.SetOrderCache(id, order)

	repository.EXPECT().GetOrderCacheById(id).Return(order, true)
	orderCached, isExist := repository.GetOrderCacheById(id)

	if !reflect.TypeOf(orderCached).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("orderCached has the wrong type")
	}

	require.True(t, isExist)
}

func TestRepositoryGetOrderByIdWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"

	repository.EXPECT().GetOrderById(id).Return(nil, errors.New("no rows expected"))
	orderDatabase, err := repository.GetOrderById(id)

	require.Nil(t, orderDatabase)
	require.Error(t, err, "no rows expected")
}

func TestRepositoryGetOrderByIdWithoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"
	order := getSubOrderDto()

	repository.EXPECT().GetOrderById(id).Return(order, nil)
	orderDatabase, err := repository.GetOrderById(id)

	if !reflect.TypeOf(orderDatabase).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("orderDatabase has the wrong type")
	}

	require.Nil(t, err)
}

func getSubOrderDto() *dto.OrderDto {
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
		*item1,
		*item2,
		*item3,
		*item4,
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
