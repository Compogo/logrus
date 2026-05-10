package metrics

import (
	"reflect"

	"github.com/Compogo/compogo/configurator"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/logrus/infrastructure/link"
	"github.com/Compogo/types/set"
	"github.com/sirupsen/logrus"
)

const (
	LevelNamesFieldName = "logger.metric.levels"
)

var LevelNamesDefault = []string{logger.Panic.String(), logger.Error.String(), logger.Warn.String()}

type Config struct {
	levelNames []string
	Levels     set.Set[logrus.Level]
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config, configurator configurator.Configurator) (*Config, error) {
	if len(config.levelNames) == 0 || reflect.DeepEqual(config.levelNames, LevelNamesDefault) {
		configurator.SetDefault(LevelNamesFieldName, LevelNamesDefault)
		config.levelNames = configurator.GetStringSlice(LevelNamesFieldName)
	}

	for _, levelName := range config.levelNames {
		level, err := logger.AllLevels.Get(levelName)
		if err != nil {
			return nil, err
		}

		logrusLevel, err := link.LoggerLevelToLogrusLevel.Get(level)
		if err != nil {
			return nil, err
		}

		config.Levels.Add(logrusLevel)
	}

	return config, nil
}
