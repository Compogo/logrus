package link

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/types/linker"
	"github.com/sirupsen/logrus"
)

var (
	// LoggerLevelToLogrusLevel связывает уровни Compogo с уровнями Logrus.
	// Используется для установки уровня логирования и фильтрации.
	LoggerLevelToLogrusLevel = linker.NewLinker[compogo.Level, logrus.Level](
		linker.Link(compogo.Panic, logrus.PanicLevel),
		linker.Link(compogo.Error, logrus.ErrorLevel),
		linker.Link(compogo.Warn, logrus.WarnLevel),
		linker.Link(compogo.Info, logrus.InfoLevel),
		linker.Link(compogo.Debug, logrus.DebugLevel),
	)
)
