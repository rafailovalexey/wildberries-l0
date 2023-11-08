package api

import (
	"encoding/json"
	mockApi "github.com/emptyhopes/orders/internal/api/mocks"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestApiGetOrderByIdWithIdIsNotUuid(t *testing.T) {
	ctrl := gomock.NewController(t)

	controller := mockApi.NewMockOrderApiInterface(ctrl)

	request := httptest.NewRequest("GET", "/v1/orders/test", nil)
	response := httptest.NewRecorder()

	controller.EXPECT().GetOrderById(response, request).Do(getResponseWithIdIsNotUuid)
	controller.GetOrderById(response, request)

	require.Equal(t, response.Code, http.StatusBadRequest)
	require.Equal(t, response.Body.String(), "id is not uuid")
}

func TestApiGetOrderByIdWithoutError(t *testing.T) {
	ctrl := gomock.NewController(t)

	controller := mockApi.NewMockOrderApiInterface(ctrl)

	request := httptest.NewRequest("GET", "/v1/orders/d4bf115a-95b8-430a-a1d2-5d21fd8c8905", nil)
	response := httptest.NewRecorder()

	controller.EXPECT().GetOrderById(response, request).Do(getResponseWithoutError)
	controller.GetOrderById(response, request)

	order := &dto.OrderDto{}

	err := json.Unmarshal(response.Body.Bytes(), order)

	if err != nil {
		t.Error(err)
	}

	require.Equal(t, response.Code, http.StatusOK)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("order has the wrong type")
	}
}

func getResponseWithIdIsNotUuid(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte("id is not uuid"))
}

func getResponseWithoutError(response http.ResponseWriter, _ *http.Request) {
	orderDto := getStubOrderDto()

	orderJson, err := json.Marshal(orderDto)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("произошла ошибка при взаимодействие с объектом"))

		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(orderJson)
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
