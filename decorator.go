package logrus

import (
	"fmt"
	"os"

	"github.com/Compogo/compogo/logger"
	"github.com/sirupsen/logrus"
)

// Decorator wraps a logrus.Logger to implement the logger.Logger interface.
// It adds application name prefixing to all log messages and supports
// creating child loggers for sub-components.
//
// The decorator maintains separate stdout and stderr loggers, allowing
// different outputs for different log levels if needed.
type Decorator struct {
	// appName is the name of the current application or component.
	// It is prefixed to all log messages.
	appName string

	// stdErr is the logrus logger used for error-level and above messages.
	stdErr *logrus.Logger

	// stdOut is the logrus logger used for info-level and below messages.
	stdOut *logrus.Logger

	// parent points to the parent logger when this is a child logger
	// created via GetLogger. It enables log message delegation up the chain.
	parent *Decorator
}

// NewDecorator creates a new Decorator instance with default logrus loggers.
// Both stdout and stderr loggers are initialized with default settings.
// The actual log level will be set later via SetLevel.
func NewDecorator() *Decorator {
	decorator := &Decorator{
		stdErr: logrus.New(),
		stdOut: logrus.New(),
	}

	decorator.stdErr.SetOutput(os.Stderr)
	decorator.stdOut.SetOutput(os.Stdout)

	formatter := &logrus.TextFormatter{}

	decorator.stdErr.SetFormatter(formatter)
	decorator.stdOut.SetFormatter(formatter)

	return decorator
}

// Panicf logs a formatted message at Panic level and then panics.
// The message is prefixed with the application name in brackets.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Panicf(s string, i ...interface{}) {
	if logger.appName != "" {
		s = fmt.Sprintf("[%s] %s", logger.appName, s)
	}

	if logger.parent != nil {
		logger.parent.Panicf(s, i...)
		return
	}

	logger.stdErr.Panicf(s, i...)
}

// Panic logs a message at Panic level and then panics.
// The application name is prepended to the arguments.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Panic(i ...interface{}) {
	if logger.appName != "" {
		i = append([]interface{}{fmt.Sprintf("[%s] ", logger.appName)}, i...)
	}

	if logger.parent != nil {
		logger.parent.Panic(i...)
		return
	}

	logger.stdErr.Panic(i...)
}

// Errorf logs a formatted message at Error level.
// The message is prefixed with the application name in brackets.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Errorf(s string, i ...interface{}) {
	if logger.appName != "" {
		s = fmt.Sprintf("[%s] %s", logger.appName, s)
	}

	if logger.parent != nil {
		logger.parent.Errorf(s, i...)
		return
	}

	logger.stdErr.Errorf(s, i...)
}

// Error logs a message at Error level.
// The application name is prepended to the arguments.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Error(i ...interface{}) {
	if logger.appName != "" {
		i = append([]interface{}{fmt.Sprintf("[%s] ", logger.appName)}, i...)
	}

	if logger.parent != nil {
		logger.parent.Error(i...)
		return
	}

	logger.stdErr.Error(i...)
}

// Warnf logs a formatted message at Warn level.
// The message is prefixed with the application name in brackets.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Warnf(s string, i ...interface{}) {
	if logger.appName != "" {
		s = fmt.Sprintf("[%s] %s", logger.appName, s)
	}

	if logger.parent != nil {
		logger.parent.Warnf(s, i...)
		return
	}

	logger.stdErr.Warnf(s, i...)
}

// Warn logs a message at Warn level.
// The application name is prepended to the arguments.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Warn(i ...interface{}) {
	if logger.appName != "" {
		i = append([]interface{}{fmt.Sprintf("[%s] ", logger.appName)}, i...)
	}

	if logger.parent != nil {
		logger.parent.Warn(i...)
		return
	}

	logger.stdErr.Warn(i...)
}

// Infof logs a formatted message at Info level.
// The message is prefixed with the application name in brackets.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Infof(s string, i ...interface{}) {
	if logger.appName != "" {
		s = fmt.Sprintf("[%s] %s", logger.appName, s)
	}

	if logger.parent != nil {
		logger.parent.Infof(s, i...)
		return
	}

	logger.stdOut.Infof(s, i...)
}

// Info logs a message at Info level.
// The application name is prepended to the arguments.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Info(i ...interface{}) {
	if logger.appName != "" {
		i = append([]interface{}{fmt.Sprintf("[%s] ", logger.appName)}, i...)
	}

	if logger.parent != nil {
		logger.parent.Info(i...)
		return
	}

	logger.stdOut.Info(i...)
}

// Debugf logs a formatted message at Debug level.
// The message is prefixed with the application name in brackets.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Debugf(s string, i ...interface{}) {
	if logger.appName != "" {
		s = fmt.Sprintf("[%s] %s", logger.appName, s)
	}

	if logger.parent != nil {
		logger.parent.Debugf(s, i...)
		return
	}

	logger.stdOut.Debugf(s, i...)
}

// Debug logs a message at Debug level.
// The application name is prepended to the arguments.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Debug(i ...interface{}) {
	if logger.appName != "" {
		i = append([]interface{}{fmt.Sprintf("[%s] ", logger.appName)}, i...)
	}

	if logger.parent != nil {
		logger.parent.Debug(i...)
		return
	}

	logger.stdOut.Debug(i...)
}

// Printf logs a formatted message at Info level (for compatibility).
// The message is prefixed with the application name in brackets.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Printf(s string, i ...interface{}) {
	if logger.appName != "" {
		s = fmt.Sprintf("[%s] %s", logger.appName, s)
	}

	if logger.parent != nil {
		logger.parent.Printf(s, i...)
		return
	}

	logger.stdOut.Printf(s, i...)
}

// Print logs a message at Info level (for compatibility).
// The application name is prepended to the arguments.
// If this is a child logger, the call is delegated to the parent.
func (logger *Decorator) Print(i ...interface{}) {
	if logger.appName != "" {
		i = append([]interface{}{fmt.Sprintf("[%s] ", logger.appName)}, i...)
	}

	if logger.parent != nil {
		logger.parent.Print(i...)
		return
	}

	logger.stdOut.Print(i...)
}

// GetStdErr returns the underlying logrus.Logger used for error output.
// This can be used for advanced configuration or adding hooks.
func (logger *Decorator) GetStdErr() *logrus.Logger {
	return logger.stdErr
}

// GetStdOut returns the underlying logrus.Logger used for standard output.
// This can be used for advanced configuration or adding hooks.
func (logger *Decorator) GetStdOut() *logrus.Logger {
	return logger.stdOut
}

// SetLevel changes the logging level for both stdout and stderr loggers.
// The level is converted from Compogo's logger.Level to logrus.Level
// using the LoggerLevelToLogrusLevel mapper.
// Returns an error if the level conversion fails.
func (logger *Decorator) SetLevel(level logger.Level) error {
	logerusLevel, err := LoggerLevelToLogrusLevel.Get(level)
	if err != nil {
		return err
	}

	logger.stdOut.SetLevel(logerusLevel)
	logger.stdErr.SetLevel(logerusLevel)

	return nil
}

// AddHook adds a logrus hook to both stdout and stderr loggers.
// This enables integration with external logging systems like Sentry,
// file logging, or custom formatters.
func (logger *Decorator) AddHook(hook logrus.Hook) {
	logger.stdOut.AddHook(hook)
	logger.stdErr.AddHook(hook)
}

// GetLogger creates a child logger with the given name.
// The child logger will prefix all messages with its own name
// in addition to the parent's name, creating a chain like:
// [app] [child] message
//
// This is useful for sub-components that need their own log identity
// while still being part of the main application.
func (logger *Decorator) GetLogger(name string) logger.Logger {
	return &Decorator{
		appName: name,
		parent:  logger,
	}
}
