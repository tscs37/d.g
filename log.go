package discorddotgo

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = nil
}

func EnableLogging() {
	log.Out = os.Stdout
}

func EnableDebug() {
	EnableLogging()
	log.Level = logrus.DebugLevel
}
