package validation

type OrderValidationInterface interface {
	GetOrderByIdValidation(id string) error
}
