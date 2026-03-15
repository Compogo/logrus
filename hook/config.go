package hook

import (
	"reflect"

	"github.com/Compogo/compogo/configurator"
	"github.com/Compogo/compogo/logger"
	logrus2 "github.com/Compogo/logrus"
	"github.com/Compogo/types/set"
	"github.com/sirupsen/logrus"
)

const (
	LevelNamesFieldName = "logger.metric.levels"
)

var LevelNamesDefault = []string{logger.Panic.String(), logger.Error.String(), logger.Warn.String()}

type Config struct {
	LevelNames []string
	Levels     set.Set[logrus.Level]
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config, configurator configurator.Configurator) (*Config, error) {
	if len(config.LevelNames) == 0 || reflect.DeepEqual(config.LevelNames, LevelNamesDefault) {
		configurator.SetDefault(LevelNamesFieldName, LevelNamesDefault)
		config.LevelNames = configurator.GetStringSlice(LevelNamesFieldName)
	}

	for _, levelName := range config.LevelNames {
		level, err := logger.Levels.Get(levelName)
		if err != nil {
			return nil, err
		}

		logrusLevel, err := logrus2.LoggerLevelToLogrusLevel.Get(level)
		if err != nil {
			return nil, err
		}

		config.Levels.Add(logrusLevel)
	}

	return config, nil
}
