package hook

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
	"github.com/Compogo/logrus"
)

var LogrusMetricComponent = &component.Component{
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provides(
			NewConfig,
			NewMetricHook,
		)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringArrayVar(&config.LevelNames, LevelNamesFieldName, LevelNamesDefault, "")
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
	PreExecute: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(decorator logrus.Decorator, hook *MetricHook) {
			decorator.AddHook(hook)
		})
	}),
}
