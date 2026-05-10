package config

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
)

var Component = &component.Component{
	Name: "logger.Logrus.config",
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.levelName, LevelNameFieldName, LevelNameDefault, "level on logger")
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
}
