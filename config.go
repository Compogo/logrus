package logrus

import (
	"strings"

	"github.com/Compogo/compogo/configurator"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/types/linker"
	"github.com/sirupsen/logrus"
)

const (
	// LevelNameFieldName defines the command-line flag name for setting the log level.
	// Example: --logger.level=debug
	LevelNameFieldName = "logger.level"
)

var (
	// LevelNameDefault defines the default log level as a string.
	// It corresponds to logger.Error (lowercase "error").
	LevelNameDefault = strings.ToLower(logger.Error.String())

	// LoggerLevelToLogrusLevel maps Compogo's logger.Level to logrus.Level.
	// This ensures type-safe conversion between the two level systems.
	LoggerLevelToLogrusLevel = linker.NewLinker[logger.Level, logrus.Level](
		linker.NewLink(logger.Panic, logrus.PanicLevel),
		linker.NewLink(logger.Error, logrus.ErrorLevel),
		linker.NewLink(logger.Warn, logrus.WarnLevel),
		linker.NewLink(logger.Info, logrus.InfoLevel),
		linker.NewLink(logger.Debug, logrus.DebugLevel),
	)
)

// Config holds the logger configuration that can be set via command-line flags
// or configuration files. It includes the log level as both string and parsed enum.
type Config struct {
	// LevelName is the string representation of the log level (e.g., "info", "debug").
	// It is populated from the command-line flag.
	LevelName string

	// Level is the parsed logger.Level enum value, converted from LevelName.
	Level logger.Level
}

// NewConfig creates a new Config instance with default values.
// The actual configuration will be applied later via Configuration function
// and command-line flag binding.
func NewConfig() *Config {
	return &Config{}
}

// Configuration applies configuration values to the Config struct.
// It reads from the configurator, sets defaults, and validates the log level.
// Returns an error if the level name cannot be parsed into a valid logger.Level.
//
// The function is designed to be used with container.Invoke in the PreRun phase.
func Configuration(config *Config, configurator configurator.Configurator) (*Config, error) {
	if config.LevelName == "" || config.LevelName == LevelNameDefault {
		configurator.SetDefault(LevelNameFieldName, LevelNameDefault)
		config.LevelName = configurator.GetString(LevelNameFieldName)
	}

	var err error
	config.Level, err = logger.Levels.Get(config.LevelName)
	if err != nil {
		return nil, err
	}

	return config, nil
}
