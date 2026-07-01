// Package logginglevel shows only the approved Debug and Error log levels.
package logginglevel

type event struct{}

// Msg logs the message.
func (event) Msg(text string) {
}

type logger struct{}

// Debug starts a debug-level log event.
func (logger) Debug() event {
	return event{}
}

// Error starts an error-level log event.
func (logger) Error() event {
	return event{}
}

var log = logger{}

// Emit writes log lines using only the approved levels.
func Emit() {
	log.Debug().Msg("starting")
	log.Error().Msg("failed")
}
