package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gookit/config/v2"
	"github.com/sirupsen/logrus"
)

var once sync.Once
var log *logrus.Logger

func GetLogger() *logrus.Logger {
	once.Do(func() {
		log = logrus.New()
		if strings.ToUpper(config.String("logFormat")) == "JSON" {
			log.Formatter = new(logrus.JSONFormatter)
		} else {
			log.SetFormatter(&logrus.TextFormatter{
				DisableColors: true,
				FullTimestamp: true,
			})
		}

		log.Out = os.Stdout
		logLevel := strings.ToUpper(config.String("logLevel"))
		switch logLevel {
		case "TRACE":
			log.Level = logrus.TraceLevel
		case "DEBUG":
			log.Level = logrus.DebugLevel
		case "INFO":
			log.Level = logrus.InfoLevel
		case "ERROR":
			log.Level = logrus.ErrorLevel
		case "FATAL":
			log.Level = logrus.FatalLevel
		case "PANIC":
			log.Level = logrus.PanicLevel
		default:
			log.Level = logrus.TraceLevel
			log.Warn(fmt.Sprintf("Unknown log level (%s), providing trace level output", logLevel))
		}
	})

	return log
}
