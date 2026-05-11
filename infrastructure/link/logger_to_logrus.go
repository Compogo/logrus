package link

import (
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/types/linker"
	"github.com/sirupsen/logrus"
)

var (
	LoggerLevelToLogrusLevel = linker.NewLinker[logger.Level, logrus.Level](
		linker.Link(logger.Panic, logrus.PanicLevel),
		linker.Link(logger.Error, logrus.ErrorLevel),
		linker.Link(logger.Warn, logrus.WarnLevel),
		linker.Link(logger.Info, logrus.InfoLevel),
		linker.Link(logger.Debug, logrus.DebugLevel),
	)
)
