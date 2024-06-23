package config

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"os"
)

func ConfigLogging() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
}
