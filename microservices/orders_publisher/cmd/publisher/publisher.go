package publisher

import (
	"context"
	natsPublisher "github.com/emptyhopes/orders_publisher/cmd/nats_publisher"
	"github.com/emptyhopes/orders_publisher/internal/provider"
	orderProvider "github.com/emptyhopes/orders_publisher/internal/provider/orders"
	"github.com/joho/godotenv"
)

type PublisherInterface interface {
	InitializeDependency(ctx context.Context) error
	InitializeEnvironment(_ context.Context) error
	InitializeProvider(_ context.Context) error
	Run()
}

type publisher struct {
	orderProvider provider.OrderProviderInterface
}

var _ PublisherInterface = (*publisher)(nil)

func NewPublisher(ctx context.Context) (*publisher, error) {
	application := &publisher{}

	err := application.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return application, nil
}

func (a *publisher) InitializeDependency(ctx context.Context) error {
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

func (a *publisher) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (a *publisher) InitializeProvider(_ context.Context) error {
	a.orderProvider = orderProvider.NewOrderProvider()

	return nil
}

func (a *publisher) Run() {
	controller := a.orderProvider.GetOrderController()

	natsPublisher.Start(controller)
}
