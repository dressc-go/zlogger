// zlogger is a small convenience wrapper around [zerolog] to add a componentname
// and timestamp to the log
//
// [zerolog]: https://github.com/rs/zerolog
package zlogger

import (
	"github.com/rs/zerolog"
	"os"
)

const (
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
)

// Set global logging level
func SetGlobalLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

// Initialize a logger, component_name will be included in the
// logger (as will be a timestamp)
func GetLogger(component string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().
		Str("component", component).
		Timestamp().
		Logger()
	return logger
}
