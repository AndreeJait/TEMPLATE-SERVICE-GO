package migration

import (
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/logs"
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/persistent/mongo"
	mm "github.com/AndreeJait/GO-ANDREE-UTILITIES/util/migration"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/config"
	nosqlmigration "github.com/AndreeJait/TEMPLATE-SERVICE-GO/migration/nosql"

	"github.com/pkg/errors"
)

func NewNoSqlMigration(conf *config.EnvConfiguration, mongo mongo.Mongo, logger logs.Logger) (mm.Tool, error) {
	nosql, err := mm.NewNoSqlMigration(mongo, nosqlmigration.Script, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate no sql migration tool")
	}
	return nosql, nil
}
