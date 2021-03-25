package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetReportCaller(true)
}
