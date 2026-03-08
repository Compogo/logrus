package logrus

import (
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
		Init: component.StepFunc(func(container container.Container) error {
			return container.Provides(
				NewConfig(),
				func() *Decorator { return decorator },
				func(decorator *Decorator) logger.Logger { return decorator },
			)
		}),
		BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
			return container.Invoke(func(config *Config) {
				flagSet.StringVar(&config.LevelName, LevelNameFieldName, LevelNameDefault, "level on logger")
			})
		}),
		PreRun: component.StepFunc(func(container container.Container) error {
			if err := container.Invoke(Configuration); err != nil {
				return err
			}

			return container.Invoke(func(decorator *Decorator, config *Config, appCfg *compogo.Config) error {
				decorator.appName = appCfg.Name
				return decorator.SetLevel(config.Level)
			})
		}),
	})
}
