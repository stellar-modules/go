package test

import (
	"github.com/sirupsen/logrus"
	"github.com/stellar-modules/go/sdk/support/log"
)

var testLogger *log.Entry

func init() {
	testLogger = log.New()
	testLogger.Entry.Logger.Formatter.(*logrus.TextFormatter).DisableColors = true
	testLogger.Entry.Logger.Level = logrus.DebugLevel
}
