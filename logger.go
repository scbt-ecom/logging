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

func InitLogger() *Logger {
	return &Logger{e}
}

func init() {
	//Logger instance creating
	l := logrus.New()

	//Logger instance settings
	l.SetReportCaller(true)
	l.SetLevel(logrus.TraceLevel)
	l.SetOutput(io.Discard)
	l.Formatter = &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "03:04:05 02/01/2006",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
		},
	}

	l.AddHook(&writerHook{
		Writer:    []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	e = logrus.NewEntry(l)
}
