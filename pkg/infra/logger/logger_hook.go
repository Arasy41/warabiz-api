package logger

import (
	"fmt"
	"net"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/ic2hrmk/lokigrus"
	"github.com/sirupsen/logrus"
)

func (l *apiLogger) InitLoki(logger *logrus.Logger) {

	appLabels := make(map[string]string)
	appLabels["app_name"] = l.cfg.Server.Name
	appLabels["app_env"] = l.cfg.Server.Env

	promtailHook, err := lokigrus.NewPromtailHook(l.cfg.Logger.Loki.URI, appLabels)
	if err != nil {
		fmt.Println("failed to connect loki err: ", err.Error())
		return
	}
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	logger.AddHook(promtailHook)
}

func (l *apiLogger) InitLogstash(logger *logrus.Logger) {

	appLabels := logrus.Fields{}
	appLabels["app_name"] = l.cfg.Server.Name
	appLabels["app_env"] = l.cfg.Server.Env

	conn, err := net.Dial("tcp", l.cfg.Logger.Logstash.URI)
	if err != nil {
		fmt.Println("failed to connect logstash err: ", err.Error())
		return
	}
	logstashHook := logrustash.New(conn, logrustash.DefaultFormatter(appLabels))
	logger.Hooks.Add(logstashHook)
}
