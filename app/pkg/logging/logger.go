package logging

import (
	"log"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func GetLogger(logLevel string) *Logger {
	l := logrus.New()
	// pretty output
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// set logging level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("cannot parse log level. %s", err)
	}

	if level == logrus.DebugLevel {
		l.SetReportCaller(true)
	}

	l.SetLevel(level)
	return &Logger{l}
}
