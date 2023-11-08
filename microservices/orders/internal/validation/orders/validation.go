package orders

import (
	"errors"
	definition "github.com/emptyhopes/orders/internal/validation"
	"regexp"
)

type validation struct{}

var _ definition.OrderValidationInterface = (*validation)(nil)

func NewOrderValidation() *validation {
	return &validation{}
}

func (v *validation) GetOrderByIdValidation(id string) error {
	if !isValidUuid(id) {
		return errors.New("id is not uuid")
	}

	return nil
}

func isValidUuid(id string) bool {
	return regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`).MatchString(id)
}
