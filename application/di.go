package application

import (
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
	}
)

func Register(container *dig.Container) error {

	// Register application here

	return nil
}
