package logrus

import (
	"context"

	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
	"github.com/Compogo/compogo/logger"
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
	decorator := NewDecorator()

	return compogo.WithLogger(decorator, &component.Component{
		Name: "logger.Logrus",
		Init: component.StepFunc(func(container container.Container) error {
			return container.Provides(
				NewConfig,
				func() *Decorator { return decorator },
				func(decorator *Decorator) logger.Panicer { return decorator },
				func(decorator *Decorator) logger.Errorer { return decorator },
				func(decorator *Decorator) logger.Warner { return decorator },
				func(decorator *Decorator) logger.Informer { return decorator },
				func(decorator *Decorator) logger.Debuger { return decorator },
				func(decorator *Decorator) logger.Printer { return decorator },
				func(decorator *Decorator) logger.Logger { return decorator },
			)
		}),
		BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
			return container.Invoke(func(config *Config) {
				flagSet.StringVar(&config.LevelName, LevelNameFieldName, LevelNameDefault, "level on logger")
			})
		}),
		Configuration: component.StepFunc(func(container container.Container) error {
			return container.Invoke(Configuration)
		}),
		PreExecute: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(decorator *Decorator, config *Config, appCfg *compogo.Config) error {
				decorator.appName = appCfg.Name
				if err := decorator.SetLevel(config.Level); err != nil {
					return err
				}

				decorator.Info("[logrus] execute 'PreExecute' step.")

				return nil
			})
		}),
		Execute: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'Execute' step.")
				return nil
			})
		}),
		PostExecute: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'PostExecute' step.")
				return nil
			})
		}),
		PreWait: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'PreWait' step.")
				return nil
			})
		}),
		Wait: component.WaitFunc(func(_ context.Context, container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'Wait' step.")
				return nil
			})
		}),
		PostWait: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'PostWait' step.")
				return nil
			})
		}),
		PreStop: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'PreStop' step.")
				return nil
			})
		}),
		Stop: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'Stop' step.")
				return nil
			})
		}),
		PostStop: component.StepFunc(func(container container.Container) error {
			return container.Invoke(func(logger logger.Logger) error {
				logger.Info("[logrus] execute 'PostStop' step.")
				return nil
			})
		}),
	})
}
