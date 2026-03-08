# Compogo Logrus 🪵

**Compogo Logrus** — это готовая интеграция [logrus](https://github.com/sirupsen/logrus) с фреймворком [Compogo](https://github.com/Compogo/compogo). Добавляется одной строкой и автоматически получает имя приложения из конфига, настраивается через флаги и поддерживает дочерние логгеры.

## 🚀 Установка

```bash
go get github.com/Compogo/logrus
```

### 📦 Использование
```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/logrus"
    "github.com/Compogo/myapp/service"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        logrus.WithLogrus(),                    // ← одна строка
        compogo.WithComponents(
            service.Component,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## ✨ Возможности
### 🎯 Автоматический префикс с именем приложения

Все сообщения автоматически получают префикс [appname]. Дочерние логгеры добавляют свой:
```go
type Service struct {
    log logger.Logger
}

func NewService(log logger.Logger) *Service {
    child := log.GetLogger("service")
    child.Info("starting") // [myapp] [service] starting
}
```

### 🎚️ Уровни логирования
Уровень задаётся через флаг --logger.level:

```bash
./myapp --logger.level=debug
```

Доступные уровни: panic, error, warn, info, debug.

### 🔧 Разделение stdout/stderr

* Info, Debug, Print → stdout
* Warn, Error, Panic → stderr

### 🪝 Поддержка хуков
```go
decorator.AddHook(sentry.NewHook(...))
```

### ⚙️ Как это работает

1. Init — регистрирует конфиг и декоратор в DI
2. BindFlags — добавляет флаг --logger.level
3. PreRun — загружает конфиг, подставляет имя приложения, устанавливает уровень
