package logging

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

type logger struct {
	*logrus.Entry
}

func (l *logger) WithField(key string, value interface{}) Logger {
	return &logger{
		l.Entry.WithField(key, value),
	}
}

func (l *logger) WithFields(fields map[string]interface{}) Logger {
	return &logger{
		l.Entry.WithFields(fields),
	}
}

func GetLogger(logLevel string) Logger {
	l := logrus.New()
	// pretty output
	l.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
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

	entry := logrus.NewEntry(l)
	return &logger{entry}
}
