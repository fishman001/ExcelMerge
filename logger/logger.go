package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

func init() {
	initStdLogger()
}

type Logger struct {
	*logrus.Logger
}

func (l *Logger) SetLevel(level string) {
	level = strings.ToUpper(level)
	switch level {
	case "INFO":
		l.Logger.SetLevel(logrus.InfoLevel)
	case "WARN":
		l.Logger.SetLevel(logrus.WarnLevel)
	case "debug":
		l.Logger.SetLevel(logrus.DebugLevel)
	default:
		l.Logger.SetLevel(logrus.ErrorLevel)
	}
}

var StdLogger *Logger

func initStdLogger() {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.InfoLevel)
	// logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	StdLogger = &Logger{
		logger,
	}
	log.SetOutput(logger.Writer())
}

func GetStdLogger() *Logger {
	return StdLogger
}
