package interfaces

import (
	"github.com/Del-Gong/DELGONG-BE-STORY/interfaces/auth"
	"github.com/Del-Gong/DELGONG-BE-STORY/interfaces/story"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(auth.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide Authentication")
	}

	if err := container.Provide(story.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide Story")
	}
	return nil
}
