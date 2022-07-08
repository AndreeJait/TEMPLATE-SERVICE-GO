package di

import (
	"context"
	"sync"

	"github.com/AndreeJait/GO-ANDREE-UTILITIES/logs"
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/persistent/mongo"
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/util/validator"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/application"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/config"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/infrastructure"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/interfaces"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/interfaces/controller"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/migration"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/shared"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

func Container() (*dig.Container, error) {
	var outer error

	once.Do(func() {
		container = dig.New()

		// External dep

		if err := container.Provide(ProvideMongo); err != nil {
			outer = err
			return
		}

		if err := container.Provide(ProvideLogger); err != nil {
			outer = err
			return
		}

		if err := container.Provide(NewValidator); err != nil {
			outer = err
			return
		}

		if err := container.Provide(shared.ProvideCloser); err != nil {
			outer = err
			return
		}

		// Internal dep
		if err := config.Register(container); err != nil {
			outer = err
			return
		}

		if err := application.Register(container); err != nil {
			outer = err
			return
		}

		if err := infrastructure.Register(container); err != nil {
			outer = err
			return
		}

		if err := interfaces.Register(container); err != nil {
			outer = err
			return
		}

		if err := controller.Register(container); err != nil {
			outer = err
			return
		}

		if err := shared.Register(container); err != nil {
			outer = err
			return
		}

		if err := migration.Register(container); err != nil {
			outer = err
			return
		}

	})

	if outer != nil {
		return nil, errors.Wrap(outer, "failed to initialize container")
	}

	return container, nil
}

func ProvideLogger(config *config.EnvConfiguration) (logs.Logger, error) {
	return logs.New(&logs.Option{
		LogFilePath: config.LoggerFileName,
		Formatter:   logs.Formatter(config.LoggerFormater),
		Level:       logs.Level(config.LoggerLevel),
		MaxSize:     config.LoggerMaxSize,
		MaxBackups:  config.LoggerMaxBackups,
		MaxAge:      config.LoggerMaxAge,
		Compress:    config.LoggerCompress,
	})
}

func NewValidator() validator.Validator {
	return validator.New()
}

func ProvideMongo(config *config.EnvConfiguration, log logs.Logger) (mongo.Mongo, error) {
	return mongo.New(context.Background(), config.MongoDatabaseUri, config.MongoDatabaseName, log)
}
