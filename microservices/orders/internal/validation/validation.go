package validation

type OrderValidationInterface interface {
	GetOrderByIdValidation(string) error
}
