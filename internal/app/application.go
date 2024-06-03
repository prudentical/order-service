package app

import (
	"order-service/internal/api"
	"order-service/internal/configuration"
	"order-service/internal/database"
	"order-service/internal/discovery"
	"order-service/internal/message"
	"order-service/internal/persistence"
	"order-service/internal/service"
	"order-service/internal/util"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type Application interface {
	Run()
}

func NewApplication() Application {
	return FxContainer{}
}

type FxContainer struct {
}

func (FxContainer) Run() {
	fx.New(
		fx.Provide(configuration.NewConfig),
		fx.Provide(NewLogger),
		fx.Provide(NewFxLogger),
		fx.Provide(ProvideEcho),
		fx.Provide(NewAppSetupManager),
		fx.Provide(discovery.NewServiceDiscovery),
		fx.Provide(util.NewValidator),
		fx.Provide(api.NewHTTPErrorHandler),
		fx.Provide(database.NewDatabaseConnection),
		fx.Provide(message.NewMessageQueueClient),
		fx.Provide(persistence.NewOrderDAO),
		fx.Provide(persistence.NewPositionDAO),
		fx.Provide(service.NewOrderService),
		fx.Provide(service.NewPositionService),
		fx.Provide(message.NewMessageHandler),
		asHandler(api.NewHealthCheck),
		asHandler(api.NewPositionHandler),
		fx.Provide(fx.Annotate(
			NewRestApp,
			fx.ParamTags(`group:"handlers"`),
		)),
		fx.WithLogger(func(log FxLogger) fxevent.Logger {
			return &log
		}),
		fx.Invoke(ManageLifeCycle),
	).Run()
}

func asHandler(handler interface{}) fx.Option {
	return fx.Provide(fx.Annotate(
		handler,
		fx.As(new(api.Handler)),
		fx.ResultTags(`group:"handlers"`),
	))
}
