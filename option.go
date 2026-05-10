package logrus

import (
	"context"

	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/logrus/infrastructure/config"
)

// WithLogrus returns a Compogo option that integrates logrus as the
// application's logging system. It automatically:
//   - Registers the logger in the DI container
//   - Adds command-line flags for log level configuration
//   - Sets up the log level from configuration
//   - Injects the application name into the logger
//
// Usage:
//
//	app := compogo.NewApp("myapp",
//	    logrus.WithLogrus(),
//	    // other options...
//	)
func WithLogrus() compogo.Option {
	l := NewLogger()

	return compogo.WithLogger(l, &component.Component{
		Name: "logger.Logrus",
		Dependencies: component.Components{
			config.Component,
		},
		Init: component.StepFunc(func(container container.Container) error {
			return container.Provides(
				func() *Logger { return l },
				func(logger *Logger) logger.Panicer { return logger },
				func(logger *Logger) logger.Errorer { return logger },
				func(logger *Logger) logger.Warner { return logger },
				func(logger *Logger) logger.Informer { return logger },
				func(logger *Logger) logger.Debuger { return logger },
				func(logger *Logger) logger.Printer { return logger },
				func(logger *Logger) logger.Logger { return logger },
			)
		}),
		Configuration: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger *Logger, config *config.Config) error {
				return logger.SetLevel(config.Level)
			})
		}),
		PreExecute: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger *Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PreExecute' step.")
				return nil
			})
		}),
		Execute: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'Execute' step.")
				return nil
			})
		}),
		PostExecute: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PostExecute' step.")
				return nil
			})
		}),
		PreWait: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PreWait' step.")
				return nil
			})
		}),
		Wait: component.WaitFunc(func(_ context.Context, container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'Wait' step.")
				return nil
			})
		}),
		PostWait: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PostWait' step.")
				return nil
			})
		}),
		PreStop: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PreStop' step.")
				return nil
			})
		}),
		Stop: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'Stop' step.")
				return nil
			})
		}),
		PostStop: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger, compogoCfg *compogo.Config) error {
				logger.GetLogger("compogo").GetLogger(compogoCfg.Name).Info("execute 'PostStop' step.")
				return nil
			})
		}),
	})
}
