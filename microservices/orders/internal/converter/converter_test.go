package converter

import (
	mockConverter "github.com/emptyhopes/orders/internal/converter/mocks"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	model "github.com/emptyhopes/orders/internal/model/orders"
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

	orderDto := getSubOrderDto()
	deliveryUid := ""
	paymentUid := ""
	orderModel := getSubOrderModel()

	order := converter.EXPECT().MapOrderDtoToOrderModel(orderDto, deliveryUid, paymentUid).Return(orderModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderPaymentDtoToOrderPaymentModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderPaymentDto := getSubOrderPaymentDto()
	orderPaymentModel := getSubOrderPaymentModel()

	order := converter.EXPECT().MapOrderPaymentDtoToOrderPaymentModel(orderPaymentDto).Return(orderPaymentModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderPaymentModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderDeliveryDtoToOrderDeliveryModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderDeliveryDto := getSubOrderDeliveryDto()
	orderDeliveryModel := getSubOrderDeliveryModel()

	order := converter.EXPECT().MapOrderDeliveryDtoToOrderDeliveryModel(orderDeliveryDto).Return(orderDeliveryModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderDeliveryModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemDtoToOrderItemModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemDto := getSubOrderItemDto()
	orderItemModel := getSubOrderItemModel()
	orderUid := ""

	order := converter.EXPECT().MapOrderItemDtoToOrderItemModel(orderItemDto, orderUid).Return(orderItemModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderItemModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemsDtoToOrderItemsModel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemsDto := getSubOrderItemsDto()
	orderItemsModel := getSubOrderItemsModel()
	orderUid := ""

	order := converter.EXPECT().MapOrderItemsDtoToOrderItemsModel(orderItemsDto, orderUid).Return(orderItemsModel)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderItemsModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderModelToOrderDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderModel := getSubOrderItemsModel()
	orderDeliveryModel := getSubOrderDeliveryModel()
	orderPaymentModel := getSubOrderPaymentModel()
	orderItemsModel := getSubOrderItemsModel()
	orderDto := getSubOrderItemsDto()

	order := converter.EXPECT().MapOrderModelToOrderDto(orderModel, orderDeliveryModel, orderPaymentModel, orderItemsModel).Return(orderDto)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderPaymentModelToOrderPaymentDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderPaymentModel := getSubOrderPaymentModel()
	orderPaymentDto := getSubOrderPaymentDto()

	order := converter.EXPECT().MapOrderPaymentModelToOrderPaymentDto(orderPaymentModel).Return(orderPaymentDto)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderPaymentModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderDeliveryModelToOrderDeliveryDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderDeliveryModel := getSubOrderDeliveryModel()
	orderDeliveryDto := getSubOrderDeliveryDto()

	order := converter.EXPECT().MapOrderDeliveryModelToOrderDeliveryDto(orderDeliveryModel).Return(orderDeliveryDto)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderDeliveryModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemModelToOrderItemDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemModel := getSubOrderItemModel()
	orderItemDto := getSubOrderItemDto()

	order := converter.EXPECT().MapOrderItemModelToOrderItemDto(orderItemModel).Return(orderItemDto)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderItemModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
}

func TestConverterMapOrderItemsModelToOrderItemsDto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	converter := mockConverter.NewMockOrderConverterInterface(ctrl)

	orderItemsModel := getSubOrderItemsModel()
	orderItemsDto := getSubOrderItemsDto()

	order := converter.EXPECT().MapOrderItemsModelToOrderItemsDto(orderItemsModel).Return(orderItemsDto)

	if !reflect.TypeOf(order).AssignableTo(reflect.TypeOf(&model.OrderItemsModel{})) {
		t.Errorf("order has the wrong type")
	}

	require.NotNil(t, order)
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

func getSubOrderModel() *model.OrderModel {
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

func getSubOrderDeliveryModel() *model.OrderDeliveryModel {
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

func getSubOrderPaymentModel() *model.OrderPaymentModel {
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

func getSubOrderItemModel() *model.OrderItemModel {
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

func getSubOrderItemsModel() *model.OrderItemsModel {
	item1 := getSubOrderItemModel()
	item2 := getSubOrderItemModel()
	item3 := getSubOrderItemModel()
	item4 := getSubOrderItemModel()

	items := model.NewOrderItemsModel(
		item1,
		item2,
		item3,
		item4,
	)

	return items
}
