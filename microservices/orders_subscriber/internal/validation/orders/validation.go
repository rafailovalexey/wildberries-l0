package orders

import (
	"errors"
	"fmt"
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	definition "github.com/emptyhopes/orders_subscriber/internal/validation"
	"regexp"
)

type validation struct{}

var _ definition.OrderValidationInterface = (*validation)(nil)

func NewOrderValidation() *validation {
	return &validation{}
}

func (v *validation) HandleOrderMessageValidation(dto *dto.OrderDto) error {
	if err := itemsIsNotEmpty(dto.Items); err != nil {
		return err
	}

	if err := isValidUuid(dto.OrderUid, "order_uid"); err != nil {
		return err
	}

	if err := isValidUuid(dto.Payment.Transaction, "payment.transaction"); err != nil {
		return err
	}

	for index, item := range *dto.Items {
		if err := isValidUuid(item.Rid, fmt.Sprintf("items[%d].rid", index)); err != nil {
			return err
		}
	}

	return nil
}

func isValidUuid(id string, field string) error {
	result := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`).MatchString(id)

	if !result {
		return errors.New(fmt.Sprintf("%s is not uuid", field))
	}

	return nil
}

func itemsIsNotEmpty(items *dto.OrderItemsDto) error {
	if len(*items) == 0 {
		return errors.New("items is not empty")
	}

	return nil
}
