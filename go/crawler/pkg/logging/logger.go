package logging

import (
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger
var l sync.Mutex

// Init log
func Init(logName string, ropts ...zap.Option) {
	l.Lock()
	defer l.Unlock()
	_logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger = _logger
}

// Logger for pansuku
func Logger() *zap.Logger {
	if logger == nil {
		Init("")
		logger.Warn("logger not initialized")
	}

	return logger
}

// Sugger log
func Suger() *zap.SugaredLogger {
	return Logger().Sugar()
}
