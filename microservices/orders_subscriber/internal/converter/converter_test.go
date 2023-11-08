package converter

import (
	mockConverter "github.com/emptyhopes/orders_subscriber/internal/converter/mocks"
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	model "github.com/emptyhopes/orders_subscriber/internal/model/orders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestConverterMapOrderDtoToOrderModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderDto := getStubOrderDto()
	deliveryUid := ""
	paymentUid := ""
	orderModel := getStubOrderModel()

	converter.EXPECT().MapOrderDtoToOrderModel(orderDto, deliveryUid, paymentUid).Return(orderModel)
	order := converter.MapOrderDtoToOrderModel(orderDto, deliveryUid, paymentUid)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderPaymentDtoToOrderPaymentModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderPaymentDto := getStubOrderPaymentDto()
	orderPaymentModel := getStubOrderPaymentModel()

	converter.EXPECT().MapOrderPaymentDtoToOrderPaymentModel(orderPaymentDto).Return(orderPaymentModel)
	order := converter.MapOrderPaymentDtoToOrderPaymentModel(orderPaymentDto)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderPaymentModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderDeliveryDtoToOrderDeliveryModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderDeliveryDto := getStubOrderDeliveryDto()
	orderDeliveryModel := getStubOrderDeliveryModel()

	converter.EXPECT().MapOrderDeliveryDtoToOrderDeliveryModel(orderDeliveryDto).Return(orderDeliveryModel)
	order := converter.MapOrderDeliveryDtoToOrderDeliveryModel(orderDeliveryDto)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderDeliveryModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemDtoToOrderItemModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemDto := getStubOrderItemDto()
	orderItemModel := getStubOrderItemModel()
	orderUid := ""

	converter.EXPECT().MapOrderItemDtoToOrderItemModel(orderItemDto, orderUid).Return(orderItemModel)
	order := converter.MapOrderItemDtoToOrderItemModel(orderItemDto, orderUid)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderItemModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemsDtoToOrderItemsModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemsDto := getStubOrderItemsDto()
	orderItemsModel := getStubOrderItemsModel()
	orderUid := ""

	converter.EXPECT().MapOrderItemsDtoToOrderItemsModel(orderItemsDto, orderUid).Return(orderItemsModel)
	order := converter.MapOrderItemsDtoToOrderItemsModel(orderItemsDto, orderUid)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderItemsModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderModelToOrderDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderModel := getStubOrderModel()
	orderDeliveryModel := getStubOrderDeliveryModel()
	orderPaymentModel := getStubOrderPaymentModel()
	orderItemsModel := getStubOrderItemsModel()
	orderDto := getStubOrderDto()

	converter.EXPECT().MapOrderModelToOrderDto(orderModel, orderDeliveryModel, orderPaymentModel, orderItemsModel).Return(orderDto)
	order := converter.MapOrderModelToOrderDto(orderModel, orderDeliveryModel, orderPaymentModel, orderItemsModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&dto.OrderDto{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderPaymentModelToOrderPaymentDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderPaymentModel := getStubOrderPaymentModel()
	orderPaymentDto := getStubOrderPaymentDto()

	converter.EXPECT().MapOrderPaymentModelToOrderPaymentDto(orderPaymentModel).Return(orderPaymentDto)
	order := converter.MapOrderPaymentModelToOrderPaymentDto(orderPaymentModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&dto.OrderPaymentDto{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderDeliveryModelToOrderDeliveryDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderDeliveryModel := getStubOrderDeliveryModel()
	orderDeliveryDto := getStubOrderDeliveryDto()

	converter.EXPECT().MapOrderDeliveryModelToOrderDeliveryDto(orderDeliveryModel).Return(orderDeliveryDto)
	order := converter.MapOrderDeliveryModelToOrderDeliveryDto(orderDeliveryModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&dto.OrderDeliveryDto{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemModelToOrderItemDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemModel := getStubOrderItemModel()
	orderItemDto := getStubOrderItemDto()

	converter.EXPECT().MapOrderItemModelToOrderItemDto(orderItemModel).Return(orderItemDto)
	order := converter.MapOrderItemModelToOrderItemDto(orderItemModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&dto.OrderItemDto{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemsModelToOrderItemsDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemsModel := getStubOrderItemsModel()
	orderItemsDto := getStubOrderItemsDto()

	converter.EXPECT().MapOrderItemsModelToOrderItemsDto(orderItemsModel).Return(orderItemsDto)
	order := converter.MapOrderItemsModelToOrderItemsDto(orderItemsModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&dto.OrderItemsDto{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
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

func getStubOrderModel() *model.OrderModel {
	order := model.NewOrderModel(
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		0,
		time.Now(),
		"",
	)

	return order
}

func getStubOrderDeliveryModel() *model.OrderDeliveryModel {
	delivery := model.NewOrderDeliveryModel(
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	)

	return delivery
}

func getStubOrderPaymentModel() *model.OrderPaymentModel {
	payment := model.NewOrderPaymentModel(
		"",
		"",
		"",
		"",
		"",
		0,
		time.Now(),
		"",
		0,
		0,
		0,
	)

	return payment
}

func getStubOrderItemModel() *model.OrderItemModel {
	item := model.NewOrderItemModel(
		0,
		"",
		0,
		"",
		"",
		0,
		"",
		0,
		0,
		"",
		0,
		"",
	)

	return item
}

func getStubOrderItemsModel() *model.OrderItemsModel {
	item1 := getStubOrderItemModel()
	item2 := getStubOrderItemModel()
	item3 := getStubOrderItemModel()
	item4 := getStubOrderItemModel()

	items := model.NewOrderItemsModel(
		item1,
		item2,
		item3,
		item4,
	)

	return items
}
