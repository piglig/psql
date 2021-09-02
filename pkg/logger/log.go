package logger

import (
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

// InitLogger initialize the logger instance
func InitLogger() {
	logger = logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}
	//if !util.IsProdEnv() {
	//	logger.SetReportCaller(true)
	//}
	logger.Out = ansicolor.NewAnsiColorWriter(os.Stdout)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Debug(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Info(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.Warn(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.Error(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.Fatalln(v)
}

func WithFields(fileds logrus.Fields) *logrus.Entry {
	return logger.WithFields(fileds)
}
