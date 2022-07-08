package config

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewEnvConfiguration); err != nil {
		return errors.Wrap(err, "Failed to provide env configuration")
	}
	return nil
}
