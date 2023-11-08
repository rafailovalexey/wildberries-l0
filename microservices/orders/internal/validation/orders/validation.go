package orders

import (
	"errors"
	"fmt"
	definition "github.com/emptyhopes/orders/internal/validation"
	"regexp"
)

type validation struct{}

var _ definition.OrderValidationInterface = (*validation)(nil)

func NewOrderValidation() *validation {
	return &validation{}
}

func (v *validation) GetOrderByIdValidation(id string) error {
	if err := isValidUuid(id, "order_uid"); err != nil {
		return err
	}

	return nil
}

func isValidUuid(id string, field string) error {
	result := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`).MatchString(id)

	if result {
		return errors.New(fmt.Sprintf("%s is not uuid", field))
	}

	return nil
}
