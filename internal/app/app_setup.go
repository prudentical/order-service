package app

import (
	"context"
	"fmt"
	"log/slog"
	"order-service/configuration"
	"order-service/internal/message"

	"go.uber.org/fx"
)

type AppSetupManager interface {
	Setup() error
	Shutdown() error
}
type appSetupManagerImpl struct {
	app      RESTApp
	messages message.MessageHandler
}

func NewAppSetupManager(app RESTApp, messages message.MessageHandler) AppSetupManager {
	return appSetupManagerImpl{app, messages}
}

func (a appSetupManagerImpl) Setup() error {
	err := a.app.setup()
	if err != nil {
		return err
	}
	err = a.messages.RegisterMessagesListener()
	return err
}

func (a appSetupManagerImpl) Shutdown() error {
	return nil
}

func ManageLifeCycle(lc fx.Lifecycle, config configuration.Config, log *slog.Logger, app RESTApp, manager AppSetupManager) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting the server")
			err := manager.Setup()
			if err != nil {
				return err
			}
			go app.server().Start(fmt.Sprintf(":%v", config.Server.Port))
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shuting down the server")
			err := manager.Shutdown()
			if err != nil {
				return err
			}
			return app.server().Shutdown(ctx)
		},
	})
}
