package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LoggerFactory struct {
	fields logrus.Fields
}

type LmptLogger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

func (l LmptLogger) Debug(args ...interface{}) { l.logger.WithFields(l.fields).Debug(args...) }
func (l LmptLogger) Debugf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Debugf(format, args...)
}
func (l LmptLogger) Info(args ...interface{}) { l.logger.WithFields(l.fields).Info(args...) }
func (l LmptLogger) Infof(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Infof(format, args...)
}
func (l LmptLogger) Warning(args ...interface{}) { l.logger.WithFields(l.fields).Warn(args...) }
func (l LmptLogger) Warningf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Warnf(format, args...)
}
func (l LmptLogger) Error(args ...interface{}) { l.logger.WithFields(l.fields).Error(args...) }
func (l LmptLogger) Errorf(format string, args ...interface{}) {
	l.logger.WithFields(l.fields).Errorf(format, args...)
}
func (l LmptLogger) WithError(err error, message string) {
	l.logger.WithFields(l.fields).WithError(err).Error(message)
}
func (l LmptLogger) WithErrorf(err error, format string, args ...interface{}) {
	l.logger.WithFields(l.fields).WithError(err).Errorf(format, args...)
}

func (f *LoggerFactory) NewJsonLogger() LmptLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	setLevel(logger)

	return LmptLogger{
		logger: logger,
		fields: f.fields,
	}
}

func InitLogger(inputFields map[string]interface{}) *LoggerFactory {
	converted := logrus.Fields{}
	for k, v := range inputFields {
		converted[k] = v
	}
	return &LoggerFactory{fields: converted}
}

func setLevel(logger *logrus.Logger) {
	level, err := logrus.ParseLevel(getEnv("LOG_LEVEL", "info"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
