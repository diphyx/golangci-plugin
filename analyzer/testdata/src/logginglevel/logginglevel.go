package logginglevel

type event struct{}

func (event) Msg(string) {}

type logger struct{}

func (logger) Debug() event { return event{} }

func (logger) Error() event { return event{} }

func (logger) Info() event { return event{} }

var log logger

func examples() {
	log.Debug().Msg("ok")
	log.Error().Msg("ok")
	log.Info().Msg("bad") // want "log level 'Info' is not allowed; use only Debug or Error"
}
