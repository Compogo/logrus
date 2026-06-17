package metrics

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
	"github.com/Compogo/logrus"
)

// LogrusMetricComponent — компонент метрик для Logrus.
// Добавляет флаги для настройки метрик и устанавливает хук в логгер.
var LogrusMetricComponent = compogo.Component{
	Name: "logger.Logrus.metrics",
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provides(
			NewConfig,
			NewMetricHook,
		)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringArrayVar(&config.levelNames, LevelNamesFieldName, LevelNamesDefault, "")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
	PreExecute: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(decorator *logrus.Logger, hook *MetricHook) {
			decorator.AddHook(hook)
		})
	}),
}
