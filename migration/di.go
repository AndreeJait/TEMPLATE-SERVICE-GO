package migration

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewNoSqlMigration); err != nil {
		return errors.Wrap(err, "failed to provide NoSql migration")
	}
	return nil
}
