package logger

import (
	"os"
	"warabiz/api/config"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

// Logger
type apiLogger struct {
	cfg    *config.Config
	logger *logrus.Logger
	entry  *logrus.Entry
}

// App Logger constructor
func NewApiLogger(cfg *config.Config) *apiLogger {
	return &apiLogger{cfg: cfg}
}

// Init logger
func (l *apiLogger) InitLogger() {

	if l.logger == nil {

		path := "log/"
		if l.cfg.Server.Env == "Staging" || l.cfg.Server.Env == "Production" {
			path = "/var/log/" + l.cfg.Server.Name + "/"
		}

		isExist, err := dirExists(path)
		if err != nil {
			panic(err)
		}

		if !isExist {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		writer, err := rotatelogs.New(
			path+l.cfg.Server.Name+"-"+"%Y%m%d.log",
			rotatelogs.WithMaxAge(-1),
			rotatelogs.WithRotationCount(l.cfg.Logger.Logrotation.MaxRotation),
			rotatelogs.WithRotationTime(l.cfg.Logger.Logrotation.RotationTime*time.Hour),
		)
		if err != nil {
			panic(err)
		}

		logger := logrus.New()
		logger.SetLevel(l.getLoggerLevel(l.cfg))

		//* Set Hook with writer & formatter for log file
		logger.Hooks.Add(lfshook.NewHook(
			writer,
			&logrus.JSONFormatter{
				TimestampFormat: time.RFC3339,
			},
		))

		//* Logstash hook
		if l.cfg.Logger.Logstash.IsActive {
			l.InitLogstash(logger)
		}

		//* Promtail Loki Hook
		if l.cfg.Logger.Loki.IsActive {
			l.InitLoki(logger)
		}

		//* Console Formatter
		if l.cfg.Server.Env == "Production" {
			logger.SetFormatter(&logrus.TextFormatter{
				DisableColors:   true,
				TimestampFormat: time.RFC3339,
			})
		} else {
			logger.SetFormatter(&easy.Formatter{
				TimestampFormat: time.RFC3339,
				LogFormat:       "[%lvl%] %time% - %msg%\n",
			})
		}

		l.logger = logger
	}
}
