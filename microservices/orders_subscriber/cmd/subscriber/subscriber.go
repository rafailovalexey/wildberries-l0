package subscriber

import (
	"context"
	natsSubscriber "github.com/emptyhopes/orders_subscriber/cmd/nats_subscriber"
	"github.com/emptyhopes/orders_subscriber/internal/provider"
	orderProvider "github.com/emptyhopes/orders_subscriber/internal/provider/orders"
	"github.com/joho/godotenv"
)

type SubscriberInterface interface {
	InitializeDependency(ctx context.Context) error
	InitializeEnvironment(_ context.Context) error
	InitializeProvider(_ context.Context) error
	Run()
}

type subscriber struct {
	orderProvider provider.OrderProviderInterface
}

var _ SubscriberInterface = (*subscriber)(nil)

func NewSubscriber(ctx context.Context) (*subscriber, error) {
	application := &subscriber{}

	err := application.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return application, nil
}

func (a *subscriber) InitializeDependency(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.InitializeEnvironment,
		a.InitializeProvider,
	}

	for _, function := range inits {
		err := function(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *subscriber) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (a *subscriber) InitializeProvider(_ context.Context) error {
	a.orderProvider = orderProvider.NewOrderProvider()

	return nil
}

func (a *subscriber) Run() {
	controller := a.orderProvider.GetOrderController()

	natsSubscriber.Start(controller)
}
