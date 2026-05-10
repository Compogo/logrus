package config

import (
	"strings"

	"github.com/Compogo/compogo/configurator"
	"github.com/Compogo/compogo/logger"
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
)

// Config holds the logger configuration that can be set via command-line flags
// or configuration files. It includes the log level as both string and parsed enum.
type Config struct {
	// levelName is the string representation of the log level (e.g., "info", "debug").
	// It is populated from the command-line flag.
	levelName string

	// Level is the parsed logger.Level enum value, converted from levelName.
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
	if config.levelName == "" || config.levelName == LevelNameDefault {
		configurator.SetDefault(LevelNameFieldName, LevelNameDefault)
		config.levelName = configurator.GetString(LevelNameFieldName)
	}

	var err error
	config.Level, err = logger.AllLevels.Get(config.levelName)
	if err != nil {
		return nil, err
	}

	return config, nil
}
