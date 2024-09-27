package logging

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var Default *logrus.Logger

func StartLogrus(logLevel uint8) {
	Default = logrus.StandardLogger()
	switch logLevel {
	case 1:
		logrus.SetLevel(logrus.FatalLevel)
	case 2:
		logrus.SetLevel(logrus.ErrorLevel)
	case 3:
		logrus.SetLevel(logrus.WarnLevel)
	case 4:
		logrus.SetLevel(logrus.InfoLevel)
	case 5:
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.WarnLevel)
		logLevel = 3
	}

	Default.Warnf("activated LogLevel: %d", logLevel)

	logrus.SetOutput(os.Stderr)
	logrus.SetReportCaller(true)

	_, fpath, _, _ := runtime.Caller(0)
	fpath = strings.TrimSuffix(fpath, "/pkg/logging/logrus.go")
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
		PadLevelText:  true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := strings.Replace(f.File, fpath, ".", -1)

			return fmt.Sprintf("%s:%d\n\t", filename, f.Line), fmt.Sprintf(" %s()", f.Function)
		},
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  false,
	})
}
