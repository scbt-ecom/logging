package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func InitLogger(logLevel string) *Logger {
	return &Logger{e}
}

func init() {
	//Logger instance creating
	log := logrus.New()

	//Logger instance settings
	log.SetLevel(logrus.TraceLevel)
	log.SetReportCaller(true)
	log.SetOutput(io.Discard)
	log.Formatter = &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "03:04:05 02/01/2006",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
		},
	}

	log.AddHook(&writerHook{
		Writer:    []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	e = logrus.NewEntry(log)
}

func SetLevel(log *Logger, logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.WithError(err).Fatal("failed to set logger level")
	}

	log.Logger.SetLevel(level)
}
