package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func SetLogInfo() {
	Log.SetLevel(logrus.DebugLevel)
	Log.SetOutput(os.Stdout)
}
