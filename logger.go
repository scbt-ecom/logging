package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
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
	l.SetLevel(logrus.TraceLevel)
	l.SetOutput(io.Discard)
	l.SetReportCaller(true)

	l.Formatter = &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "03:04:05 02/01/2006",
		//CallerPrettyfier: caller(),
	}

	l.AddHook(&writerHook{
		Writer:    []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	e = logrus.NewEntry(l)
}

func (l *Logger) SetLevel(logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		l.WithError(err).Fatalf("failed to set logger level: %s", logLevel)
	}

	l.Logger.SetLevel(level)
}

func (l *Logger) WithExtraField(key string, value any) *Logger {
	return &Logger{l.WithField(key, value)}
}

type Fields logrus.Fields

func (l *Logger) WithExtraFields(fields Fields) *Logger {
	return &Logger{l.WithFields(logrus.Fields(fields))}
}

//func caller() func(*runtime.Frame) (function string, file string) {
//	return func(f *runtime.Frame) (function string, file string) {
//		p, _ := os.Getwd()
//
//		return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, p), f.Line)
//	}
//}
//
//func (l *Logger) SetFormatter() {
//	l.Logger.SetFormatter(&logrus.TextFormatter{
//		DisableColors:    true,
//		FullTimestamp:    true,
//		TimestampFormat:  "03:04:05 02/01/2006",
//		CallerPrettyfier: caller(),
//	})
//}
