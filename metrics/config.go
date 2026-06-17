package metrics

import (
	"reflect"

	"github.com/Compogo/compogo"
	"github.com/Compogo/logrus/infrastructure/link"
	"github.com/Compogo/types/set"
	"github.com/sirupsen/logrus"
)

// LevelNamesFieldName — имя поля в конфигурации для уровней метрик.
const LevelNamesFieldName = "logger.metric.levels"

// LevelNamesDefault — уровни по умолчанию: Panic, Error, Warn.
var LevelNamesDefault = []string{compogo.Panic.String(), compogo.Error.String(), compogo.Warn.String()}

// Config содержит конфигурацию метрик логгера.
type Config struct {
	levelNames []string
	Levels     set.Set[logrus.Level]
}

// NewConfig создаёт новую конфигурацию метрик.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию метрик из Configurator.
// Парсит имена уровней в значения Logrus.
func Configuration(config *Config, configurator compogo.Configurator) (*Config, error) {
	if len(config.levelNames) == 0 || reflect.DeepEqual(config.levelNames, LevelNamesDefault) {
		configurator.SetDefault(LevelNamesFieldName, LevelNamesDefault)
		config.levelNames = configurator.GetStringSlice(LevelNamesFieldName)
	}

	for _, levelName := range config.levelNames {
		level, err := compogo.AllLevels.Get(levelName)
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
