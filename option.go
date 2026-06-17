package logrus

import (
	"context"

	"github.com/Compogo/compogo"
	"github.com/Compogo/logrus/infrastructure/config"
)

var (
	// logger — синглтон логгера, используемый во всём приложении.
	logger = NewLogger()

	// component — компонент Logrus для Compogo.
	// Регистрирует логгер в DI-контейнере и настраивает его уровень.
	component = compogo.Component{
		Name: "logger.Logrus",
		Dependencies: compogo.Components{
			config.Component,
		},
		Init: compogo.StepFunc(func(container compogo.Container) error {
			return container.Provides(
				func() *Logger { return logger },
				func(logger *Logger) compogo.Panicer { return logger },
				func(logger *Logger) compogo.Errorer { return logger },
				func(logger *Logger) compogo.Warner { return logger },
				func(logger *Logger) compogo.Informer { return logger },
				func(logger *Logger) compogo.Debuger { return logger },
				func(logger *Logger) compogo.Printer { return logger },
				func(logger *Logger) compogo.Logger { return logger },
			)
		}),
		Configuration: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger *Logger, config *config.Config) error {
				return logger.SetLevel(config.Level)
			})
		}),
		PreExecute: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger *Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PreExecute' step.")
				return nil
			})
		}),
		Execute: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'Execute' step.")
				return nil
			})
		}),
		PostExecute: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PostExecute' step.")
				return nil
			})
		}),
		PreWait: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PreWait' step.")
				return nil
			})
		}),
		Wait: compogo.WaitFunc(func(_ context.Context, container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'Wait' step.")
				return nil
			})
		}),
		PostWait: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PostWait' step.")
				return nil
			})
		}),
		PreStop: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PreStop' step.")
				return nil
			})
		}),
		Stop: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'Stop' step.")
				return nil
			})
		}),
		PostStop: compogo.StepFunc(func(container compogo.Container) error {
			return container.Invoke(func(logger compogo.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PostStop' step.")
				return nil
			})
		}),
	}
)

// WithLogrus возвращает опцию для подключения Logrus к приложению Compogo.
//
// Пример:
//
//	app := compogo.NewApp("myapp",
//	    logrus.WithLogrus(),
//	    // другие опции...
//	)
func WithLogrus() compogo.Option {
	l := NewLogger()

	return compogo.WithLogger(l, &component)
}
