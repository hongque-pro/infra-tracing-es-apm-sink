package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var loggerMutex sync.RWMutex // guards access to global logger state

// loggers is the set of loggers in the system
var loggers = make(map[string]*zap.SugaredLogger)

var levels = make(map[string]zap.AtomicLevel)
var defaultLevel zapcore.Level = zapcore.WarnLevel

var logCore = &lockedMultiCore{}

func GetLogger(name string) *zap.SugaredLogger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()
	log, ok := loggers[name]
	if !ok {
		levels[name] = zap.NewAtomicLevelAt(defaultLevel)
		log = zap.New(logCore).
			WithOptions(zap.IncreaseLevel(levels[name])).
			Named(name).
			Sugar()

		loggers[name] = log
	}

	return log
}
