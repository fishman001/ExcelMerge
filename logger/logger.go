package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
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
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: false})
	StdLogger = &Logger{
		logger,
	}
	writer, err := mkLogFile()
	if err != nil {
		log.Println(err)
	} else {
		logger.SetOutput(writer)
	}
	log.SetOutput(logger.Writer())
}

func GetStdLogger() *Logger {
	return StdLogger
}

func mkLogFile() (io.Writer, error) {
	format := time.Now().Format("ExcelMerge-200601021504.log")
	writer, err := os.OpenFile(path.Join(".", format), os.O_CREATE|os.O_WRONLY, 0764)
	if err != nil {
		return nil, err
	}
	return writer, nil

}
