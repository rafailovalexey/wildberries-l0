package orders

import (
	"fmt"
	model "github.com/emptyhopes/orders/internal/model/orders"
	"github.com/emptyhopes/orders/internal/repository"
	"time"
)

type Repository struct{}

var _ repository.OrdersRepositoryInterface = &Repository{}

func (r *Repository) GetOrderById(id string) (*model.OrderModel, error) {
	value, isExist := repository.Cache.Get(id)
	fmt.Println(value, isExist)

	if !isExist {
		fmt.Println("set value")
		repository.Cache.Set(id, "1", 5*time.Minute)
	}

	return nil, nil
}
