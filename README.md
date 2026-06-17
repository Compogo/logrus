# Compogo Logrus

[![Go Reference](https://pkg.go.dev/badge/github.com/Compogo/logrus.svg)](https://pkg.go.dev/github.com/Compogo/logrus)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Адаптер [Logrus](https://github.com/sirupsen/logrus) для фреймворка [Compogo](https://github.com/Compogo/compogo).

Реализует интерфейс `compogo.Logger` через Logrus с поддержкой:

* Иерархических логгеров (`GetLogger`)
* Метрик Prometheus (счётчики по уровням логирования)
* Гибкой конфигурации через Compogo
* Разделения stdout/stderr

## Установка

```shell
go get github.com/Compogo/logrus
```

## Быстрый старт

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/logrus"
)

func main() {
    app := compogo.NewApp("myapp",
        logrus.WithLogrus(),
        compogo.WithConfigurator(configurator, configuratorCmp),
        compogo.WithContainer(container, containerCmp),
        compogo.WithOsSignalCloser(),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## Использование

### Базовое логирование

```go
// Логирование с разными уровнями
logger.Info("Application started")
logger.Warn("Memory usage is high")
logger.Error("Failed to connect to database")
logger.Debug("Request details", "path", "/api/v1")
```

### Иерархические логгеры

```go
// Создание вложенных логгеров
httpLogger := logger.GetLogger("http")
dbLogger := logger.GetLogger("database")

httpLogger.Info("Server listening on :8080") // "[http] Server listening on :8080"
dbLogger.Info("Connected to PostgreSQL")     // "[database] Connected to PostgreSQL"

// Вложенность: compogo -> myapp -> http
appLogger := logger.GetLogger("compogo").GetLogger("myapp")
appLogger.Info("Starting") // "[compogo][myapp] Starting"
```

## Конфигурация через флаги

```shell
# Уровень логирования
go run main.go --logger.level=debug

# Метрики для каких уровней собирать
go run main.go --logger.metric.levels=panic,error,warn,info
```

## Метрики Prometheus

Пакет предоставляет готовый компонент метрики для счётчиков логов по уровням.

#### Метрика:

```plantuml
compogo_log_level_count{app="myapp",level="error"} 5
compogo_log_level_count{app="myapp",level="warn"} 12
compogo_log_level_count{app="myapp",level="info"} 42
```

#### Настройка через конфигурацию:
```go
// По умолчанию собираются только panic, error, warn
// Можно настроить через флаг или конфиг
configurator.SetDefault("logger.metric.levels", []string{"panic", "error", "warn", "info"})
```

## API

### Logger

Реализует интерфейс `compogo.Logger`:

```go
type Logger interface {
    // Уровни логирования
    Panicf(string, ...interface{})
    Panic(...interface{})
    Errorf(string, ...interface{})
    Error(...interface{})
    Warnf(string, ...interface{})
    Warn(...interface{})
    Infof(string, ...interface{})
    Info(...interface{})
    Debugf(string, ...interface{})
    Debug(...interface{})
    Printf(string, ...interface{})
    Print(...interface{})

    // Иерархия
    GetLogger(name string) Logger
}
```

### Дополнительные методы

```go
// Установка уровня логирования
logger.SetLevel(compogo.Debug) error

// Добавление своего хука Logrus
logger.AddHook(hook logrus.Hook)

// Доступ к внутренним логгерам
logger.GetStdOut() *logrus.Logger
logger.GetStdErr() *logrus.Logger
```

## Зависимости

* [Compogo](https://github.com/Compogo/compogo) — основной фреймворк
* [Logrus](https://github.com/sirupsen/logrus) — библиотека логирования
* [Prometheus](https://github.com/prometheus/client_golang) — метрики

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
