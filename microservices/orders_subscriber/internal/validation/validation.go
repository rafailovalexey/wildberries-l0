package validation

import dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"

type OrderValidationInterface interface {
	HandleOrderMessageValidation(*dto.OrderDto) error
}
