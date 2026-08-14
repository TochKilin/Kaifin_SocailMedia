package logs

import (
	// standard library
	"fmt"

	"github.com/rs/zerolog"
)

// use when app run to show log
func NewLog(logLevel string) {
	switch logLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		// Info
		// Warn
		// Error
		// Fatal
		// Panic
	case "tracing":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		// Error
		// Fatal
		// Panic
	default:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
	fmt.Println("log level", zerolog.GlobalLevel())
}
