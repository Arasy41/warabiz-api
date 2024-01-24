package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"warabiz/api/config"

	"github.com/sirupsen/logrus"
)

func (l *apiLogger) getEntry() *logrus.Entry {
	if l.entry == nil {
		l.entry = logrus.NewEntry(l.logger)
	}

	file, line := l.getCallerInfo(l.cfg.Logger.CallerSkipper)
	if file == "" || line == 0 {
		return l.entry
	}
	return l.entry.WithField("caller", fmt.Sprintf("%v:%v", file, line))
}

func (l *apiLogger) getCallerInfo(skip int) (string, int) {

	if l.cfg.Logger.DisableCaller {
		return "", 0
	}

	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("Could not retrieve caller information\n")
	}

	if skip == l.cfg.Logger.CallerSkipper {
		if getPackageName(pc) == "exception" {
			return l.getCallerInfo(l.cfg.Logger.CallerSkipper + 1)
		} else if getPackageName(pc) == "main" {
			return l.getCallerInfo(4)
		}
	}

	file = formatFilePath(file)
	return file, line
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"trace": logrus.TraceLevel,
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(cfg *config.Config) logrus.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return logrus.DebugLevel
	}
	return level
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func getPackageName(pc uintptr) string {
	functionName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(functionName, ".")
	return strings.Join(parts[:len(parts)-1], ".")
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
