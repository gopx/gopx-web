/*
Package log is responsible for whole application logging system.
It manages many logging settings e.g severity level, output formatting etc.
*/
package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{}
	if env, ok := os.LookupEnv("DEBUG"); ok && env == "1" {
		log.Level = logrus.DebugLevel
	} else {
		log.Level = logrus.InfoLevel
	}
}

// Debug logs a message with the specified context at level Debug on the standard logger.
func Debug(context, msg string) {
	log.WithFields(logrus.Fields{
		"context": context,
	}).Debug(msg)
}

// Info logs a message with the specified context at level Info on the standard logger.
func Info(context, msg string) {
	log.WithFields(logrus.Fields{
		"context": context,
	}).Info(msg)
}

// Warn logs a message with the specified context at level Warn on the standard logger.
func Warn(context, msg string) {
	log.WithFields(logrus.Fields{
		"context": context,
	}).Warn(msg)
}

// Error logs a message with the specified context at level Error on the standard logger.
func Error(context, msg string) {
	log.WithFields(logrus.Fields{
		"context": context,
	}).Error(msg)
}

// Fatal logs a message with the specified context at level Fatal on the standard logger.
func Fatal(context, msg string) {
	log.WithFields(logrus.Fields{
		"context": context,
	}).Fatal(msg)
}

// Panic logs a message with the specified context at level Panic on the standard logger.
func Panic(context, msg string) {
	log.WithFields(logrus.Fields{
		"context": context,
	}).Panic(msg)
}
