package logrus

import (
	"os"

	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/logrus/infrastructure/link"
	"github.com/Compogo/types/linker"
	"github.com/sirupsen/logrus"
)

// Logger wraps a logrus.Logger to implement the logger.Logger interface.
// It adds application name prefixing to all log messages and supports
// creating child loggers for sub-components.
//
// The decorator maintains separate stdout and stderr loggers, allowing
// different outputs for different log levels if needed.
type Logger struct {
	name string

	// stdErr is the logrus logger used for error-level and above messages.
	stdErr *logrus.Logger

	// stdOut is the logrus logger used for info-level and below messages.
	stdOut *logrus.Logger

	// parent points to the parent logger when this is a child logger
	// created via GetLogger. It enables log message delegation up the chain.
	parent *Logger

	childes *linker.Linker[string, *Logger]
}

// NewLogger creates a new Logger instance with default logrus loggers.
// Both stdout and stderr loggers are initialized with default settings.
// The actual log level will be set later via SetLevel.
func NewLogger() *Logger {
	l := &Logger{
		stdErr: logrus.New(),
		stdOut: logrus.New(),
	}

	l.stdErr.SetOutput(os.Stderr)
	l.stdOut.SetOutput(os.Stdout)

	formatter := &logrus.TextFormatter{}

	l.stdErr.SetFormatter(formatter)
	l.stdOut.SetFormatter(formatter)

	return l
}

// Panicf logs a formatted message at Panic level and then panics.
// The message is prefixed with the application name in brackets.
// If this is a child logger, the call is delegated to the parent.
func (logger *Logger) Panicf(s string, i ...interface{}) {
	if logger.name != "" {
		s = "[%s]" + s
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = logger.name
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
func (logger *Logger) Panic(i ...interface{}) {
	if logger.name != "" {
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = "[" + logger.name + "]"
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
func (logger *Logger) Errorf(s string, i ...interface{}) {
	if logger.name != "" {
		s = "[%s]" + s
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = logger.name
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
func (logger *Logger) Error(i ...interface{}) {
	if logger.name != "" {
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = "[" + logger.name + "]"
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
func (logger *Logger) Warnf(s string, i ...interface{}) {
	if logger.name != "" {
		s = "[%s]" + s
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = logger.name
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
func (logger *Logger) Warn(i ...interface{}) {
	if logger.name != "" {
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = "[" + logger.name + "]"
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
func (logger *Logger) Infof(s string, i ...interface{}) {
	if logger.name != "" {
		s = "[%s]" + s
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = logger.name
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
func (logger *Logger) Info(i ...interface{}) {
	if logger.name != "" {
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = "[" + logger.name + "]"
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
func (logger *Logger) Debugf(s string, i ...interface{}) {
	if logger.name != "" {
		s = "[%s]" + s
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = logger.name
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
func (logger *Logger) Debug(i ...interface{}) {
	if logger.name != "" {
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = "[" + logger.name + "]"
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
func (logger *Logger) Printf(s string, i ...interface{}) {
	if logger.name != "" {
		s = "[%s]" + s
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = logger.name
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
func (logger *Logger) Print(i ...interface{}) {
	if logger.name != "" {
		i = append(i, 0)
		copy(i[1:], i[:len(i)-1])
		i[0] = "[" + logger.name + "]"
	}

	if logger.parent != nil {
		logger.parent.Print(i...)
		return
	}

	logger.stdOut.Print(i...)
}

// GetStdErr returns the underlying logrus.Logger used for error output.
// This can be used for advanced configuration or adding hooks.
func (logger *Logger) GetStdErr() *logrus.Logger {
	return logger.stdErr
}

// GetStdOut returns the underlying logrus.Logger used for standard output.
// This can be used for advanced configuration or adding hooks.
func (logger *Logger) GetStdOut() *logrus.Logger {
	return logger.stdOut
}

// SetLevel changes the logging level for both stdout and stderr loggers.
// The level is converted from Compogo's logger.Level to logrus.Level
// using the LoggerLevelToLogrusLevel mapper.
// Returns an error if the level conversion fails.
func (logger *Logger) SetLevel(level logger.Level) error {
	logrusLevel, err := link.LoggerLevelToLogrusLevel.Get(level)
	if err != nil {
		return err
	}

	logger.stdOut.SetLevel(logrusLevel)
	logger.stdErr.SetLevel(logrusLevel)

	return nil
}

// AddHook adds a logrus hook to both stdout and stderr loggers.
// This enables integration with external logging systems like Sentry,
// file logging, or custom formatters.
func (logger *Logger) AddHook(hook logrus.Hook) {
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
func (logger *Logger) GetLogger(name string) logger.Logger {
	if logger.childes == nil {
		logger.childes = linker.NewLinker[string, *Logger]()
	}

	if !logger.childes.Has(name) {
		logger.childes.Add(name, &Logger{
			name:   name,
			parent: logger,
		})
	}

	l, _ := logger.childes.Get(name)
	return l
}
