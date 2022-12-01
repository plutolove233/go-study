package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var logger = logrus.New()

func main() {
	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		"system-%Y%m%d.log",
		rotatelogs.WithLinkName("system.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		println("failed create rotatelogs: ", err.Error())
		return
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(lfHook)

	logger.WithFields(logrus.Fields{
		"animal": "dogs",
		"size":   10,
	}).Info("a group of animal appeared")
}
