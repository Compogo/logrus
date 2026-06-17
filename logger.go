package logrus

import (
	"os"

	"github.com/Compogo/compogo"
	"github.com/Compogo/logrus/infrastructure/link"
	"github.com/Compogo/types/linker"
	"github.com/sirupsen/logrus"
)

// Logger реализует интерфейс compogo.Logger на основе Logrus.
// Поддерживает иерархическую структуру логгеров через GetLogger,
// раздельные потоки вывода (stdout/stderr) и добавление хуков.
//
// Иерархия:
//   - Корневой логгер создаётся через NewLogger()
//   - Вложенные логгеры создаются через GetLogger()
//   - Каждый вложенный логгер добавляет префикс [имя] к сообщениям
//
// Пример:
//
//	logger := NewLogger()
//	httpLogger := logger.GetLogger("http")
//	httpLogger.Info("Server started") // выведет: "[http] Server started"
type Logger struct {
	name    string
	stdErr  *logrus.Logger
	stdOut  *logrus.Logger
	parent  *Logger
	childes *linker.Linker[string, *Logger]
}

// NewLogger создаёт новый корневой логгер.
// Настраивает вывод в stdout/stderr и устанавливает TextFormatter.
//
// Пример:
//
//	logger := NewLogger()
//	logger.Info("Application started")
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

// Panicf логирует сообщение с уровнем Panic и вызывает panic.
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

// Panic логирует сообщение с уровнем Panic и вызывает panic.
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

// Errorf логирует форматированное сообщение с уровнем Error.
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

// Error логирует сообщение с уровнем Error.
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

// Warnf логирует форматированное сообщение с уровнем Warn.
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

// Warn логирует сообщение с уровнем Warn.
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

// Infof логирует форматированное сообщение с уровнем Info.
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

// Info логирует сообщение с уровнем Info.
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

// Debugf логирует форматированное сообщение с уровнем Debug.
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

// Debug логирует сообщение с уровнем Debug.
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

// Printf логирует форматированное сообщение без уровня (аналог fmt.Printf).
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

// Print логирует сообщение без уровня (аналог fmt.Print).
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

// GetStdErr возвращает внутренний логгер для stderr.
// Позволяет напрямую работать с Logrus-логгером.
func (logger *Logger) GetStdErr() *logrus.Logger {
	return logger.stdErr
}

// GetStdOut возвращает внутренний логгер для stdout.
// Позволяет напрямую работать с Logrus-логгером.
func (logger *Logger) GetStdOut() *logrus.Logger {
	return logger.stdOut
}

// SetLevel устанавливает уровень логирования для stdout и stderr.
//
// Пример:
//
//	logger.SetLevel(compogo.Debug)
func (logger *Logger) SetLevel(level compogo.Level) error {
	logrusLevel, err := link.LoggerLevelToLogrusLevel.Get(level)
	if err != nil {
		return err
	}

	logger.stdOut.SetLevel(logrusLevel)
	logger.stdErr.SetLevel(logrusLevel)

	return nil
}

// AddHook добавляет хук Logrus к обоим логгерам (stdout и stderr).
// Используется для добавления метрик, сентинел и других расширений.
//
// Пример:
//
//	logger.AddHook(&MetricHook{})
func (logger *Logger) AddHook(hook logrus.Hook) {
	logger.stdOut.AddHook(hook)
	logger.stdErr.AddHook(hook)
}

// GetLogger возвращает дочерний логгер с указанным именем.
// Дочерние логгеры кэшируются и переиспользуются.
// Сообщения от дочерних логгеров получают префикс [имя].
//
// Реализует интерфейс compogo.Logger.
//
// Пример:
//
//	httpLogger := logger.GetLogger("http")
//	httpLogger.Info("Server started") // "[http] Server started"
func (logger *Logger) GetLogger(name string) compogo.Logger {
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
