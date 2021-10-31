package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogErrorIfNotNil provides the provided err as an argument to the logging function using the provided template
// so it is equivalent to LogErrorIf(err != nil, template, err)
func LogErrorIfNotNil(err error, template string) {
	logIf(err != nil, zapcore.ErrorLevel, template, err)
}

func LogErrorIf(condition bool, template string, args ...interface{}) {
	logIf(condition, zapcore.ErrorLevel, template, args...)
}

func LogWarnIf(condition bool, template string, args ...interface{}) {
	logIf(condition, zapcore.WarnLevel, template, args...)
}

func LogInfoIf(condition bool, template string, args ...interface{}) {
	logIf(condition, zapcore.InfoLevel, template, args...)
}

func LogDebugIf(condition bool, template string, args ...interface{}) {
	logIf(condition, zapcore.DebugLevel, template, args...)
}

// logIf method for checking a condition and logging if condition is met (usually checking if error is nil)
func logIf(condition bool, level zapcore.Level, template string, args ...interface{}) {
	if condition {
		switch level {
		case zapcore.DebugLevel:
			zap.S().Debugf(template, args...)
		case zapcore.InfoLevel:
			zap.S().Infof(template, args...)
		case zapcore.ErrorLevel:
			zap.S().Errorf(template, args...)
		case zapcore.FatalLevel:
			zap.S().Fatalf(template, args...)
		case zapcore.WarnLevel:
			zap.S().Warnf(template, args...)
		case zapcore.PanicLevel:
			zap.S().Panicf(template, args...)
		case zapcore.DPanicLevel:
			zap.S().DPanicf(template, args...)
		}
	}
}

type FiberLogWriter struct{}

func (l *FiberLogWriter) Write(p []byte) (n int, err error) {
	zap.S().Info(string(p))
	return len(p), nil
}
