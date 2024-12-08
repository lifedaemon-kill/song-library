package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Log *logrus.Logger

func GetNewLogger() (*logrus.Logger, error) {
	logFilePath := "internal/pkg/logger/api.log"

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log := logrus.New()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetLevel(logrus.TraceLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return log, nil
}
