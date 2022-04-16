package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Logger struct {
	log    *zap.Logger
	closed bool
	mutex  sync.Mutex
}

func NewLogger(encoding Encoding, path string, level Level, callerSkip int) (*Logger, error) {
	cfg := &zap.Config{
		Level:            NewLogLevel(level),
		Encoding:         NewEncoding(encoding),
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{path},
		ErrorOutputPaths: []string{path},
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err := cfg.Build(zap.AddCallerSkip(1 + callerSkip))

	return &Logger{log: log}, err
}

func (l *Logger) Close() {
	l.mutex.Lock()
	if !l.closed {
		l.closed = true
		if l.log != nil {
			l.log.Sync()
		}
	}
	l.mutex.Unlock()
}

func (l *Logger) CloneWithCallerSkip(callerSkip int) *Logger {
	return &Logger{
		log: l.log.WithOptions(zap.AddCallerSkip(callerSkip)),
	}
}

func (l *Logger) Fatal(s string, fields ...Field) {
	l.log.Fatal(s, fields...)
}

func (l *Logger) Err(s string, fields ...Field) {
	l.log.Error(s, fields...)
}

func (l *Logger) Warn(s string, fields ...Field) {
	l.log.Warn(s, fields...)
}

func (l *Logger) Info(s string, fields ...Field) {
	l.log.Info(s, fields...)
}

func (l *Logger) Debug(s string, fields ...Field) {
	l.log.Debug(s, fields...)
}
