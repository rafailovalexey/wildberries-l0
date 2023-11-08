package repository

import (
	"errors"
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	mockRepository "github.com/emptyhopes/orders_subscriber/internal/repository/mocks"
	"github.com/emptyhopes/orders_subscriber/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestRepositoryGetOrdersCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	cache := storage.NewCache()

	cacheItems := cache.GetCache()

	repository.EXPECT().GetOrdersCache().Return(cacheItems)
	cacheItemsWithRepository := repository.GetOrdersCache()

	if !reflect.TypeOf(cacheItemsWithRepository).AssignableTo(reflect.TypeOf(&map[string]storage.CacheItem{})) {
		t.Errorf("cacheItemsWithRepository has the wrong type")
	}
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
	orderDto := getStubOrderDto()

	repository.EXPECT().SetOrderCache(id, orderDto).Return()
	repository.SetOrderCache(id, orderDto)

	repository.EXPECT().GetOrderCacheById(id).Return(orderDto, true)
	orderCached, isExist := repository.GetOrderCacheById(id)

	if !reflect.TypeOf(orderCached).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("orderCached has the wrong type")
	}

	require.True(t, isExist)
}

func TestRepositorySetOrderCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"
	orderDto := getStubOrderDto()

	repository.EXPECT().SetOrderCache(id, orderDto).Return()
	repository.SetOrderCache(id, orderDto)

	repository.EXPECT().GetOrderCacheById(id).Return(orderDto, true)
	orderCached, isExist := repository.GetOrderCacheById(id)

	if !reflect.TypeOf(orderCached).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("orderCached has the wrong type")
	}

	require.True(t, isExist)
}

func TestRepositoryDeleteOrderCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"
	orderDto := getStubOrderDto()

	repository.EXPECT().SetOrderCache(id, orderDto).Return()
	repository.SetOrderCache(id, orderDto)

	repository.EXPECT().GetOrderCacheById(id).Return(orderDto, true)
	orderCached, isExist := repository.GetOrderCacheById(id)

	if !reflect.TypeOf(orderCached).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("orderCached has the wrong type")
	}

	require.True(t, isExist)

	repository.EXPECT().DeleteOrderCacheById(id).Return()
	repository.DeleteOrderCacheById(id)

	repository.EXPECT().GetOrderCacheById(id).Return(nil, false)
	orderCachedAfterDelete, isExist := repository.GetOrderCacheById(id)

	require.Nil(t, orderCachedAfterDelete)
	require.False(t, isExist)
}

func TestRepositoryCreateOrderWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	orderDto := getStubOrderDto()

	repository.EXPECT().CreateOrder(orderDto).Return(errors.New(""))
	err := repository.CreateOrder(orderDto)

	require.Error(t, err)
}

func TestRepositoryCreateOrderWithoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	orderDto := getStubOrderDto()

	repository.EXPECT().CreateOrder(orderDto).Return(nil)
	err := repository.CreateOrder(orderDto)

	require.Nil(t, err)
}

func getStubOrderDeliveryDto() *dto.OrderDeliveryDto {
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

func getStubOrderPaymentDto() *dto.OrderPaymentDto {
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

func getStubOrderItemDto() *dto.OrderItemDto {
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

func getStubOrderItemsDto() *dto.OrderItemsDto {
	item1 := getStubOrderItemDto()
	item2 := getStubOrderItemDto()
	item3 := getStubOrderItemDto()
	item4 := getStubOrderItemDto()

	items := dto.NewOrderItemsDto(
		item1,
		item2,
		item3,
		item4,
	)

	return items
}

func getStubOrderDto() *dto.OrderDto {
	delivery := getStubOrderDeliveryDto()
	payment := getStubOrderPaymentDto()
	items := getStubOrderItemsDto()

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
