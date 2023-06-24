package util

import (
	"flag"
	"os"

	"github.com/shivam904455/Home-Automation/model"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Out = os.Stdout
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

}

// func  SetLoger()*logrus.Logger{

func SetLoger() *logrus.Logger {

	// Define a command-line flag for log level
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "Log level (debug, info, warn, error)")

	// Parse the command-line flags
	flag.Parse()
	// Set the log level based on the flag value
	switch *logLevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
	return Logger
}

func Log(logLevel, pkgLevel, functionName string, message, parameater interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parameater == nil {
			Logger.Debugf("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Debugf("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.Value, parameater)
		}
	case model.LogLevelWarning:
		if parameater == nil {
			Logger.Warnf("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Warnf("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.Value, parameater)
		}
	case model.LogLevelError:
		if parameater == nil {
			Logger.Errorf("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Errorf("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.LogLevelError, parameater)
		}
	default:
		if parameater == nil {
			Logger.Infof("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Infof("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.Value, parameater)
		}
	}
}
