package config

import (
	"strings"

	"github.com/Compogo/compogo"
)

// LevelNameFieldName — имя поля в конфигурации для уровня логирования.
const LevelNameFieldName = "logger.level"

var (
	// LevelNameDefault — значение по умолчанию для уровня логирования (Error).
	LevelNameDefault = strings.ToLower(compogo.Error.String())
)

// Config содержит конфигурацию логгера.
type Config struct {
	levelName string
	Level     compogo.Level
}

// NewConfig создаёт новую конфигурацию.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию из Configurator и парсит уровень логирования.
// Если уровень не задан, используется Error.
//
// Пример:
//
//	config, err := Configuration(&Config{}, configurator)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Configuration(config *Config, configurator compogo.Configurator) (*Config, error) {
	if config.levelName == "" || config.levelName == LevelNameDefault {
		configurator.SetDefault(LevelNameFieldName, LevelNameDefault)
		config.levelName = configurator.GetString(LevelNameFieldName)
	}

	var err error
	config.Level, err = compogo.AllLevels.Get(config.levelName)
	if err != nil {
		return nil, err
	}

	return config, nil
}
