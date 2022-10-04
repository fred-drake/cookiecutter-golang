package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	var log = logrus.New()
	if os.Getenv("LOG_FORMAT") == "JSON" {
		log.Formatter = new(logrus.JSONFormatter)
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}
	
	log.Out = os.Stdout
	if os.Getenv("TRACE") != "" {
		log.Level = logrus.TraceLevel
		log.Trace("Trace log level set.")
	} else if os.Getenv("DEBUG") != "" {
		log.Level = logrus.DebugLevel
		log.Debug("Debug log level set.")
	} else {
		log.Level = logrus.InfoLevel
	}

	return log
}
