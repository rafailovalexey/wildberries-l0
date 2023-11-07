package application

import (
	"context"
	httpServer "github.com/emptyhopes/orders/cmd/http_server"
	"github.com/emptyhopes/orders/internal/provider"
	orderProvider "github.com/emptyhopes/orders/internal/provider/orders"
	"github.com/joho/godotenv"
)

type ApplicationInterface interface {
	InitializeDependency(ctx context.Context) error
	InitializeEnvironment(_ context.Context) error
	InitializeProvider(_ context.Context) error
	Run()
}

type application struct {
	orderProvider provider.OrderProviderInterface
}

var _ ApplicationInterface = (*application)(nil)

func NewApplication(ctx context.Context) (*application, error) {
	application := &application{}

	err := application.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return application, nil
}

func (a *application) InitializeDependency(ctx context.Context) error {
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

func (a *application) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (a *application) InitializeProvider(_ context.Context) error {
	a.orderProvider = orderProvider.NewOrderProvider()

	return nil
}

func (a *application) Run() {
	api := a.orderProvider.GetOrderApi()

	httpServer.Run(api)
}
