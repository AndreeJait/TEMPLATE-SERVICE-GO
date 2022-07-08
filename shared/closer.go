package shared

import (
	"context"

	"github.com/AndreeJait/GO-ANDREE-UTILITIES/logs"
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/persistent/mongo"
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/util/validator"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
		Mongo     mongo.Mongo
		Logger    logs.Logger
		Validator validator.Validator
	}
	Closer struct {
		Mongo  mongo.Mongo
		Logger logs.Logger
	}
)

func (c *Closer) Close() {
	if err := c.Mongo.Client().Disconnect(context.Background()); err != nil {
		c.Logger.Errorf("failed to close mongodb connection %s", err)
	}
}

func ProvideCloser(holder Holder) (*Closer, error) {
	return &Closer{
		Mongo:  holder.Mongo,
		Logger: holder.Logger,
	}, nil
}
