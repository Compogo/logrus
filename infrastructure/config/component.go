package config

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
)

// Component — компонент конфигурации логгера.
// Добавляет флаг для уровня логирования и загружает конфигурацию.
var Component = &compogo.Component{
	Name: "logger.Logrus.config",
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.levelName, LevelNameFieldName, LevelNameDefault, "level on logger")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
}
