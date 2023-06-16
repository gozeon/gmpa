package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func SetLogInfo() {
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
}
