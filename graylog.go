package logging

import (
	graylog "github.com/gemnasium/logrus-graylog-hook"
)

func AddGraylogHook(log *Logger, graylogUrl, containerName string) {
	hook := graylog.NewGraylogHook(graylogUrl, map[string]interface{}{"container_name": containerName})
	log.Logger.AddHook(hook)
}
