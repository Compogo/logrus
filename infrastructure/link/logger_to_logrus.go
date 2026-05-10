package link

import (
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/types/linker"
	"github.com/sirupsen/logrus"
)

var (
	LoggerLevelToLogrusLevel = linker.NewLinker[logger.Level, logrus.Level](
		linker.NewLink(logger.Panic, logrus.PanicLevel),
		linker.NewLink(logger.Error, logrus.ErrorLevel),
		linker.NewLink(logger.Warn, logrus.WarnLevel),
		linker.NewLink(logger.Info, logrus.InfoLevel),
		linker.NewLink(logger.Debug, logrus.DebugLevel),
	)
)
