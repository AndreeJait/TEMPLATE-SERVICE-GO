package shared

import "go.uber.org/dig"

type GlobalHolder struct {
	dig.In
}

func Register(container *dig.Container) error {

	return nil
}
