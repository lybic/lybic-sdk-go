package lybic

// Logger is an interface for logging messages at different levels.
//
//		This interface defines methods for logging debug, info, warning, and error messages.
//	 Support: zap（SugaredLogger）、logrus、zerolog、slog
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// NewEmptyLogger returns a Logger that does not log anything.
type emptyLogger struct{}

func (emptyLogger) Debug(...interface{}) {}

func (emptyLogger) Info(...interface{}) {}

func (emptyLogger) Warn(...interface{}) {}

func (emptyLogger) Error(...interface{}) {}

func (emptyLogger) Debugf(format string, args ...interface{}) {}

func (emptyLogger) Infof(format string, args ...interface{}) {}

func (emptyLogger) Warnf(format string, args ...interface{}) {}

func (emptyLogger) Errorf(format string, args ...interface{}) {}
