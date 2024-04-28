package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"order-service/internal/api"
	"order-service/internal/app"
	"order-service/internal/configuration"
	"order-service/internal/database"
	"order-service/internal/discovery"
	"order-service/internal/message"
	"order-service/internal/persistence"
	"order-service/internal/service"
	"order-service/internal/util"
)

func main() {
	fx.New(
		fx.Provide(configuration.NewConfig),
		fx.Provide(app.NewLogger),
		fx.Provide(app.NewFxLogger),
		fx.Provide(app.ProvideEcho),
		fx.Provide(app.NewAppSetupManager),
		fx.Provide(discovery.NewServiceDiscovery),
		fx.Provide(util.NewValidator),
		fx.Provide(api.NewHTTPErrorHandler),
		fx.Provide(database.NewDatabaseConnection),
		fx.Provide(message.NewMessageQueueClient),
		fx.Provide(persistence.NewOrderDAO),
		fx.Provide(service.NewOrderService),
		fx.Provide(message.NewMessageHandler),
		asHandler(api.NewHealthCheck),
		asHandler(api.NewOrderHandler),
		fx.Provide(fx.Annotate(
			app.NewRestApp,
			fx.ParamTags(`group:"handlers"`),
		)),
		fx.WithLogger(func(log app.FxLogger) fxevent.Logger {
			return &log
		}),
		fx.Invoke(app.ManageLifeCycle),
	).Run()
}

func asHandler(handler interface{}) fx.Option {
	return fx.Provide(fx.Annotate(
		handler,
		fx.As(new(api.Handler)),
		fx.ResultTags(`group:"handlers"`),
	))
}
