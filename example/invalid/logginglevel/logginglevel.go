// Package logginglevel contains a forbidden log level.
package logginglevel

type event struct{}

// Msg logs the message.
func (event) Msg(text string) {
}

type logger struct{}

// Info starts an info-level log event.
func (logger) Info() event {
	return event{}
}

var log = logger{}

func example() {
	log.Info().Msg("bad")
}
