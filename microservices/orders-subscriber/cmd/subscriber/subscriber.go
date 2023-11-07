package subscriber

import (
	"context"
	natsSubscriber "github.com/emptyhopes/orders-subscriber/cmd/nats-subscriber"
	"github.com/emptyhopes/orders-subscriber/internal/provider"
	orderProvider "github.com/emptyhopes/orders-subscriber/internal/provider/orders"
	"github.com/joho/godotenv"
)

type Application struct {
	orderProvider provider.OrderProviderInterface
}

func NewApplication(ctx context.Context) (*Application, error) {
	application := &Application{}

	err := application.initializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return application, nil
}

func (a *Application) initializeDependency(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initializeEnvironment,
		a.initializeProvider,
	}

	for _, function := range inits {
		err := function(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) initializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (a *Application) initializeProvider(_ context.Context) error {
	a.orderProvider = orderProvider.NewOrderProvider()

	return nil
}

func (a *Application) Run() {
	natsSubscriber.Start(a.orderProvider.GetOrderController())
}
