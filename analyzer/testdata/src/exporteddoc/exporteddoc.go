package exporteddoc

// Good does a thing.
func Good() {}

func Bad() {} // want "exported function 'Bad' should have a documentation comment"

func unexported() {}

type Thing struct{}

// Run runs the thing.
func (thing *Thing) Run() {}

func (thing *Thing) Stop() {} // want "exported method 'Stop' should have a documentation comment"
