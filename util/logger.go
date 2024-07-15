package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func Initialize() {
	Logger.SetFormatter(&log.JSONFormatter{})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(log.WarnLevel)
}
